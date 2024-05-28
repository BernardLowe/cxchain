// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/** 
 * @title Ballot
 * @dev Implements voting process along with vote delegation
 */
contract TContract {
    uint256 num;
    uint256 ba;

    constructor() {
        num = 1;
    }

    function setNum(uint256 _num) external {
        num = num;
    }

    function add() external {
        num++;
    }

}