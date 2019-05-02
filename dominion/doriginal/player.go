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
	deck            []CardID
	handCards       []CardID
	cardPlayingArea []CardID
	discardPile     []CardID

	actions int
	buys    int
	coins   int
}

func init() {
	fmt.Println("import d_original/player")
}

type PlayerID int

func NewPlayerIDGenerator() func() PlayerID {
	var next int
	return func() PlayerID {
		next++
		return PlayerID(next)
	}
}

/*
func (r *Player) String() string {

}
*/

func (r *Player) JoinGame() {

}

// Gain a card to
type GainedCard int

const (
	ToDiscardPile GainedCard = 0 + iota
	ToDeck
	ToHand
)

// GainCard is gain a card from Supply
func (r *Player) GainCard(id CardID, to GainedCard) {
	switch to {
	case ToDiscardPile:
		r.discardPile = append(r.discardPile, id)
	case ToDeck:
		r.deck = append(r.deck, id)
	case ToHand:
		r.handCards = append(r.handCards, id)
	}
}

// DrawCard is draw cards from deck to hand
func (r *Player) DrawCard(cnt int) {
}

// Shuffle is make new deck
func (r *Player) ShuffleDeck() {
}

// TranshCard is trash card to trash
func (r *Player) TrashCard() {
}

func (r *Player) PlayCard(card *Card) error {
	//card.Action()
	return nil
}
