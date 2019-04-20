package main

import (
	"bufio"
	"fmt"
	dom "github.com/vincentgong3mm/golang/dominion/d_original"
	"os"
)

func waitExit() {
	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')
}

func main() {
	fmt.Println("Let's start Dominion!")

	np := dom.CreateNewPlayer()

	fmt.Println(np)

	waitExit()
}
