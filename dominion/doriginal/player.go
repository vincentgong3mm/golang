package doriginal

import "fmt"

// add counter -> Player.index
type Player struct {
	name      string
	index     int
	handCards []Card
	deckCards []Card

	actions int
	buys    int
}

func init() {
	fmt.Println("import d_original/player")
}

func CreateNewPlayer(name string) *Player {
	np := Player{name: name}

	return &np
}

func (r *Player) JoinGame() {

}

// DrawCard is draw cards from deck to hand
func (r *Player) DrawCard(cnt int) {
}

// MakeNewDesk is make new deck
func (r *Player) MakeNewDeck() {
}

// TranshCard is trash card to trash
func (r *Player) TrashCard() {
}
