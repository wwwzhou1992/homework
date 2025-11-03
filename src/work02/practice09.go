package main

import "fmt"

/*
*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*
*/
// 方法1：接收切片指针作为参数
func doubleSliceElements(ptr *[]int) {
	// 通过指针访问切片并修改每个元素
	for i := 0; i < len(*ptr); i++ {
		(*ptr)[i] = (*ptr)[i] * 2
	}
}

// 方法2：使用更简洁的语法（推荐）
func doubleSliceElementsV2(slicePtr *[]int) {
	// Go允许直接对切片指针使用索引操作
	slice := *slicePtr // 解引用获取切片
	for i := range slice {
		slice[i] = slice[i] * 2
	}
}

// 方法3：如果不需要改变切片本身（长度/容量），可以直接传递切片
func doubleSliceDirectly(slice []int) {
	for i := range slice {
		slice[i] = slice[i] * 2
	}
}

func main() {
	// 示例1：使用切片指针
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始切片: %v\n", numbers)

	doubleSliceElements(&numbers)
	fmt.Printf("方法1处理后: %v\n", numbers)

	// 示例2：重置切片并使用第二种方法
	numbers = []int{1, 2, 3, 4, 5}
	doubleSliceElementsV2(&numbers)
	fmt.Printf("方法2处理后: %v\n", numbers)

	// 示例3：演示直接传递切片（更常用）
	numbers = []int{1, 2, 3, 4, 5}
	doubleSliceDirectly(numbers)
	fmt.Printf("直接传递切片: %v\n", numbers)

	// 示例4：演示为什么有时需要传递切片指针
	fmt.Println("\n--- 切片追加场景 ---")
	originalSlice := []int{1, 2, 3}
	fmt.Printf("追加前 - 切片: %v, 长度: %d, 容量: %d\n",
		originalSlice, len(originalSlice), cap(originalSlice))

	appendToSlice(originalSlice)
	fmt.Printf("追加后(无指针) - 切片: %v, 长度: %d, 容量: %d\n",
		originalSlice, len(originalSlice), cap(originalSlice))

	// 使用指针来真正修改切片
	appendToSliceWithPointer(&originalSlice)
	fmt.Printf("追加后(有指针) - 切片: %v, 长度: %d, 容量: %d\n",
		originalSlice, len(originalSlice), cap(originalSlice))
}

// 这个函数无法修改外部的切片，因为append可能返回新的切片
func appendToSlice(slice []int) {
	slice = append(slice, 4, 5, 6)
	fmt.Printf("  函数内部切片: %v, 长度: %d\n", slice, len(slice))
}

// 这个函数可以真正修改外部的切片
func appendToSliceWithPointer(slicePtr *[]int) {
	*slicePtr = append(*slicePtr, 4, 5, 6)
	fmt.Printf("  函数内部切片: %v, 长度: %d\n", *slicePtr, len(*slicePtr))
}
