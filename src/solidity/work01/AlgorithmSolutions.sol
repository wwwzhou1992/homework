// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AlgorithmSolutions {
    // ========== 2. 反转字符串 (Reverse String) ==========
    function reverseString(string memory str) public pure returns (string memory) {
        bytes memory strBytes = bytes(str);
        uint256 length = strBytes.length;
        
        // 使用双指针法反转
        for (uint256 i = 0; i < length / 2; i++) {
            bytes1 temp = strBytes[i];
            strBytes[i] = strBytes[length - 1 - i];
            strBytes[length - 1 - i] = temp;
        }
        
        return string(strBytes);
    }
    
    // ========== 3. 罗马数字转整数 (Roman to Integer) ==========
    function romanToInt(string memory s) public pure returns (uint256) {
        bytes memory roman = bytes(s);
        uint256 length = roman.length;
        uint256 result = 0;
        uint256 prev = 0;
        
        // 从右向左遍历
        for (uint256 i = 0; i < length; i++) {
            uint256 current = getRomanValue(roman[i]);
            result += current;
            
            // 如果当前值大于前一个值，需要减去2倍的prev（因为之前已经加过了）
            if (current > prev) {
                result -= 2 * prev;
            }
            prev = current;
        }
        
        return result;
    }
    
    // 辅助函数：获取罗马数字字符对应的值
    function getRomanValue(bytes1 romanChar) private pure returns (uint256) {
        // 比较 ASCII 码
        if (romanChar == 0x49) return 1; // 'I'
        if (romanChar == 0x56) return 5; // 'V'
        if (romanChar == 0x58) return 10; // 'X'
        if (romanChar == 0x4C) return 50; // 'L'
        if (romanChar == 0x43) return 100; // 'C'
        if (romanChar == 0x44) return 500; // 'D'
        if (romanChar == 0x4D) return 1000; // 'M'
        return 0;
    }
    
    // ========== 4. 整数转罗马数字 (Integer to Roman) ==========
    function intToRoman(uint256 num) public pure returns (string memory) {
        require(num >= 1 && num <= 3999, "Number must be between 1 and 3999");
        
        // 使用查表法
        string[13] memory romanNumerals = [
            "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"
        ];
        
        // 显式指定 uint256 类型
        uint256[13] memory values = [
            uint256(1000), 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1
        ];
        
        bytes memory result;
        
        for (uint256 i = 0; i < romanNumerals.length; i++) {
            while (num >= values[i]) {
                bytes memory roman = bytes(romanNumerals[i]);
                for (uint256 j = 0; j < roman.length; j++) {
                    result = abi.encodePacked(result, roman[j]);
                }
                num -= values[i];
            }
        }
        
        return string(result);
    }
    
    // 整数转罗马数字的另一种实现（避免类型转换问题）
    function intToRomanV2(uint256 num) public pure returns (string memory) {
        require(num >= 1 && num <= 3999, "Number must be between 1 and 3999");
        
        // 使用固定大小的数组
        string[13] memory romanNumerals = [
            "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"
        ];
        
        uint256[] memory values = new uint256[](13);
        values[0] = 1000;
        values[1] = 900;
        values[2] = 500;
        values[3] = 400;
        values[4] = 100;
        values[5] = 90;
        values[6] = 50;
        values[7] = 40;
        values[8] = 10;
        values[9] = 9;
        values[10] = 5;
        values[11] = 4;
        values[12] = 1;
        
        bytes memory result;
        
        for (uint256 i = 0; i < romanNumerals.length; i++) {
            while (num >= values[i]) {
                bytes memory roman = bytes(romanNumerals[i]);
                for (uint256 j = 0; j < roman.length; j++) {
                    result = abi.encodePacked(result, roman[j]);
                }
                num -= values[i];
            }
        }
        
        return string(result);
    }
    
    // ========== 5. 合并两个有序数组 (Merge Sorted Array) ==========
    function mergeSortedArrays(
        uint256[] memory nums1,
        uint256[] memory nums2
    ) public pure returns (uint256[] memory) {
        uint256 m = nums1.length;
        uint256 n = nums2.length;
        uint256[] memory merged = new uint256[](m + n);
        
        uint256 i = 0; // nums1 指针
        uint256 j = 0; // nums2 指针
        uint256 k = 0; // merged 指针
        
        // 合并两个有序数组
        while (i < m && j < n) {
            if (nums1[i] <= nums2[j]) {
                merged[k] = nums1[i];
                i++;
            } else {
                merged[k] = nums2[j];
                j++;
            }
            k++;
        }
        
        // 将剩余元素添加到 merged
        while (i < m) {
            merged[k] = nums1[i];
            i++;
            k++;
        }
        
        while (j < n) {
            merged[k] = nums2[j];
            j++;
            k++;
        }
        
        return merged;
    }
    
    // ========== 6. 二分查找 (Binary Search) ==========
    function binarySearch(
        uint256[] memory nums,
        uint256 target
    ) public pure returns (int256) {
        if (nums.length == 0) return -1;
        
        uint256 left = 0;
        uint256 right = nums.length - 1;
        
        while (left <= right) {
            uint256 mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                return int256(mid);
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                // 防止下溢
                if (mid == 0) break;
                right = mid - 1;
            }
        }
        
        return -1; // 没有找到
    }
    
    // 二分查找的另一个版本
    function binarySearchV2(
        uint256[] memory nums,
        uint256 target
    ) public pure returns (int256) {
        uint256 left = 0;
        uint256 right = nums.length;
        
        while (left < right) {
            uint256 mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                return int256(mid);
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        
        return -1;
    }
    



    
    // ========== 罗马数字转整数（简化版） ==========
    function romanToIntSimple(string memory s) public pure returns (uint256) {
        bytes memory roman = bytes(s);
        uint256 length = roman.length;
        uint256 result = 0;
        uint256 prev = 0;
        
        for (uint256 i = 0; i < length; i++) {
            uint256 current = _getRomanValue(roman[i]);
            result += current;
            
            if (current > prev) {
                result -= 2 * prev;
            }
            prev = current;
        }
        
        return result;
    }
    
    function _getRomanValue(bytes1 c) private pure returns (uint256) {
        // 直接比较字符
        //反转字符串："abcde" → "edcba"

//罗马数字转整数："MCMXCIV" → 1994

//整数转罗马数字：1994 → "MCMXCIV"，2024 → "MMXXIV"

//合并数组：[1,3,5] + [2,4,6] → [1,2,3,4,5,6]

//二分查找：在 [1,3,5,7,9,11] 中查找 7 → 索引 3
        if (c == 'I') return 1;
        if (c == 'V') return 5;
        if (c == 'X') return 10;
        if (c == 'L') return 50;
        if (c == 'C') return 100;
        if (c == 'D') return 500;
        if (c == 'M') return 1000;
        return 0;
    }
}