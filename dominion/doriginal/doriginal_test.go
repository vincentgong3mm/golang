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
	/*
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
	*/

	fmt.Println(gb)
}

func TestPlayCard(t *testing.T) {
	gman := dom.CreateNewGameBox()

	gman.CreateAllCard()

	fmt.Println(gman)

	p1 := gman.CreateNewPlayer("jong")
	fmt.Println("begin player", p1)

	p1.AddDiscardPileToDeck()
	fmt.Println("AddDiscardPileToDeck", p1)
	fmt.Println(p1)

	p1.DrawCard(5)
	fmt.Println("DrawCard", p1)
	p1.DrawCard(5)
	fmt.Println("DrawCard", p1)
	p1.DrawCard(5)
	fmt.Println("DrawCard", p1)
	err := p1.DrawCard(5)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCleanUp(t *testing.T) {
	gman := dom.CreateNewGameBox()
	gman.CreateAllCard()

	//	fmt.Println(gman)

	p1 := gman.CreateNewPlayer("jong")
	fmt.Println(p1)

	for i := 0; i < 5; i++ {
		fmt.Println("Draw and Cleanup", i)
		p1.DrawCard(5)
		fmt.Println(p1)
		p1.CleanUp()
		fmt.Println(p1)
	}

}

func TestBuyCard(t *testing.T) {
	gman := dom.CreateNewGameBox()
	gman.CreateAllCard()

	p1 := gman.CreateNewPlayer("jong")
	fmt.Println(p1)

	for i := 0; i < 3; i++ {
		p1.BuyFromSupply(dom.Festival, dom.ToDiscardPile)
		fmt.Println(p1)
	}
}

func CreateGameBox() *dom.GameBox {
	gman := dom.CreateNewGameBox()
	gman.CreateAllCard()

	return gman
}

func TestSupply(t *testing.T) {
	gman := CreateGameBox()

	gman.RegistCardToSuppy(dom.SetFirstGame, 2)
	fmt.Println(gman.StringSupply())

	p1 := gman.CreateNewPlayer("jong")
	fmt.Println(p1)

	p2 := gman.CreateNewPlayer("seong")
	fmt.Println(p2)

	fmt.Println(gman.StringSupply())
}
