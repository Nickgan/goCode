// 指定编译器版本，^0.8.20 代表0.8.20 ~ 0.9.0之间均可编译
// 0.8.x 内置溢出/下溢检查，无需引入SafeMath
pragma solidity ^0.8.0;

contract localObj {

    // 状态变量，address类型
    address public admin;

    // hash值
    bytes32  public hash;

    // 部署时后产生的随机数
    uint256  public randnum;

    // 构造函数
    constructor() {
        admin = msg.sender; //msg.sender是调用者
        hash = blockhash(0);    // 返回0块的hash值
        // 利用时间戳，调用者账号，hash值共同模拟随机值，得到一个100以内的数
        randnum = uint256(keccak256(abi.encode(now, msg.sender, hash))) % 100;
    }

}