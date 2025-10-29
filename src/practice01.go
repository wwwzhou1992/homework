package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func JudgeTwice(arr []string) string {
	//声明一个Map<int,int>记录出现的数组中的内容和次数
	timeVal := make(map[string]int)
	var value string
	for _, i2 := range arr {
		if timeVal[i2] > 0 {
			timeVal[i2]++
		} else {
			timeVal[i2] = 1
		}
	}
	for k, v := range timeVal {
		if v == 1 {
			fmt.Println(k)
			value = k
		}
	}
	return value
}