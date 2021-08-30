package main

import "fmt"

func main() {
	// 声明一个int类型的切片
	var slice1 []int
	fmt.Printf("slice1: %T,%v\n", slice1, slice1)

	// 声明并初始化一个string类型的切片
	var slice2 = []int{1, 2, 3}
	fmt.Printf("slice2: %T,%v\n", slice2, slice2)

	// 使用make函数创建切片
	slice3 := make([]int, 2, 5)
	fmt.Printf("slice3: %T,%v\n", slice3, slice3)

	// 使用切片操作创建切片(坐闭右开)
	slice4 := slice2[0:2]
	fmt.Printf("slice3: %T,%v\n", slice4, slice4)

	// slice只是底层数组的一个表示，对slice所坐的任何修改都将反映在底层数组中
	// 当多个切片共享相同的底层数组时，对每个切片的修改都会对底层数组进行修改，继而影响其他共享该底层数组的切片
	slice5 := slice2
	slice5[2] = 10086
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice5: ", slice5)

	// 切片的长度len(): 切片中元素的数量
	// 切片的容量cap(): 底层数组中元素的数量
	fmt.Println("slice3 len: ", len(slice3))
	fmt.Println("slice3 cap: ", cap(slice3))

	// 空切片: 一个切片在未初始化之前默认为nil，长度为0
	fmt.Printf("slice1 val: %v\n", slice1)
	fmt.Printf("slice1 len: %d\n", len(slice1))
	fmt.Printf("slice1 == nil: %v\n", slice1 == nil)

	// append()函数: 向slice里面追加一个或者多个元素，返回一个一样类型的slice，允许往空slice中追加元素
	slice1 = append(slice1, 3, 4, 5)
	fmt.Println("newSlice1: ", slice1)

	// copy()函数: 从源slice中复制元素到目标slice，并且返回复制的元素的个数，要求目标slice必须初始化
	// copy()方法不会在源slice和目标slice中建立联系
	slice6 := make([]int, 2, 6)
	count := copy(slice6, slice2)
	fmt.Printf("拷贝的元素个数：%d\n", count)
	fmt.Printf("拷贝后的目标slice：%v\n", slice6)
}
