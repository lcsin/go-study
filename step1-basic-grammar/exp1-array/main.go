package main

import "fmt"

func main() {
	// 声明一个int类型，大小为3的数组
	var array1 [3]int
	fmt.Printf("array1: %v\n", array1)

	// 声明并初始化数组
	var array2 = [3]int{1, 2, 3}
	fmt.Printf("array2: %v\n", array2)

	// 根据数组的元素个数初始化数组
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("array3: %v\n", array3)

	// 数组的访问
	fmt.Printf("array3[2]的值: %d\n", array3[2])

	// 数组的长度
	fmt.Printf("array3的长度: %d\n", len(array3))

	// for... range 数组遍历
	for idx, val := range array3 {
		fmt.Printf("元素下标：%d, 元素值：%v\n", idx, val)
	}

	// for... len 数组遍历
	for i := 0; i < len(array3); i++ {
		fmt.Printf("元素下标：%d, 元素值：%v\n", i, array3[i])
	}

	// 数组是值类型, 数组间的传递传递的是原始数组的副本
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr2[1] = 10086
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)

	// 数组的大小是类型的一部分，不同长度的数组是不同类型
	var a [3]int
	var b [2]int
	fmt.Printf("数组a类型：%T\n", a)
	fmt.Printf("数组b类型：%T\n", b)

	// 数组间的比较需要相同类型、相同长度才能进行比较，且当元素都一致时，才为true

	fmt.Printf("array1 == array2: %v\n", array1 == array2)
}
