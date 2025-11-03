package main

/*
*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。



示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。
示例 3：

输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。


提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits 不包含任何前导 0。
*/
import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始加1
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}

		// 有进位，当前位设为0，继续处理前一位
		digits[i] = 0
	}

	// 如果所有位都有进位（如999+1=1000），需要在最前面添加1
	return append([]int{1}, digits...)
}

func main() {
	testCases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9},
		{0},
		{1, 9, 9},
	}

	for _, test := range testCases {
		result := plusOne(test)
		fmt.Printf("输入: %v\n", test)
		fmt.Printf("输出: %v\n", result)
		fmt.Println("---")
	}
}
