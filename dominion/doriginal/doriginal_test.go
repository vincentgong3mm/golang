package doriginal_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	dom "github.com/vincentgong3mm/golang/dominion/doriginal"
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

func TestCardType(t *testing.T) {
	return
	//ct := dom.CardTypeAction
	var ct dom.CardType

	ct = dom.CardTypeAction
	fmt.Println(ct)
}

func TestCreateCard(t *testing.T) {

	gb := dom.CreateNewGameBox()

	vCard := gb.CreateCard("Village", []dom.CardType{dom.CardTypeAction}, 3)
	sCard := gb.CreateCard("Smithy", []dom.CardType{dom.CardTypeAction, dom.CardTypeVictory}, 3)

	fmt.Println(gb)

	fmt.Println(vCard.TermString())
	fmt.Println(sCard.TermString())
}
