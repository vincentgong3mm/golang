package main

import (
	"fmt"

	"github.com/vincentgong3mm/golang/study2019/exch"
	//"github.com/vincentgong3mm/golang/study2019/exregexp"
	rt "github.com/vincentgong3mm/golang/study2019/routine"

	"sync"
)

func main() {
	fmt.Println("hi..")

	exch.TestChannel()
	//exregexp.TestRegexp()

	var wait sync.WaitGroup
	wait.Add(2)

	go rt.Say(&wait, "11111")
	go rt.Say(&wait, "22222")

	/*
		go func(s string) {
			defer wait.Done()
			for i := 0; i < 5; i++ {
				fmt.Println(s)
				time.Sleep(500 * time.Millisecond)
			}
		}("111")

		go func(s string) {
			defer wait.Done()
			for i := 0; i < 5; i++ {
				fmt.Println(s)
				time.Sleep(500 * time.Millisecond)
			}
		}("222")
	*/

	wait.Wait()

}
