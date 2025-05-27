require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.28",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200, // Default value, adjust if needed
      },
      viaIR: true, // Enable IR pipeline
    },
  },
  networks: {
    hardhat: {},
    // Add testnet/mainnet configs if needed
  },
};