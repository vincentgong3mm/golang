package dominion_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	dom "github.com/vincentgong3mm/golang/dominionboardgame/dominion"
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

	gb := dom.CreateNewGameMan()
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
	gman := dom.CreateNewGameMan()

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
	gman := dom.CreateNewGameMan()
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
	gman := dom.CreateNewGameMan()
	gman.CreateAllCard()

	p1 := gman.CreateNewPlayer("jong")
	fmt.Println(p1)

	for i := 0; i < 3; i++ {
		p1.BuyCard(dom.Festival)
		fmt.Println(p1)
	}
}

func CreateGameMan() *dom.GameMan {
	gman := dom.CreateNewGameMan()
	gman.CreateAllCard()

	return gman
}

func CreateGameManAndSetSuppy() *dom.GameMan {
	gman := dom.CreateNewGameMan()
	gman.CreateAllCard()

	gman.RegistCardToSuppy(dom.SetFirstGame, 2)

	return gman
}

func TestSupply(t *testing.T) {
	gman := CreateGameMan()

	gman.RegistCardToSuppy(dom.SetFirstGame, 2)
	fmt.Println(gman.StringSupply())

	// create player1
	p1 := gman.CreateNewPlayer("jong")
	fmt.Println(p1)

	// create player2
	p2 := gman.CreateNewPlayer("seong")
	fmt.Println(p2)

	// print supply
	fmt.Println(gman.StringSupply())
}

func CreateTwoPlayer(g *dom.GameMan) {
	g.CreateNewPlayer("jong")
	g.CreateNewPlayer("seong")
}

func TestDrawAndCleanUp(t *testing.T) {
	logger := dom.GetLogInstance()

	// test create game
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	//fmt.Println(gman)

	// test get player
	p1 := gman.GetPlayer(1)
	p2 := gman.GetPlayer(2)
	fmt.Println(p1, p2)

	// test draw
	p1.DrawCard(5)
	p2.DrawCard(5)
	fmt.Println(p1, p2)

	p1.BuyCard(dom.Festival)
	p2.BuyCard(dom.Market)
	fmt.Println(p1, p2)

	logger.Println("TestDrawAndCleanUp")
}

func TestShowAllCard(t *testing.T) {

	logger := dom.GetLogInstance()

	// test create game
	gman := CreateGameManAndSetSuppy()
	fmt.Println(gman)

	logger.Println("TestShowAllCard")
}

func TestDeck(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)

	// test get player
	p1 := gman.GetPlayer(1)

	fmt.Println(p1)
	revealCards, _ := p1.RevealTopCardFromDeck(3)
	fmt.Println("revealCards:", revealCards)
	fmt.Println(p1)

	popCards, _ := gman.TrashTopCardFromDeck(p1, 3)
	fmt.Println("PopCards:", popCards)
	fmt.Println(p1)
	fmt.Println(gman)
}