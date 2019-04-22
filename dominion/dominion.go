package main

import (
	"bufio"
	"fmt"
	"os"

	dom "github.com/vincentgong3mm/golang/dominion/doriginal"
)

func waitExit() {
	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')
}

func main() {
	fmt.Println("Let's start Dominion!")

	np := dom.CreateNewPlayer("jong")
	fmt.Println(np)

	np2 := dom.CreateNewPlayer("seong")
	fmt.Println(np2)

	waitExit()
}
