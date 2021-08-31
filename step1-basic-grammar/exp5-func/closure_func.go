package main

import "fmt"

func main() {
	res1 := increment()      //res1 = fun
	fmt.Printf("%T\n", res1) //func() int
	fmt.Println(res1)

	v1 := res1()
	fmt.Println(v1) //1
	v2 := res1()
	fmt.Println(v2) //2
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment()
	fmt.Println(res2)
	v3 := res2()
	fmt.Println(v3) //1
	fmt.Println(res2())

	fmt.Println(res1())
}

func increment() func() int { //外层函数
	//1.定义了一个局部变量
	i := 0
	//2.定义了一个匿名函数，给变量自增并返回
	fun := func() int { //内层函数
		i++
		return i
	}
	//3.返回该匿名函数
	return fun
}
