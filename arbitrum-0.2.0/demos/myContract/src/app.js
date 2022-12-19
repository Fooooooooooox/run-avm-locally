/* eslint-env browser */
"use strict";

var $ = require("jquery");
const ethers = require("ethers");
const ArbProvider = require("arb-provider-ethers");

require("bootstrap/dist/css/bootstrap.min.css");

let App = {
  provider: null,
  contracts: {},
  account: "0x0",
  hasVoted: false,

  init: function() {
    return App.initWeb3();
  },

  initWeb3: async function() {
    // Modern dapp browsers...
    let web3Provider = null;
    if (window.ethereum) {
      web3Provider = window.ethereum;
      try {
        // Request account access
        await window.ethereum.enable();
      } catch (error) {
        // User denied account access...
        console.error("User denied account access");
      }
    }
    // Legacy dapp browsers...
    else if (window.web3) {
      web3Provider = window.web3.currentProvider;
    }
    // If no injected web3 instance is detected, fall back to Ganache
    // 用ganache启动的本地l1
    else {
      web3Provider = new ethers.providers.JsonRpcProvider(
        "http://localhost:7545"
      );
    }

    const contracts = require("../compiled.json");
    // 启动l2，l2的provider内还包含一个l1的provider(用ganache启动的本地l1) 为什么？withdraw需要在l1上转账，发送交易/状态数据到l1的inbox合约
    App.provider = new ArbProvider(
      "http://localhost:1235",
      contracts,
      new ethers.providers.Web3Provider(web3Provider)
    );
    return App.initContract();
  },

  initContract: async function() {
    var network = await App.provider.getNetwork();
    const myContract = require("../build/contracts/myContract.json");
    let address = myContract.networks[network.chainId.toString()].address;
    let myContractContractRaw = new ethers.Contract(
      address,
      myContract.abi,
      App.provider
    );
    let wallet = await App.provider.getSigner(0);
    App.contracts.myContract = myContractContractRaw.connect(wallet); // eslint-disable-line require-atomic-updates
    App.account = await wallet.getAddress(); // eslint-disable-line require-atomic-updates
    let greeting = App.contracts.myContract.sayHelloWorld();
    console.log("===== the greeting is", greeting);
  }
};

$(function() {
  $(window).on("load", function() {
    App.init();
  });
});
