package main

import "fmt"

func main() {
	// go中的字符串是一个字节切片，采用UTF-8编码(UTF-8编码中，一个中文一般占3个字节)
	str := "hello,world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v ", str[i])
	}
	fmt.Println()
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()

	// 字符串和字节的转换
	slice1 := []byte{65, 66, 67}
	s1 := string(slice1)
	fmt.Println(s1)
	s2 := "XYZ"
	slice2 := []byte(s2)
	fmt.Println(slice2)

	// 字符串不能修改
	//s2[1] = 'B'
}
