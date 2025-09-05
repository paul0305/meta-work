// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

contract Andy {



    function romanToInt(string memory s) public pure returns (uint256) {
        //字符串转为bytes数组
        uint length = bytes(s).length;
        uint256 result = 0;
        //定义数组存储罗马数组对应的10进制数
        uint256[256] memory values;
 /*罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000*/
        // 设置每个字符的值
        values[uint8(bytes1('I'))] = 1;
        values[uint8(bytes1('V'))] = 5;
        values[uint8(bytes1('X'))] = 10;
        values[uint8(bytes1('L'))] = 50;
        values[uint8(bytes1('C'))] = 100;
        values[uint8(bytes1('D'))] = 500;
        values[uint8(bytes1('M'))] = 1000;

        for (uint i = 0; i < length; i++) {
            uint8 currentKey = uint8(bytes(s)[i]);
            uint256 currentValue = values[currentKey];
            uint256 nextValue = i + 1 < length ? values[uint8(bytes(s)[i + 1])] : 0;

            if (currentValue < nextValue) {
                result += nextValue - currentValue;
                i++; // 跳过下一个字符
            } else {
                result += currentValue;
            }
        }

        return result;
    }

    //二分查找(Binary Search)
    //在一个有序数组中查找目标值,返回目标值的下标
    function binarySearchRecursive(uint256[] memory array, uint target) public pure  returns (uint256 index) {
        //数组下标lef right middle
        uint256 left = 0;
        uint256 right = array.length - 1;
        uint256 middle;
        //循环查找
        while (left <= right) {
            //中间值
            middle = (left + right) / 2;
            if (array[middle] == target){
                return middle;
            } else if (target > array[middle]) {
                left = middle + 1;
            } else if (target < array[middle]) {
                right = middle - 1;
            }
        }
        //返回未找到值
        return uint256(0);
    }
    

}
