package main

import "fmt"

func main() {
	go f1()

}

func f1() {
	fmt.Println("hello,world!")
}
