// 指定编译器版本，^0.8.20 代表0.8.20 ~ 0.9.0之间均可编译
// 0.8.x 内置溢出/下溢检查，无需引入SafeMath
pragma solidity ^0.8.0;

contract mapping_demo {

    //定义address 和name 的mapping映射
    mapping(address => string) public addr_name;

    constructor(){
        addr_name[msg.sender] = "yekai";
    }

    // 设置地址对应的姓名
    function setName(string memory _name) public {
        addr_name[msg.sender] = _name;
    }
}
