package routine_test

import (
	"fmt"
	rt "github.com/vincentgong3mm/golang/study2019/routine"
	"sync"
	"testing"
)

func TestRoutine(t *testing.T) {
	fmt.Println("TestRouine")
	var wait sync.WaitGroup
	wait.Add(2)

	go rt.Say(&wait, "11111")
	go rt.Say(&wait, "22222")

	wait.Wait()
}

func TestMin(t *testing.T) {
	m := rt.ParallelMin([]int{100, 24, 30, 7, 9}, 2)
	fmt.Println(m)
}
