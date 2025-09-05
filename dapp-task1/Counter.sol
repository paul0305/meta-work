// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

//0xE1957595819648F74B45F653FD6eEc40E082f4F5
//0xE1957595819648F74B45F653FD6eEc40E082f4F5
contract Counter{

    uint256 public nb;

    function addNb()public  returns (uint256){
        nb += 1;
        return nb;
    }

    function getNb() public view returns(uint256){
        return nb;
    }

}
