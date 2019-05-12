package main

import (
	"bufio"
	"fmt"
	"os"

	dom "github.com/vincentgong3mm/golang/dominionboardgame/dominion"
)

func waitExit() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')

	fmt.Println(in)

}

func main() {
	fmt.Println("Let's start Dominion!")

	gman := dom.CreateNewGameMan()
	fmt.Println(gman)
	gman.SetInputFromBuffer()
	gman.WriteInBuffer("asdfasfd99999")

	fmt.Println("Input String :")
	str, _ := gman.ReadInput()
	fmt.Println("str=", str)

	waitExit()
}
