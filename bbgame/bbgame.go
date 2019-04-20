package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ResultType is
type ResultType int

type ResultCount struct {
	rt    ResultType
	count int
}

const (
	Ball = 0 + iota
	Strike
	Out
)

const (
	MaxNumberCount = 3
)

var Results = [...]string{
	"Ball",
	"Strike",
	"Out",
}

func (r ResultType) String() string {
	return Results[r%MaxNumberCount]
}

// AINumber is
type AINumber struct {
	number []int
}

func newAINumber() *AINumber {
	ain := AINumber{}

	for i := range ain.number {
		ain.number[i] = -1
	}

	return &ain
}

func (r *AINumber) createNewNunmber() {

	rand.Seed(time.Now().UnixNano())

	tmp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})

	r.number = tmp[:MaxNumberCount]
}

func (r *AINumber) checkNumer(n []int) []ResultType {
	checkOneNumber := func(position int, value int) ResultType {
		for i, v := range r.number {
			if v == value {
				if i == position {
					return Strike
				} else {
					return Ball
				}
			}
		}
		return Out
	}

	var rt []ResultType
	for i, v := range n {
		tmp := checkOneNumber(i, v)
		rt = append(rt, tmp)
	}

	return rt

}

func inputNumber() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter three number ")
	in, _ := reader.ReadString('\n')

	//fmt.Println(in)

	var rs []string
	var rn []int
	in = in[:len(in)-1]

	rs = strings.Split(in, " ")

	for _, s := range rs {
		tn, _ := strconv.Atoi(s)
		rn = append(rn, tn)
	}

	return rn
}

func main() {

	fmt.Println("Start Baseball Game.")

	ain := newAINumber()

	ain.createNewNunmber()

	// Print AI Number
	//fmt.Println(ain)

	for {
		in := inputNumber()

		//fmt.Println(in)

		result := ain.checkNumer(in)

		//rc := []ResultCount{Ball, 0}
		rc := make(map[ResultType]int)

		for _, v := range result {
			rc[v]++
		}

		isEnd := true

		ss := ""
		for _, v := range in {
			ss += fmt.Sprintf("%d ", v)
		}

		rs := ""
		for i, v := range rc {
			//fmt.Println(v, i)
			rs += fmt.Sprintf("|%d %s", v, i)

			if i != Strike {
				isEnd = false
			}
		}
		rs += "|"

		fmt.Println(ss, "=> ", rs)

		if isEnd == true {
			fmt.Println("Well done.")
			break
		}
	}
}
