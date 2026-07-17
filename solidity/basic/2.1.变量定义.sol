// 指定编译器版本，^0.8.20 代表0.8.20 ~ 0.9.0之间均可编译
// 0.8.x 内置溢出/下溢检查，无需引入SafeMath
pragma solidity ^0.8.0;


contract varDefine{

    int256 public AuthAge; //定义作者年龄
    byte32 public AuthHash; //定义作者hash值
    string public AuthName; //定义作者姓名
    uint256 AuthSal;    //定义作者薪水

    constructor(int256 _AuthAge, string memory _AuthName, uint256 _AuthSal){
        AuthName = _AuthName;
        AuthAge = _AuthAge;
        AuthSal = _AuthSal;
        // keccak256 以太坊使用的椭圆曲线算法，用来计算hash值，只接收一个参数，因此需要经过内置的abi.encode将多个参数打包成一个参数，
        // keccak256 返回值是 bytes32 的定长数组，也就是hash值
        AuthHash = keccak256(abi.encode(AuthName, AuthAge, AuthSal));
    }

}