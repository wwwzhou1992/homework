package main

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。

nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略
*/
import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用双指针法
	i := 0                           // 慢指针，指向当前唯一元素的位置
	for j := 1; j < len(nums); j++ { // 快指针，遍历整个数组
		if nums[j] != nums[i] {
			// 找到新的唯一元素，移动到慢指针的下一个位置
			i++
			nums[i] = nums[j]
		}
	}

	// 返回唯一元素的数量（索引+1）
	return i + 1
}

func main() {
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 2, 3},
		{1, 1, 1},
		{},
	}

	for _, test := range testCases {
		fmt.Printf("原数组: %v\n", test)
		k := removeDuplicates(test)
		fmt.Printf("去重后长度: %d\n", k)
		fmt.Printf("前 %d 个元素: %v\n", k, test[:k])
		fmt.Println("---")
	}
}
