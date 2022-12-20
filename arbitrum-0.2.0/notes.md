## prerequisites

```shell
n install 10.16.3

npm install -g truffle@5.0.33
```

install then link python & python3:
(use [this installer](https://blog.shahednasser.com/how-to-install-python-2-on-mac-12-3/) to install 2.7 on mac)

``shell
ln -s -f /usr/local/bin/python3.10 /usr/local/bin/python3
ln -s -f /usr/local/bin/python2.7 /usr/local/bin/python3
```

## deploy on local mock layer1
先用ganache启动一个本地网络，作为l1
把eth跨链桥部署到l1上

```shell
cd packages/arb-bridge-eth
npm i --save-dev
truffle migrate
```

## deploy contract with 

```shell
npm i packages/arb-provider-truffle --save-dev
npm i packages/arb-provider-ethers --save-dev
```

用truffle init创建了一个hello world项目：arbitrum-0.2.0/demos/myContract，用avm compiler进行编译：

```shell
cd demos/myContract
truffle migrate --network arbitrum
```

运行truffle migrate --network arbitrum之后本应该生成compiled.json和contract.ao文件。

1. compile.json文件是编译成evm的代码（truffle本身有的功能）
2. contract.ao文件是编译成avm的代码（arb魔改之后的avm编译器）

arb-provider-truffle是arb对truffle做了一层封装，用truffle编译出compiled.json之后，调用arbc-truffle.py来把compiled.json转换成contract.ao。

但是运行完之后只生成了compiled.json，是因为arbc-truffle.py还没有放入环境变量。

解决：
```shell
cd arbitrum-0.2.0/packages/arb-compiler-evm
pip instal .
```

现在可以直接用arbc-truffle了：
``shell
arbc-truffle arbitrum-0.2.0/demos/myContract/compiled.json arbitrum-0.2.0/demos/myContract/contract.ao 
```
成功生成contract.ao 

// TODO: 详细看arbitrum-0.2.0/packages/arb-compiler-evm 

## wallet.sendTransaction

发送一笔交易的详细过程（先把交易摘要发送到链上，再有validator执行并上传状态根）：
1. 用户给l2发送一笔交易，l2的avm执行这个交易并生成一个新的状态。
2. 管理者把状态变化转换成assertion（交易摘要），质押后上传到l1的inbox合约
3. validator从l1 inbox中监听到交易后在本地进行验证，将生成的新的状态上传到l1 （如果其他的validator发现l1上的状态根和自己本地计算出来的不一样，就发起挑战）

向l2发送交易之前，需要启动l2的validator：

当前的这个版本启动validator的时候需要启动coordinatorServer和followerServer。其中coordinator是将状态根上传到l1的，follower是读取l1上的交易在本地执行，监听是否有争议的。

### 启动validator
