package monitor

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/MadBase/MadNet/blockchain/dkg/dkgtasks"
	"github.com/MadBase/MadNet/blockchain/objects"
	"github.com/MadBase/MadNet/config"
	"github.com/MadBase/MadNet/consensus/admin"
	"github.com/MadBase/MadNet/logging"
	"github.com/dgraph-io/badger/v2"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/sirupsen/logrus"
)

var (
	// ErrUnknownRequest a service was invoked but couldn't figure out which
	ErrUnknownRequest = errors.New("unknown request")

	// ErrUnknownResponse only used when response to a service is not of the expected type
	ErrUnknownResponse = errors.New("response isn't in expected form")
)

// Monitor describes required functionality to monitor Ethereum
type Monitor interface {
	StartEventLoop(accounts.Account) (chan<- bool, error)
	GetStatus() <-chan string
}

type monitor struct {
	database     Database
	bus          Bus
	logger       *logrus.Logger
	tickInterval time.Duration
	timeout      time.Duration
	statusMsg    chan string
	adminHandler *admin.Handlers
	typeRegistry *objects.TypeRegistry
}

// NewMonitor creates a new Monitor
func NewMonitor(db Database, bus Bus, adminHandler *admin.Handlers, tickInterval time.Duration, timeout time.Duration) (Monitor, error) {

	logger := logging.GetLogger("monitor")

	rand.Seed(time.Now().UnixNano())

	// Type registry is used to bidirectionally map a type name string to it's reflect.Type
	// -- This lets us use a wrapper class and unmarshal something where we don't know its type
	//    in advance.
	tr := &objects.TypeRegistry{}

	tr.RegisterInstanceType(&dkgtasks.CompletionTask{})
	tr.RegisterInstanceType(&dkgtasks.DisputeTask{})
	tr.RegisterInstanceType(&dkgtasks.GPKJDisputeTask{})
	tr.RegisterInstanceType(&dkgtasks.GPKSubmissionTask{})
	tr.RegisterInstanceType(&dkgtasks.KeyshareSubmissionTask{})
	tr.RegisterInstanceType(&dkgtasks.MPKSubmissionTask{})
	tr.RegisterInstanceType(&dkgtasks.PlaceHolder{})
	tr.RegisterInstanceType(&dkgtasks.RegisterTask{})
	tr.RegisterInstanceType(&dkgtasks.ShareDistributionTask{})

	return &monitor{
		adminHandler: adminHandler,
		database:     db,
		bus:          bus,
		logger:       logger,
		statusMsg:    make(chan string, 1),
		tickInterval: tickInterval,
		timeout:      timeout,
		typeRegistry: tr,
	}, nil
}

func (mon *monitor) GetStatus() <-chan string {
	return mon.statusMsg
}

// StartEventLoop starts the event loop
func (mon *monitor) StartEventLoop(acct accounts.Account) (chan<- bool, error) {

	logger := mon.logger

	// Load or create initial state
	logger.Info(strings.Repeat("-", 80))
	initialState, err := mon.database.FindState()
	if err != nil {
		logger.Warnf("could not find previous state: %v", err)
		if err != badger.ErrKeyNotFound {
			return nil, err
		}

		logger.Info("Setting initial state to defaults...")
		startingBlock := config.Configuration.Ethereum.StartingBlock
		schedule := NewSequentialSchedule(mon.typeRegistry, mon.adminHandler)
		dkgState := objects.NewDkgState(acct)

		initialState = objects.NewMonitorState(dkgState, schedule)
		initialState.HighestBlockFinalized = uint64(startingBlock)
		initialState.HighestBlockProcessed = uint64(startingBlock)
	}

	initialState.InSync = false
	logger.Info("Current state:")
	logger.Infof("...Ethereum in sync: %v", initialState.EthereumInSync)
	logger.Infof("...Highest block finalized: %v", initialState.HighestBlockFinalized)
	logger.Infof("...Highest block processed: %v", initialState.HighestBlockProcessed)
	logger.Infof("...Monitor tick interval: %v", mon.tickInterval.String())
	logger.Info(strings.Repeat("-", 80))

	cancelChan := make(chan bool)
	go mon.eventLoop(initialState, cancelChan)

	return cancelChan, nil
}

func (mon *monitor) eventLoop(state *objects.MonitorState, cancelChan <-chan bool) error {

	done := false
	for !done {
		select {
		case done = <-cancelChan:
			mon.logger.Warnf("Received cancel request for event loop.")
		case tick := <-time.After(500 * time.Millisecond):
			err := mon.eventLoopTick(state, tick, done)
			if err != nil {
				mon.logger.Warnf("Error occurred: %v", err)
			}
		}
	}

	return nil
}

func (mon *monitor) eventLoopTick(state *objects.MonitorState, tick time.Time, shutdownRequested bool) error {

	logger := mon.logger

	// Make a backup of state to monitor for changes
	originalState := state.Clone()

	// Every tick we request events be processed and we require it doesn't overlap with the next
	resp, err := mon.bus.Request(SvcWatchEthereum, mon.timeout, state)
	if err != nil {
		logger.Warnf("Could not request SvcWatchEthereum: %v", err)
		return err
	}

	select {
	case responseValue := <-resp.Response():
		switch value := responseValue.(type) {
		case *objects.MonitorState:
			diff := originalState.Diff(state)
			if len(diff) > 0 {
				select {
				case mon.statusMsg <- fmt.Sprintf("State \xce\x94 %v", diff):
				default:
				}
				mon.database.UpdateState(value)
			}
			return nil
		case error:
			logger.Warnf("SvcWatchEthereum() : %v", value)
		default:
			logger.Errorf("SvcWatchEthereum() invalid return type: %v", value)
		}
	case to := <-resp.Timeout():
		logger.Warnf("SvcWatchEthereum() : Timeout %v", to)
	}

	return nil
}
