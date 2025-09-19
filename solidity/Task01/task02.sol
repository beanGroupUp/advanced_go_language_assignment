// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title RomanNumeralsConverter
 * @dev 将整数转换为罗马数字的合约
 * @notice 支持1到3999之间的整数转换
 */
contract RomanNumeralsConverter {
    
    /**
     * @dev 将整数转换为罗马数字
     * @param num 需要转换的整数，必须在1到3999之间
     * @return 罗马数字字符串
     */
    function toRoman(uint256 num) public pure returns (string memory) {
        // 验证输入范围
        require(num >= 1 && num <= 3999, "Number out of range (1-3999)");
        
        // 定义罗马数字的值和对应的符号
        //在 Solidity 中，数组字面量的类型是由其第一个元素的类型决定的。当我们写 [1000, 900, 500, ...] 时，Solidity 编译器会根据第一个元素 1000 的类型来推断整个数组的类型。
        uint256[13] memory values = [uint256(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        
        string[13] memory symbols = [
            "M", "CM", "D", "CD",
            "C", "XC", "L", "XL",
            "X", "IX", "V", "IV",
            "I"
        ];
        
        // 初始化结果字符串
        string memory result = "";
        
        // 遍历所有可能的符号
        for (uint256 i = 0; i < values.length; i++) {
            // 当当前值可以被减去时
            while (num >= values[i]) {
                num -= values[i];
                //这行代码是 Solidity 中字符串拼接的标准方法。让我详细解释每个部分：
                result = string(abi.encodePacked(result, symbols[i]));
            }
        }
        
        return result;
    }
    
    /**
     * @dev 测试函数，验证转换是否正确
     * @param num 测试数字
     * @return 罗马数字字符串
     */
    function testConversion(uint256 num) public pure returns (string memory) {
        return toRoman(num);
    }
}