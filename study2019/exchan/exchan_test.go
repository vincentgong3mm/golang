package exchan

import (
	"fmt"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for num := range c {
		fmt.Println(num)
	}
}

func TestExchan(t *testing.T) {
	fmt.Println("Exchan")
}

func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func TestFibonacci(t *testing.T) {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ",")
	}
}

func BabyNames(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)

		for _, f := range first {
			for _, s := range second {
				//fmt.Print("in func", f, s, "\n")
				c <- string(f) + string(s)
			}
		}
	}()

	return c
}

func TestBabyNames(t *testing.T) {
	for n := range BabyNames("abcdef", "01910") {
		fmt.Print(n, ", ")
	}
}

func TestCloseChannel(t *testing.T) {
	c := make(chan int)
	close(c)

	a, b := <-c

	fmt.Println(a, b)

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func TestPlusOne(t *testing.T) {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
}

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func TestChain(t *testing.T) {

}

func TestFanOut(t *testing.T) {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}
