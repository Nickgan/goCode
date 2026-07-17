// 指定编译器版本，^0.8.20 代表0.8.20 ~ 0.9.0之间均可编译
// 0.8.x 内置溢出/下溢检查，无需引入SafeMath
pragma solidity ^0.8.0;

contract  HelloWorld {

    //状态变量：存储在链上的字符串
    string  public helloText;

    //构造函数：合约部署时仅执行一次
    constructor() {
        helloText = "Hello World";
    }

    // 读取字符串view函数（不消耗gas,只读）
    function getHello() public view returns(string){
        return helloText;
    }

    // 修改字符串函数（消耗gas,写入链
    function setHello(string  memory _helloText) public {
        helloText = _helloText;
    }
}
