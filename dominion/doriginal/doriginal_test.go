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

	gb.CreateCard("Village",
		[]dom.CardType{dom.CardTypeAction},
		3)
	gb.CreateCard("Smithy",
		[]dom.CardType{dom.CardTypeAction},
		3)
	gb.CreateCard("Mill",
		[]dom.CardType{dom.CardTypeAction, dom.CardTypeVictory},
		3)
	gb.CreateCard("Gold",
		[]dom.CardType{dom.CardTypeTreasure},
		6)
	gb.CreateCard("Silver",
		[]dom.CardType{dom.CardTypeTreasure},
		3)
	gb.CreateCard("Copper",
		[]dom.CardType{dom.CardTypeTreasure},
		0)
	gb.CreateCard("Curse",
		[]dom.CardType{dom.CardTypeCurse},
		0)
	gb.CreateCard("Market",
		[]dom.CardType{dom.CardTypeAction},
		5)

	fmt.Println(gb)

	//fmt.Println(vCard.TermString())
	//fmt.Println(sCard.TermString())
}
