// 指定编译器版本，^0.8.20 代表0.8.20 ~ 0.9.0之间均可编译
// 0.8.x 内置溢出/下溢检查，无需引入SafeMath
pragma solidity ^0.8.0;

contract func_demo1 {

    address admin;
    uint256 count;

    function setCount(uint256 _count) public {
        count = _count;
    }

    function setCount2(uint256 _count) external {
        count = _count;
    }

    function withDraw() payable {

    }

    function getCount() public view returns (uint256){
        return count;
    }
}
