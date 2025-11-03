package main

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
import "fmt"

func singleNumber(nums []int) int {
	// 创建一个 map 来记录每个数字出现的次数
	countMap := make(map[int]int)

	// 第一次遍历：统计每个数字出现的次数
	for _, num := range nums {
		countMap[num]++
	}

	// 第二次遍历：找出出现次数为1的数字
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return -1 // 如果没有找到，返回-1（根据题目保证会找到）
}

// 更高效的解法：使用异或运算
func singleNumberXOR(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	testCases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
	}

	fmt.Println("使用 Map 方法:")
	for _, test := range testCases {
		result := singleNumber(test)
		fmt.Printf("数组 %v 中只出现一次的数字是: %d\n", test, result)
	}

	fmt.Println("\n使用 XOR 方法:")
	for _, test := range testCases {
		result := singleNumberXOR(test)
		fmt.Printf("数组 %v 中只出现一次的数字是: %d\n", test, result)
	}
}
