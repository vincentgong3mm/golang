package calc_test

import (
	"fmt"
	"github.com/vincentgong3mm/golang/study2019/calc"
	"testing"
)

func TestSum(t *testing.T) {
	s := calc.Sum(1, 2, 3)

	fmt.Println("aaaa")

	if s != 6 {
		t.Error("Wrong result")
	}

	fmt.Println("bbbb")

	fmt.Println(s)
}
