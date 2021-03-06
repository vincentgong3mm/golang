package dominion_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

	_, err := p1.DrawCard(5)
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
	/*
		gman := dom.CreateNewGameMan()
		gman.CreateAllCard()

		p1 := gman.CreateNewPlayer("jong")
		fmt.Println(p1)

		for i := 0; i < 3; i++ {
			p1.BuyCard(dom.Festival)
			fmt.Println(p1)
		}
	*/
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

	gman.BuyCard(dom.Festival, p1)
	gman.BuyCard(dom.Market, p1)
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

	//	fmt.Println(p1)
	p1.RevealTopCardFromDeck(3)
	//fmt.Println("revealCards:", revealCards)
	//fmt.Println(p1)

	gman.TrashTopCardFromDeck(p1, 1)
	//fmt.Println("PopCards:", popCards)
	//fmt.Println(p1)
	//fmt.Println(gman)

	p1.DrawCard(2)
	fmt.Println(gman)
	p1.PlayCardFromHand(0, gman)
	fmt.Println(gman)

}

func TestGameMan(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	// test get player
	p1 := gman.GetPlayer(1)

	fmt.Println(gman)

	fmt.Println(p1)
}

func TestInput(t *testing.T) {
	gman := CreateGameManAndSetSuppy()

	var in bytes.Buffer
	in.Write([]byte("abcd\n"))

	fmt.Println("Input Hand Card Index :")

	str, _ := gman.ReadInput()
	fmt.Println(str)

	//fmt.Println("Your Index :", in)
}

func TestInput2(t *testing.T) {
	gman := CreateGameManAndSetSuppy()

	gman.SetInputFromBuffer()
	gman.WriteInBuffer("100")

	str, _ := gman.ReadInput()
	fmt.Println(str)
}

func TestPlayAction(t *testing.T) {
	fmt.Println("TestPlayAction")
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	p1 := gman.GetPlayer(1)
	fmt.Println(p1)

	p1.GainCardGM(dom.Market)
	p1.GainCardGM(dom.Village)
	p1.GainCardGM(dom.Smithy)
	p1.GainCardGM(dom.Festival)
	//p1.GainCardGM(dom.Copper)
	//p1.GainCardGM(dom.Silver)
	fmt.Println(p1)

	p1.PlayCardFromHand(0, gman)
	fmt.Println(p1)
	p1.PlayCardFromHand(0, gman)
	fmt.Println(p1)
	p1.PlayCardFromHand(0, gman)
	fmt.Println(p1)
	err := p1.PlayCardFromHand(0, gman)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)

	/*
		p1.PlayCardFromHand(0, gman)
		fmt.Println(p1)
		p1.PlayCardFromHand(0, gman)
		fmt.Println(p1)
		p1.PlayCardFromHand(0, gman)
		fmt.Println(p1)
	*/

}

func TestArtisan(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)

	p1.GainCardGM(dom.Market)
	p1.GainCardGM(dom.Artisan)
	fmt.Println(p1)

	gman.SetInputFromBuffer()
	p1.PlayCardFromHand(0, gman)
	fmt.Println(p1)

	// for Artisan add buffer 1
	gman.WriteInBuffer("3\n") // gain card 7(Festival), 7 is supply's index
	gman.WriteInBuffer("3\n") // put Market onto player's deck
	p1.PlayCardFromHand(0, gman)
	fmt.Println(p1)
}

func TestChapel(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)

	p1.GainCardGM(dom.Copper)
	p1.GainCardGM(dom.Chapel)
	p1.GainCardGM(dom.Estate)
	p1.GainCardGM(dom.Market)
	p1.GainCardGM(dom.Artisan)
	fmt.Println(p1)

	gman.SetInputFromBuffer()
	gman.WriteInBuffer("2\n") // gain card 7(Festival), 7 is supply's index
	gman.WriteInBuffer("1\n") // put Market onto player's deck
	if err := p1.PlayCardFromHand(1, gman); err != nil {
		fmt.Println(err)
	}
	/*
		gman.WriteInBuffer("3\n") // put Market onto player's deck
		p1.PlayCardFromHand(0, gman)
		fmt.Println(p1)
	*/
}

func TestCellar(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)

	p1.GainCardGM(dom.Copper)
	p1.GainCardGM(dom.Chapel)
	p1.GainCardGM(dom.Estate)
	p1.GainCardGM(dom.Cellar)
	p1.GainCardGM(dom.Laboratory)
	fmt.Println(p1)
	gman.SetInputFromBuffer()
	gman.WriteInBuffer("2\n")
	gman.WriteInBuffer("\n")

	if err := p1.PlayCardFromHand(3, gman); err != nil {
		fmt.Println(err)
	}
	if err := p1.PlayCardFromHand(2, gman); err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1)
}

