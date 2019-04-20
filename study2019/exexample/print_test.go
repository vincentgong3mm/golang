package main

import "fmt"

func ExamplePrint() {
	fmt.Println("Example---")

	a := 10
	fmt.Println("aaa")

	b := 40
	fmt.Println("bbb")

	c := a + b
	fmt.Println(c)
	// Output:
	// Example---
	// aaa
	// ccc
	// 30
}
