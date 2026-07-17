// 指定编译器版本，^0.8.20 代表0.8.20 ~ 0.9.0之间均可编译
// 0.8.x 内置溢出/下溢检查，无需引入SafeMath
pragma solidity ^0.8.0;


contract function_status_modify_demo {


    uint256 public balance;

    function deposit() public payable {
        require(msg.value > 0, "msg.value 必须是一个大于0的数。");
        balance += msg.value;
    }

    function withdraw(uint256 _amount) public {
        require(_amount <= balance, "余额不足");
        balance -= _amount;
        payable(msg.sender).transfer(_amount);
    }
}
