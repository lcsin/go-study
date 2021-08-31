package main

import "fmt"

func main() {
	res1, res2 := func2(2, 3)
	fmt.Println(res1, res2)

	f := func1
	fmt.Printf("func1: %T\n", func1)
	fmt.Printf("f: %T\n", f)
	fmt.Println(f(1, 2))
}

func func1(a, b int) int {
	return a + b
}

func func2(a, b int) (res1 int, res2 int) {
	res1 = a + b
	res2 = a - b
	return res1, res2
}
