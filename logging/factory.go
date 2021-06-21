package logging

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/MadBase/MadNet/constants"
	"github.com/sirupsen/logrus"
)

type loggerDetails struct {
	sync.Once
	loggers map[string]*logrus.Logger
}

var loggers loggerDetails // map[string]*logrus.Logger{}

//LogFormatter applies consistent formatting to every message
type LogFormatter struct {
	frm func(*logrus.Entry) ([]byte, error)
}

//Format satisfies logrus' Format interface while staying flexible
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.frm != nil {
		return f.frm(entry)
	}

	genericFormatter := logrus.TextFormatter{PadLevelText: true, TimestampFormat: "1-2|15:04:05.000", FullTimestamp: true}

	return genericFormatter.Format(entry)
}

//LogWriter struct used to provide an io.Writer
type LogWriter struct {
	logger *logrus.Logger
	level  logrus.Level
}

func (logWriter *LogWriter) Write(p []byte) (n int, err error) {
	logWriter.logger.Log(logWriter.level, strings.TrimRight(string(p), "\n"))
	return len(p), nil
}

func (ld *loggerDetails) init() {
	ld.Do(func() {
		loggers.loggers = make(map[string]*logrus.Logger, len(constants.ValidLoggers))
		for _, loggerName := range constants.ValidLoggers {
			formatter := &LogFormatter{frm: func(entry *logrus.Entry) ([]byte, error) {
				defaultFormat, e := (&logrus.TextFormatter{
					PadLevelText:    true,
					TimestampFormat: "1-2|15:04:05.000",
					FullTimestamp:   true}).Format(entry)

				label := fmt.Sprintf("%-10s ", loggerName)

				line := bytes.Join([][]byte{
					[]byte(label), defaultFormat},
					[]byte(" "))

				return line, e
			}}

			logger := logrus.New()
			logger.SetFormatter(formatter)
			logger.SetLevel(logrus.InfoLevel)

			loggers.loggers[loggerName] = logger
		}
	})
}

//GetLogger either returns an existing logger for package specified or creates a new one
func GetLogger(name string) *logrus.Logger {
	loggers.init()
	logger, exists := loggers.loggers[name]
	if !exists {
		panic(fmt.Sprintf("Invalid logger requested: %v", name))
	}

	return logger
}

//GetLogWriter returns an io.Writer that maps to the named logger at the specified level
func GetLogWriter(pkgName string, level logrus.Level) *LogWriter {
	logger := GetLogger(pkgName)

	return &LogWriter{logger, level}
}

// GetKnownLoggers returns all loggers currently configured
func GetKnownLoggers() []*logrus.Logger {
	loggers.init()
	ret := make([]*logrus.Logger, 0)
	for _, logger := range loggers.loggers {
		ret = append(ret, logger)
	}

	return ret
}
