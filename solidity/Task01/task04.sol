// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


contract BinarySearch{

    function binarySearch(uint256[] calldata arr,uint256 target) external pure returns (uint256){
        //处理空数组情况
        if(arr.length == 0){
            //返回uint256 类型能表示的最大数值
            return type(uint256).max;
        }

        //初始化左边界和右边界
        uint256 left = 0;
        uint256 right = arr.length -1;

        //当左边界小于等于有边界时继续搜索
        while (left <= right){
            //计算中间索引（防止溢出）
            //防止在计算中间索引时，左边界 left 和右边界 right 的直接相加 (left + right) 可能导致的整数溢出。
            /*
            (right - left)：先计算区间长度。这个值通常不会很大（尤其是在搜索的后期），远小于 type(uint256).max，因此它本身溢出的风险极低。
            (right - left) / 2：计算出区间长度的一半。
            left + ...：将左边界加上一半的长度，得到精确的中间点。
            */
            uint256 mid = left + (right - left) / 2;

            //检查中间值是否等于目标值
            if(arr[mid] == target){
                return mid;
            }

            //如果中间值小于目标值，说明目标在右半部分
            else if (arr[mid] < target){
                left = mid + 1;
            }//如果中间值大于目标值，说明目标在左半部分
            else{
                right  = mid -1 ;
            }
        }

        //未找到目标，返回最大值表示失败
        return type(uint256).max;
    }
}