package exdiscoverygo_test

import (
	"fmt"
	exd "github.com/vincentgong3mm/golang/study2019/exdiscoverygo"
	"strings"
	"testing"
)

// simple function
func TestAddOne(t *testing.T) {
	n := []int{1, 2, 3, 4}
	exd.AddOne(n)
	fmt.Println(n)
}

// higher-order function
func TestReadFromPrint(t *testing.T) {
	r := strings.NewReader("1111\nbbbb\n")
	err := exd.ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})

	if err != nil {
		fmt.Println(err)
	}

}

// closure
func TestReadFrom_append(t *testing.T) {
	r := strings.NewReader("111\n2222\nccc")

	var lines []string
	err := exd.ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(lines)
}

// generator #1
func TestNewIntGenerator(t *testing.T) {
	gen := exd.NewIntGenerator()
	fmt.Println(gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen())
}

// generator #2
func TestNewIntGenerator_multiple(t *testing.T) {
	gen1 := exd.NewIntGenerator()
	gen2 := exd.NewIntGenerator()

	fmt.Println(gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1())

}

// generator #3
func TestNetVextexIDGenerator(t *testing.T) {
	gen := exd.NewVertexIDGenerator()
	fmt.Println(gen())
}

// type func name
func TestBinOpToBinSub(t *testing.T) {
	sub := exd.BinOpToBinSub(func(a, b int) int {
		return a + b
	})

	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
}
