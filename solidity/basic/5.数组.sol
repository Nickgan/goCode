// 指定编译器版本，^0.8.20 代表0.8.20 ~ 0.9.0之间均可编译
// 0.8.x 内置溢出/下溢检查，无需引入SafeMath
pragma solidity ^0.8.0;

contract arry_demo {

    string[5] public names;
    uint256[] public ages;

    constructor(){

        // 修改定长数组的元素
        names[0] = "yekai";

        // 动态数组追加元素
        ages.push(18);
    }

    // 获取数组长度方法
    function getLength() public view returns (uint256, uint256){
        return (names.length, ages.length);
    }
}
