var Election = artifacts.require("./myContract.sol");

module.exports = function(deployer) {
  deployer.deploy(myContract);
};
