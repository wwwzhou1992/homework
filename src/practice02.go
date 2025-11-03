/*
*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
package main

import "fmt"

func isValid(s string) bool {
	// 使用切片模拟栈
	stack := make([]byte, 0)

	// 创建括号映射表
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for i := 0; i < len(s); i++ {
		char := s[i]

		// 如果是右括号
		if matchingLeft, isRight := pairs[char]; isRight {
			// 检查栈是否为空或者栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != matchingLeft {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	// 最后栈应该为空才表示所有括号都正确匹配
	return len(stack) == 0
}

// 将 main 改为 testBrackets 避免冲突
func testBrackets() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("程序发生异常: %v\n", r)
		}
	}()

	testCases := []struct {
		input    string
		expected bool
	}{
		{"()", true},     // 应该通过
		{"()[]{}", true}, // 应该通过
		{"(]", false},    // 应该失败
		{"([)]", false},  // 应该失败
		{"{[]}", true},   // 应该通过
		{"]", false},     // 应该失败
	}

	for _, test := range testCases {
		result := isValid(test.input)
		status := "FAIL"
		if result == test.expected {
			status = "PASS"
		}
		fmt.Printf("[%s] isValid(%q) = %t (expected %t)\n", status, test.input, result, test.expected)
	}
}

func main() {
	// 调用括号验证测试
	testBrackets()

	// 您原来的 main.go 代码也可以放在这里
}
