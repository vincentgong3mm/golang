package doriginal

import "fmt"

// 이렇게 까지는 필요하지 않음.
// player의 turn에 해야할 순서
// type Actioner interface {
// 	action() error
// 	buy() error
// 	clean() error
// 	draw() error
// }

// add counter -> Player.index
type Player struct {
	name            string
	index           int
	deck            []Card
	handCards       []Card
	cardPlayingArea []Card
	discardPile     []Card

	actions int
	buys    int
	coins   int
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

func (r *Player) PlayCard(card *Card) error {
	//card.Action()
	return nil
}
