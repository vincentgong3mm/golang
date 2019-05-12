package main

import (
	"bufio"
	"fmt"
	"os"

	dom "github.com/vincentgong3mm/golang/dominionboardgame/dominion"
)

func waitExit() {
	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')
}

func main() {
	fmt.Println("Let's start Dominion!")

	gman := dom.CreateNewGameMan()
	fmt.Println(gman)
	waitExit()
}
