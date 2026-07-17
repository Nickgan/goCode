// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ModifierDemo {

    bool public status;

    // 自定义修饰符
    modifier  myModify (){
        require(status == true, "status must be true");
        _;
    }


    function getStatus() public view myModify returns(bool){
        return status;
    }

}



