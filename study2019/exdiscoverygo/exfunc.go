package exdiscoverygo

import (
	"bufio"
	"fmt"
	"io"
)

func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

type VertexID int

func NewVertexIDGenerator_old() func() VertexID {
	var next VertexID
	return func() VertexID {
		next++
		return next
	}
}

type BinOp func(int, int) int
type BinSub func(int, int) int

func BinOpToBinSub(f BinOp) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}

// function abs
func NewVertexIDGenerator() func() VertexID {
	gen := NewIntGenerator()
	return func() VertexID {
		return VertexID(gen())
	}
}
