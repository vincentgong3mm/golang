package calc_test

import (
	"calc"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	s := calc.Sum(1, 2, 3)

	if s != 6 {
		t.Error("Wrong result")
	}

	fmt.Println(s)
}