func TestWorkshop(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)

	p1.GainCardGM(dom.Copper)
	p1.GainCardGM(dom.Chapel)
	p1.GainCardGM(dom.Estate)
	p1.GainCardGM(dom.Cellar)
	p1.GainCardGM(dom.Workshop)

	fmt.Println(p1)
	gman.SetInputFromBuffer()
	gman.WriteInBuffer("12\n") // select supply index

	if err := p1.PlayCardFromHand(4, gman); err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1)

}

func TestChan(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)

	p1.SendGameManMessage(&dom.MessageGameMan{Msg: dom.MsgFinishGame})
	p1.SendGameManMessage(&dom.MessageGameMan{Msg: dom.MsgCloseGame})
	p1.SendGameManMessage(&dom.MessageGameMan{Msg: dom.MsgFinishGame})
	p1.SendGameManMessage(&dom.MessageGameMan{Msg: dom.MsgFinishGame})
	//p1.SendPlayMessage(MessagePlay{})

	p1.SendPlayMessage(&dom.MessagePlay{Msg: dom.MsgThisPlayCard, Step: 7})
	p1.SendPlayMessage(&dom.MessagePlay{Msg: dom.MsgOtherDoEffect, Step: 2})
	p1.SendPlayMessage(&dom.MessagePlay{Msg: dom.MsgThisPlayCard, Step: 3})

	fmt.Println(p1)

	time.Sleep(500 * time.Millisecond)
}

func TestWitch(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)
	p2 := gman.GetPlayer(2)

	p1.DrawCard(4)
	p2.DrawCard(5)

	p1.GainCardGM(dom.Witch)

	if err := p1.PlayCardFromHand(4, gman); err != nil {
		fmt.Println(err)
	}

	time.Sleep(100 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	fmt.Println(p1)
	fmt.Println(p2)

	time.Sleep(100 * time.Millisecond)

	time.Sleep(1000 * time.Millisecond)

}

func TestCouncilRoom(t *testing.T) {
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)
	p2 := gman.GetPlayer(2)

	p1.DrawCard(4)
	p2.DrawCard(5)

	p1.GainCardGM(dom.CouncilRoom)

	time.Sleep(1000 * time.Millisecond)

	if err := p1.PlayCardFromHand(4, gman); err != nil {
		fmt.Println(err)
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println(p1)
	fmt.Println(p2)

	time.Sleep(1000 * time.Millisecond)

}

func TestMine(t *testing.T) {
	fmt.Println("TestMine")

	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)

	//p1.DrawCard(4)
	p1.GainCardGM(dom.Estate)
	p1.GainCardGM(dom.Copper)
	p1.GainCardGM(dom.Estate)
	p1.GainCardGM(dom.Silver)
	p1.GainCardGM(dom.Mine)

	fmt.Println(p1)

	gman.SetInputFromBuffer()
	gman.WriteInBuffer("1\n") // select supply index
	gman.SetInputFromBuffer()
	gman.WriteInBuffer("1\n") // select supply index

	if err := p1.PlayCardFromHand(4, gman); err != nil {
		fmt.Println(err)
	}

	fmt.Println(p1)

	for i := 0; i < 20; i++ {
		fmt.Println("-------------------============================================")
	}
	gman.Wait()

}

func TestOnePlay(t *testing.T) {
	fmt.Println("TestPlay")
	g := dom.CreateNewGameMan()
	g.CreateAllCard()
	g.RegistCardToSuppy(dom.SetFirstGame, 1)

	p := g.CreateNewPlayer("gong")

	p.DrawCard(5)

	g.SetInputFromBuffer()
	g.WriteInBuffer("1\n") // select supply index
	cardIndexInHand, s := g.ReadInput(">>>> choose index in your hand. #")

	fmt.Println(cardIndexInHand, s)

	if err := p.PlayCardFromHand(cardIndexInHand, g); err != nil {
		fmt.Println(err)
	}

	// player 상태 메시지 보여주고 입력 받아서 진행
	// 1. 액션상태 표시 : 액션 진행할 카드 선택
	//		- 액션 없는 카드 사용 못하세 제한
	// 2. buy 상태 표시, 현재 coin개수 + treature  카드 표시
	// 3. clean up

	g.Wait()
}

func TestAtoi(t *testing.T) {
	n, err := strconv.Atoi("b")

	fmt.Println(n, err)
}
