package exch

import "fmt"

func init() {
	fmt.Println("call exch init function")
}

// TestChannel :
func TestChannel() {
	fmt.Println("tsetChannel")

	testLocalFunction()
}

func testLocalFunction() {
	fmt.Println("local Function")
}
