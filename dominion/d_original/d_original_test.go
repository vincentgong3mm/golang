package d_original_test

import (
	"bufio"
	"fmt"
	dom "github.com/vincentgong3mm/golang/dominion/d_original"
	"os"
	"testing"
)

func waitExit() {
	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')
}

func TestMain(t *testing.T) {
	return
	fmt.Println("Let's start Dominion!")

	np := dom.CreateNewPlayer("jong")
	fmt.Println(np)

	np2 := dom.CreateNewPlayer("seong")
	fmt.Println(np2)

	waitExit()
}

func TestGameBox(t *testing.T) {
	gb := dom.CreateNewGameBox()

	nc := dom.CreateNewCard("Village")
	gb.CreateCard(*nc)

	fmt.Println(gb)

}
