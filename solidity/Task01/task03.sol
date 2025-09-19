// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

//合并两个有序数组 (Merge Sorted Array) - Solidity 实现
//由于 Solidity 主要用于区块链智能合约开发，处理数组操作需要特别注意 Gas 消耗和效率问题。


contract ArrayMerger{
    function mergeSortedArrays(uint[] memory arr1,uint[] memory arr2)
    public pure returns(uint[] memory){
        uint totalLength = arr1.length + arr2.length;
        uint[] memory merged = new uint[](totalLength);

        //在 Solidity 中，变量声明语法与其他语言（如 JavaScript 或 C++）不同。不能在同一行使用逗号分隔声明多个变量。必须分别声明每个变量：
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] <= arr2[j]) {
                merged[k] = arr1[i];
                i++;
            } else {
                merged[k] = arr2[j];
                j++;
            }
            //结果数组的下角标
            k++;
        }
        //此时两个数组不可能同时有值
        //如果arr1 还有剩余元素，全部添加近结果中
        while(i < arr1.length){
            merged[k] = arr1[i];
            i++;
            k++;
        }

        //如果arr2还有剩余元素则全部添加
        while(j < arr2.length){
            merged[k] = arr2[j];
            j++;
            k++;
        }
            return merged; 
        }



function testMerge(uint[] memory arr1, uint[] memory arr2) 
        public 
        pure 
        returns (uint[] memory) 
    {
        return mergeSortedArrays(arr1, arr2);
    }



}