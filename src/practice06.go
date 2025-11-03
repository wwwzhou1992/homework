package main

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
*
*/
import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 结果数组
	result := [][]int{}
	// 当前合并区间
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		// 如果当前区间与下一个区间有重叠
		if current[1] >= intervals[i][0] {
			// 合并区间，取结束位置的较大值
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			// 没有重叠，将当前区间加入结果
			result = append(result, current)
			// 更新当前区间为下一个区间
			current = intervals[i]
		}
	}

	// 添加最后一个区间
	result = append(result, current)

	return result
}

func main() {
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {2, 3}},
		{{1, 4}, {0, 0}},
		{},
		{{1, 4}},
	}

	for _, test := range testCases {
		result := merge(test)
		fmt.Printf("输入: %v\n", test)
		fmt.Printf("输出: %v\n", result)
		fmt.Println("---")
	}
}
