package exdiscoverygo_test

import (
	"fmt"
	"github.com/vincentgong3mm/golang/study2019/exdiscoverygo"
	"testing"
)

func TestAddOne(t *testing.T) {
	n := []int{1, 2, 3, 4}
	exdiscoverygo.AddOne(n)
	fmt.Println(n)
}
