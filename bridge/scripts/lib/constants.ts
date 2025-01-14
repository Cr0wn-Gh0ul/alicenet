import "process";
export const MULTICALL_GAS_LIMIT = "10000000";
// contract names
export const ALICENET_FACTORY = "AliceNetFactory";
export const END_POINT = "MockEndPoint";
export const LOCK_UP = "Lockup";
export const MOCK_FACTORY = "MockFactory";
export const MOCK = "MockBaseContract";
export const STAKING_ROUTER_V1 = "StakingRouterV1";
export const UTILS = "Utils";
// function names
export const INITIALIZER = "initialize";
export const DEPLOY_CREATE2 = "deployCreate2";
export const DEPLOY_CREATE = "deployCreate";
export const DEPLOY_METAMORPHIC = "deployMetamorphic";
export const UPGRADE_PROXY = "upgradeProxy";
// Factory event names
export const DEPLOYED_PROXY = "DeployedProxy";
export const DEPLOYED_RAW = "DeployedRaw";
export const DEPLOYED_TEMPLATE = "DeployedTemplate";
export const PROXY = "Proxy";
// Event variable names
export const CONTRACT_ADDR = "contractAddr";

// Hardhat CLI Task names
export const TASK_DEPLOY_ALL_CONTRACTS = "deploy-all-contracts";
export const TASK_DEPLOY_CREATE = "deploy-create";
export const TASK_DEPLOY_CREATE_AND_REGISTER = "deploy-create-and-register";
export const TASK_DEPLOY_CREATE2 = "deploy-create2";
export const TASK_DEPLOY_FACTORY = "deploy-factory";
export const DEPLOY_FACTORY = "deployFactory";
export const TASK_DEPLOY_METAMORPHIC = "deploy-metamorphic";
export const DEPLOY_PROXY = "deployProxy";
export const TASK_DEPLOY_PROXY = "deploy-proxy";
export const TASK_DEPLOY_CONTRACTS = "deploy-contracts";
export const DEPLOY_UPGRADEABLE_PROXY = "deployUpgradeableProxy";
export const TASK_DEPLOY_UPGRADEABLE_PROXY = "deploy-upgradeable-proxy";
export const DEPLOYMENT_ARG_PATH = `../scripts/base-files/deploymentArgs`;
export const DEPLOYMENT_LIST_PATH = `../scripts/base-files/deploymentList`;
export const FACTORY_DEPLOYMENT: string = "factory";
export const FACTORY_STATE_PATH = `../scripts/generated/factoryState`;
export const LOGIC_ADDR = "LogicAddress";
export const META_ADDR = "MetaAddress";
export const MOCK_INITIALIZABLE = "MockInitializable";
export const MULTI_CALL_DEPLOY_PROXY = "multiCallDeployProxy";
export const TASK_MULTI_CALL_DEPLOY_PROXY = "multi-call-deploy-proxy";
export const TASK_MULTI_CALL_DEPLOY_METAMORPHIC =
  "multi-call-deploy-metamorphic";
export const TASK_FULL_MULTI_CALL_DEPLOY_PROXY = "full-multi-call-deploy-proxy";
export const MULTI_CALL_UPGRADE_PROXY = "multiCallUpgradeProxy";
export const TASK_MULTI_CALL_UPGRADE_PROXY = "multi-call-upgrade-proxy";
export const ONLY_PROXY = "onlyProxy";
export const PROXY_ADDR = "ProxyAddress";
export const RECEIPT = "receipt";
export const UPGRADE_DEPLOYED_PROXY = "upgradeDeployedProxy";
export const TASK_UPGRADE_DEPLOYED_PROXY = "upgrade-deployed-proxy";
export const TASK_UPGRADE_PROXY = "upgrade-proxy";
export const UPGRADEABLE_DEPLOYMENT: string = "deployUpgradeable";
export const TASK_UPGRADEABLE_DEPLOYMENT: string = "deploy-upgradeable";
// default paths
export const DEPLOYMENT_LIST_FPATH = "/deploymentList";
export const DEPLOYMENT_ARGS_TEMPLATE_FPATH = "/deploymentArgsTemplate";
export const DEFAULT_CONFIG_DIR = "../scripts/base-files";
export const DEFAULT_FACTORY_STATE_OUTPUT_DIR = "../scripts/generated";
export const BASE_CONFIG_PATH = `../scripts/base-files/baseConfig`;
export const HARDHAT_CHAIN_ID = 1337;
export const env = (): string => {
  let _env = process.env.DEPLOYMENT_ENVIRONMENT;
  if (typeof _env === "undefined") {
    _env = "dev";
  }
  return _env;
};
