/*
匿名函数的作用：构造高阶函数
*/
package main

import "fmt"

func main() {
	// 定义一个匿名函数，会直接进行调用，且通常只能使用一次，除非将匿名函数赋值给一个变量，通过变量来进行多次调用
	func() {
		fmt.Println("匿名函数被执行...")
	}()

	// 定义一个带参数的匿名函数，在函数体后面的括号内进行传参和调用
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	// 定义一个带返回值的匿名函数，使用一个变量接收返回值
	r := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(r)

	// 匿名函数赋值给一个变量，该变量存储匿名函数的函数体的内存地址，可以通过变量()来多次调用
	f := func(a, b int) int {
		return a + b
	}
	// f是函数的地址
	fmt.Println(f)
	// 可以通过f()多次调用匿名函数
	fmt.Println(f(1, 2))
}
