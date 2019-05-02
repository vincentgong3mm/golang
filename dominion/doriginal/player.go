package doriginal

import (
	"fmt"
	"math/rand"
)

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
	deck            CardIDs
	handCards       CardIDs
	cardPlayingArea CardIDs
	discardPile     CardIDs

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

func (r Player) String() string {
	s := ""
	s += fmt.Sprintf("Name:%s(ID:%d)\n", r.name, r.index)

	s += fmt.Sprintf("+Deck:")
	s += fmt.Sprintf("%s", r.deck)

	s += fmt.Sprintf("+Hand:")
	s += fmt.Sprintf("%s", r.handCards)

	s += fmt.Sprintf("+CardPlayingArea:")
	s += fmt.Sprintf("%s", r.cardPlayingArea)

	s += fmt.Sprintf("+DiscardPile:")
	s += fmt.Sprintf("%s", r.discardPile)

	/*
		for _, v := range r.handCards {
			s += fmt.Sprintf("%s(%d)-", v, v)
		}
		s = strings.TrimRight(s, "-")
		s += "]\n"
	*/

	return s
}

func (r *Player) AddDiscardPileToDeck() {
	// shuffle discard pile
	r.discardPile.Shuffle()

	// add iscard pile to deck
	r.addCardsToDeckBottom(&r.discardPile)
	r.discardPile = r.discardPile[0:0]

	//	r.deck.ShuffleCardIDs()
}

func (r *Player) addCardsToDeckBottom(cards *CardIDs) {
	r.deck = append(r.deck, *cards...)
}

func (r CardIDs) Shuffle() {
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
}

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
func (r *Player) DrawCard(cnt int) error {
	/*
		//
		if len(r.deck) < 0 {
			// add to deck
		}

		for i := 0; i < cnt; i++ {
			r.handCards = append(r.handCards, r.deck[0:cnt])
		}
	*/

	return nil
}

// TranshCard is trash card to trash
func (r *Player) TrashCard() {
}

func (r *Player) PlayCard(card *Card) error {
	//card.Action()
	return nil
}
