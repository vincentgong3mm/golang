package d_original

import "fmt"

type Card int

type Player struct {
	handCards []Card
	deckCards []Card

	actions int
	buys    int
}

func init() {
	fmt.Println("import d_original/player")
}

func CreateNewPlayer() *Player {
	np := Player{}

	return &np
}
