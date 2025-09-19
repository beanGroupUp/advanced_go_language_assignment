// SPDX-License-Identifier: MIT
// 使用MIT许可证
pragma solidity ^0.8.0;
// 指定Solidity编译器版本为0.8.0及以上

contract RomanToInteger {
    // 定义一个名为RomanToInteger的合约
    
    function romanToInt(string memory s) public pure returns (uint256) {
        // 公开的纯函数，将罗马数字字符串转换为整数
        // 参数s: 内存中的罗马数字字符串
        // 返回值: 转换后的整数值
        
        bytes memory sBytes = bytes(s);
        // 将输入的字符串转换为字节数组，以便逐个字符处理
        
        uint256 result = 0;
        // 初始化结果为0，用于累加计算最终值
        
        for (uint256 i = 0; i < sBytes.length; i++) {
            // 遍历字节数组中的每个字符
            
            uint256 currentValue = getValue(sBytes[i]);
            // 获取当前字符对应的数值
            
            if (i < sBytes.length - 1) {
                // 如果不是最后一个字符，检查下一个字符
                
                uint256 nextValue = getValue(sBytes[i + 1]);
                // 获取下一个字符对应的数值
                
                if (currentValue < nextValue) {
                    // 如果当前字符值小于下一个字符值（特殊情况，如IV、IX等）
                    
                    result -= currentValue;
                    // 从结果中减去当前值（因为这是减法表示法）
                } else {
                    // 否则（正常情况，当前字符值大于或等于下一个字符值）
                    
                    result += currentValue;
                    // 向结果中添加当前值
                }
            } else {
                // 如果是最后一个字符
                
                result += currentValue;
                // 直接向结果中添加当前值
            }
        }
        return result;
        // 返回最终的计算结果
    }
    
    function getValue(bytes1 c) private pure returns (uint256) {
        // 私有辅助函数，将单个罗马数字字符转换为对应的数值
        // 参数c: 字节类型的字符
        // 返回值: 字符对应的数值
        
        if (c == 'I') return 1;
        // 如果是字符'I'，返回1
        if (c == 'V') return 5;
        // 如果是字符'V'，返回5
        if (c == 'X') return 10;
        // 如果是字符'X'，返回10
        if (c == 'L') return 50;
        // 如果是字符'L'，返回50
        if (c == 'C') return 100;
        // 如果是字符'C'，返回100
        if (c == 'D') return 500;
        // 如果是字符'D'，返回500
        if (c == 'M') return 1000;
        // 如果是字符'M'，返回1000
        return 0;
        // 如果不是有效的罗马数字字符，返回0（根据题目保证，这种情况不会发生）
    }
}