package main

import "fmt"

/*
高阶函数：根据go语言的数据类型的特点，可以将一个函数作为另一个函数的参数和返回值
    fun1(),fun2()
    将fun1函数作为了fun2这个函数的参数。
            fun2函数：就叫高阶函数
                接收了一个函数作为参数的函数，高阶函数

            fun1函数：回调函数
                作为另一个函数的参数的函数，叫做回调函数。
*/
func main() {
	fmt.Println("oper add: ", oper(1, 2, add))
	fmt.Println("oper sub: ", oper(1, 2, sub))
	fmt.Println("oper div: ", oper(2, 2, func(a, b int) int {
		return a / b
	}))
	fmt.Println("oper mul: ", oper(1, 2, func(a, b int) int {
		return a * b
	}))
}

// add 和操作
func add(a, b int) int {
	return a + b
}

// sub 差操作
func sub(a, b int) int {
	return a - b
}

// oper 运算操作
func oper(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}
