import toml from "@iarna/toml";
import { BigNumberish, BytesLike, ContractReceipt } from "ethers";
import fs from "fs";
import { FACTORY_STATE_PATH } from "../constants";
import { readFactoryState } from "./deploymentConfigUtil";

export type FactoryData = {
  address: string;
  owner?: string;
  gas?: number;
};

export type DeployCreateData = {
  name: string;
  address: string;
  factoryAddress: string;
  gas?: number;
  constructorArgs?: any;
  receipt?: ContractReceipt;
};
export type MetaContractData = {
  metaAddress: string;
  salt: string;
  templateName: string;
  templateAddress: string;
  factoryAddress: string;
  gas?: number;
  initCallData: string;
  receipt?: ContractReceipt;
};
export type TemplateData = {
  name: string;
  address: string;
  factoryAddress: string;
  gas?: number;
  receipt?: ContractReceipt;
  constructorArgs?: string;
};

export interface FactoryConfig {
  [key: string]: any;
}
export type ProxyData = {
  proxyAddress: string;
  salt: BytesLike;
  logicName?: string;
  logicAddress?: string;
  factoryAddress: string;
  gas?: BigNumberish;
  receipt?: ContractReceipt;
  initCallData?: BytesLike;
};

export async function getDefaultFactoryAddress(
  network: string,
  usrPath?: string
): Promise<string> {
  const path =
    usrPath === undefined
      ? FACTORY_STATE_PATH
      : usrPath.replace(/\/+$/, "") + "/factoryState";
  // fetch whats in the factory config file
  const config = await readFactoryState(path);
  return config[network].defaultFactoryAddress;
}

async function writeFactoryState(
  network: string,
  fieldName: string,
  fieldData: any,
  usrPath?: string
) {
  const path =
    usrPath === undefined
      ? FACTORY_STATE_PATH
      : usrPath.replace(/\/+$/, "") + "/factoryState";
  let factoryStateConfig;
  if (fs.existsSync(path)) {
    factoryStateConfig = await readFactoryState(path);
    factoryStateConfig[network] =
      factoryStateConfig[network] === undefined
        ? {}
        : factoryStateConfig[network];
    factoryStateConfig[network][fieldName] = fieldData;
  } else {
    factoryStateConfig = {
      [network]: {
        [fieldName]: fieldData,
      },
    };
  }
  const data = toml.stringify(factoryStateConfig);
  fs.writeFileSync(path, data);
}

export async function updateDefaultFactoryData(
  network: string,
  data: FactoryData,
  usrPath?: string
) {
  await writeFactoryState(network, "defaultFactoryData", data, usrPath);
  await writeFactoryState(
    network,
    "defaultFactoryAddress",
    data.address,
    usrPath
  );
}

export async function updateDeployCreateList(
  network: string,
  data: DeployCreateData,
  usrPath?: string
) {
  await updateList(network, "rawDeployments", data, usrPath);
}

export async function updateTemplateList(
  network: string,
  data: TemplateData,
  usrPath?: string
) {
  await updateList(network, "templates", data, usrPath);
}

export async function updateProxyList(
  network: string,
  data: ProxyData,
  usrPath?: string
) {
  await updateList(network, "proxies", data, usrPath);
}

export async function updateMetaList(
  network: string,
  data: MetaContractData,
  usrPath?: string
) {
  await updateList(network, "staticContracts", data, usrPath);
}
export async function updateList(
  network: string,
  fieldName: string,
  data: any,
  usrPath?: string
) {
  const factoryStateConfig = await readFactoryState(usrPath);
  const output: Array<any> =
    factoryStateConfig[network][fieldName] === undefined
      ? []
      : factoryStateConfig[network][fieldName];
  output.push(data);
  if (data.receipt !== undefined) {
    data.receipt = undefined;
  }
  // write new data to config file
  await writeFactoryState(network, fieldName, output, usrPath);
}