// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Reverser{
    //反转一个字符串。输入 "abcde"，输出 "edcba"
    function reversStr(string memory input)public  pure returns (string memory){
         bytes memory temp = bytes(input); //将字符串转为字符数组
         uint i = 0;
         uint j = temp.length -1 ;
         while(i<temp.length/2 && i != j) {
             bytes1 a = temp[i];
             temp[i] = temp[j];
             temp[j] =a; 
             ++i;
             --j;
         }
         return string(temp);
         
    }
}