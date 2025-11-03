package main

import "fmt"

/*
*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案
*
*/
func twoSum(nums []int, target int) []int {
	// 创建哈希表存储数字和对应的索引
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算需要的补数
		complement := target - num

		// 检查补数是否在哈希表中
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}

		// 将当前数字和索引存入哈希表
		numMap[num] = i
	}

	return nil // 根据题目假设，总会有一个解
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{0, 4, 3, 0}, 0},
		{[]int{-1, -2, -3, -4, -5}, -8},
	}

	for _, test := range testCases {
		result := twoSum(test.nums, test.target)
		fmt.Printf("nums = %v, target = %d\n", test.nums, test.target)
		fmt.Printf("输出: %v\n", result)
		fmt.Printf("验证: nums[%d] + nums[%d] = %d + %d = %d\n",
			result[0], result[1], test.nums[result[0]], test.nums[result[1]], test.target)
		fmt.Println("---")
	}
}
