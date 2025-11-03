package main

/*
*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*
*/
import "fmt"

// 定义一个函数，接收整数指针作为参数
func increaseByTen(ptr *int) {
	// 通过指针修改指向的值
	*ptr = *ptr + 10
}

func main() {
	// 声明并初始化一个整数变量
	num := 5
	fmt.Printf("原始值: %d\n", num)

	// 获取变量的地址（指针）并传递给函数
	increaseByTen(&num)

	// 输出修改后的值
	fmt.Printf("修改后的值: %d\n", num)

	// 额外演示：直接使用指针变量
	var anotherNum int = 20
	var ptr *int = &anotherNum

	fmt.Printf("\n另一个示例:\n")
	fmt.Printf("原始值: %d\n", anotherNum)
	fmt.Printf("指针指向的值: %d\n", *ptr)

	increaseByTen(ptr)
	fmt.Printf("修改后的值: %d\n", anotherNum)
	fmt.Printf("指针指向的新值: %d\n", *ptr)
}
