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

func CreateGameManAndSetSuppy() *dom.GameMan {
	gman := dom.CreateNewGameMan()
	gman.CreateAllCard()

	gman.RegistCardToSuppy(dom.SetFirstGame, 2)

	return gman
}

func CreateTwoPlayer(g *dom.GameMan) {
	g.CreateNewPlayer("jong")
	g.CreateNewPlayer("seong")
}

func main() {
	/*
		fmt.Println("Let's start Dominion!")

		gman := dom.CreateNewGameMan()
		fmt.Println(gman)
		gman.SetInputFromBuffer()
		gman.WriteInBuffer("10\n11\n")

		fmt.Println("Input String :")
		str, _ := gman.ReadInput()
		fmt.Println("str=", str)
		str, _ = gman.ReadInput()
		fmt.Println("str=", str)
	*/
	gman := CreateGameManAndSetSuppy()
	CreateTwoPlayer(gman)
	fmt.Println(gman)
	p1 := gman.GetPlayer(1)

	/*

		// test Artisan
		p1.GainCardGM(dom.Market)
		p1.GainCardGM(dom.Artisan)
		fmt.Println(p1)

		//gman.SetInputFromBuffer()
		p1.PlayCardFromHand(0, gman)
		fmt.Println(p1)

		// for Artisan add buffer 1
		//gman.WriteInBuffer("7") // gain card 7(Festival), 7 is supply's index
		//gman.WriteInBuffer("1")  // put Market onto player's deck
		p1.PlayCardFromHand(0, gman)
		fmt.Println(p1)
	*/

	// tet chapel
	p1.GainCardGM(dom.Copper)
	p1.GainCardGM(dom.Chapel)
	p1.GainCardGM(dom.Estate)
	p1.GainCardGM(dom.Market)
	p1.GainCardGM(dom.Artisan)
	fmt.Println(p1)

	if err := p1.PlayCardFromHand(1, gman); err != nil {
		fmt.Println(err)
	}

	waitExit()
}
