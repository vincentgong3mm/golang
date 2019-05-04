package doriginal

import (
	"errors"
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
	ID              PlayerID
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
	s += fmt.Sprintf("@Player:%s(ID:%d)\n", r.name, r.ID)

	s += fmt.Sprintf("+Action#%d\n", r.actions)
	s += fmt.Sprintf("+Buy#%d\n", r.buys)
	s += fmt.Sprintf("+Coin#%d\n", r.coins)
	s += fmt.Sprintf("+Deck")
	s += fmt.Sprintf("%s", r.deck)

	s += fmt.Sprintf("+Hand")
	s += fmt.Sprintf("%s", r.handCards)

	s += fmt.Sprintf("+CardPlayingArea")
	s += fmt.Sprintf("%s", r.cardPlayingArea)

	s += fmt.Sprintf("+DiscardPile")
	s += fmt.Sprintf("%s", r.discardPile)

	return s
}

func (r *Player) AddDiscardPileToDeck() {
	// shuffle discard pile
	r.discardPile.Shuffle()

	// add iscard pile to deck
	r.addCardsToDeckBottom(&r.discardPile)

	// empty discard pile
	r.discardPile = r.discardPile[0:0]
}

func (r *Player) addCardsToDeckBottom(cards *CardIDs) {
	r.deck = append(r.deck, *cards...)
}

func (r CardIDs) Shuffle() {
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
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
	if len(r.deck) < cnt {
		// add to deck
		r.AddDiscardPileToDeck()
	}

	if len(r.deck) < cnt {
		return errors.New(fmt.Sprintf("not enough deck. deck is %d < %d", len(r.deck), cnt))
	}

	tmpCards := r.deck[0:cnt]
	r.deck = r.deck[cnt:len(r.deck)]
	r.handCards = append(r.handCards, tmpCards...)

	r.buys = 1
	r.actions = 1
	r.coins = 0

	return nil
}

func (r *Player) PlayActionCard(index int) {
}

func (r *Player) PlayTreasureCard(index int) {
}

func (r *Player) CleanUp() {
	// empty hand cards to discardpile
	r.discardPile = append(r.discardPile, r.handCards...)

	// emptly hand cards
	r.handCards = r.handCards[:0]
}

func (r *Player) BuyCard(card CardID) error {
	if r.buys <= 0 {
		return errors.New(fmt.Sprintf("can't buy. buy count is %d", r.buys))
	}

	// check Supply
	// if len(xxxx)

	r.buys--
	r.GainCard(card, ToDiscardPile)

	return nil
}

// TranshCard is trash card to trash
func (r *Player) TrashCard() {
}

func (r *Player) PlayCard(card *Card) error {
	//card.Action()
	return nil
}
