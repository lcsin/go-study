package main

import "fmt"

func main() {
	//defer f1()
	//defer f2()
	//defer f3()

	n := 5
	defer f4(n)
	n++
	fmt.Println(n)
}

func f1() {
	fmt.Println("f1...")
}

func f2() {
	fmt.Println("f2...")
}

func f3() {
	fmt.Println("f3...")
}

func f4(n int) {
	fmt.Println(n)
}
