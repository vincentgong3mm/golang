package main

import "fmt"

type MyInter interface {
	PrintLog()
	PrintName()
}

type MyDataA struct {
	RealName  string
	RealValue int
}

type MyDataB struct {
	TempName   string
	TempValue  int
	TempValut2 int
}

func (r *MyDataA) PrintLog() {
	fmt.Println("MyDataA")
}

func (r *MyDataA) PrintName() {
	fmt.Println(r.RealName)
}

func PrintData(d ...MyInter) {
	for _, v := range d {
		v.PrintLog()
		v.PrintName()
	}
}

func main() {
	fmt.Println("vim-go")

	a := MyDataA{RealName: "aaaa", RealValue: 10}

	PrintData(&a)
}
