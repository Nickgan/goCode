// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;


contract  EventsDemo {

    // 定义充值事件
    event Deposit(address indexed _from, uint indexed _value, string note);


    // 充值函数
    function deposit() public payable {

        // 触发事件
        emit Deposit(msg.sender, msg.value,"转账测试value 必须大于0value 必须大于0");
    }


    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}