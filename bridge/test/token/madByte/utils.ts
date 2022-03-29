import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { expect } from "../../chai-setup";
import { Fixture, getFixture } from "../../setup";
import { getState, init, showState } from "./setup";

describe("Testing MadByte Utils methods", async () => {
  let admin: SignerWithAddress;
  let user: SignerWithAddress;
  let user2: SignerWithAddress;
  let fixture: Fixture;
  let marketSpread = 4;

  beforeEach(async function () {
    fixture = await getFixture();
    let signers = await ethers.getSigners();
    [admin, user, user2] = signers;
    await init(fixture);
    showState("Initial", await getState(fixture));
  });

  it("Should calculate correct bounding curves", async () => {
    let eth = 10000;
    let madBytes = await fixture.madByte.ethToMadByte(
      await fixture.madByte.getPoolBalance(),
      ethers.utils.parseEther((eth / marketSpread).toString())
    );
    expect(madBytes).to.be.equal(BigInt("936764568799449143863271"));
    eth = 20000;
    madBytes = await fixture.madByte.ethToMadByte(
      await fixture.madByte.getPoolBalance(),
      ethers.utils.parseEther((eth / marketSpread).toString())
    );
    expect(madBytes).to.be.equal(BigInt("1005000000000000000000000"));
    eth = 25000;
    madBytes = await fixture.madByte.ethToMadByte(
      await fixture.madByte.getPoolBalance(),
      ethers.utils.parseEther((eth / marketSpread).toString())
    );
    expect(madBytes).to.be.equal(BigInt("1007899288252135716968558"));
  });
});