const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FooDecoder", function () {
  let FooDecoder;
  let fooDecoder;

  beforeEach(async function () {
    FooDecoder = await ethers.getContractFactory("FooDecoder");
    fooDecoder = await FooDecoder.deploy();
  });

  it("should decode specific ABI-encoded bytes", async function () {
    // The encoded data you want to test
    const encodedData = "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000030390000000000000000000000000000000000000000000000000000000000bc614e0000000000000000000000000000000000000000000000000000011f71fb04cb000000000000000000000000000000000000000000000000ab54a98ceb1f0ad200000000000000000000000000000000000000018ee90ff6c373e0ee4e3f0ad2000000000000000000000000000000000000000000000000000000000000000100000000000000000000000090abcdef1234567890abcdef1234567890abcdef00000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000180000000000000000000000000000000000000000000000000000000000000000d48656c6c6f2c204170746f7321000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009627974655f646174610000000000000000000000000000000000000000000000";

    // Call the decode function
    const result = await fooDecoder.decode(encodedData);

    // Expected values based on the encoded data
    expect(result[0]).to.equal(32); // uint8
    expect(result[1]).to.equal(12); // uint16
    expect(result[2]).to.equal(12345); // uint32
    expect(result[3]).to.equal(BigInt("12345678")); // uint64
    expect(result[4]).to.equal(BigInt("1234567890123")); // uint128
    expect(result[5]).to.equal(BigInt("12345678901234567890123456789")); // uint256
    expect(result[6]).to.equal(BigInt("123456789012345678901234567890123456789")); // This appears to be another large number
    expect(result[7]).to.equal(true); // bool
    expect(result[8]).to.equal("0x90abcdef1234567890abcdef1234567890abcdef"); // address
    expect(result[9]).to.equal("Hello, Aptos!"); // string
    expect(result[10]).to.equal("0x627974655f64617461"); // bytes - "byte_data" in hex

    console.log("Decoded results:");
    console.log("uint8:", result[0]);
    console.log("uint16:", result[1]);
    console.log("uint32:", result[2]);
    console.log("uint64:", result[3]);
    console.log("uint128:", result[4]);
    console.log("uint256:", result[5]);
    console.log("bool:", result[6]);
    console.log("address:", result[7]);
    console.log("string:", result[8]);
    console.log("bytes:", result[9]);
  });
});