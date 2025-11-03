package main

/*
*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*
*/
import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// 将当前前缀与每个字符串进行比较，逐步缩短前缀
		for j := 0; j < len(prefix); j++ {
			// 如果当前字符串长度不够，或者字符不匹配
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
		// 如果前缀已经为空，直接返回
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// 另一种解法：纵向扫描
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 遍历第一个字符串的每个字符
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		// 检查其他字符串在相同位置的字符
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度不够或字符不匹配
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"", "b"},
		{"a"},
		{"abc", "abc", "abc"},
		{"ab", "a"},
	}

	fmt.Println("方法1结果:")
	for _, test := range testCases {
		result := longestCommonPrefix(test)
		fmt.Printf("字符串数组 %v 的最长公共前缀是: \"%s\"\n", test, result)
	}

	fmt.Println("\n方法2结果:")
	for _, test := range testCases {
		result := longestCommonPrefix2(test)
		fmt.Printf("字符串数组 %v 的最长公共前缀是: \"%s\"\n", test, result)
	}
}
