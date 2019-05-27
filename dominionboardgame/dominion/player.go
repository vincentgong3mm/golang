package dominion

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	deck            CardIDs // index 0 is the top card
	handCards       CardIDs
	HandCards       Cards
	cardPlayingArea CardIDs
	discardPile     CardIDs

	actions int
	buys    int
	coins   int

	chanGameMan chan MessageGameMan
	chanPlay    chan MessagePlay
}

func init() {
	fmt.Println("import dominon/player")

}

func (p *Player) InitChan() {
	p.chanGameMan = make(chan MessageGameMan)
	p.chanPlay = make(chan MessagePlay)

	go p.DoChanMessage()
}

func (p *Player) SendGameManMessage(msg *MessageGameMan) {
	p.chanGameMan <- *msg
}

func (p *Player) SendPlayMessage(msg *MessagePlay) {
	p.chanPlay <- *msg
}

func (p *Player) DoChanMessage() {
	fmt.Println("Run DoChanMessage", p.name, p.ID)
	for {
		select {
		case g := <-p.chanGameMan:
			fmt.Println("Receive From chanGameMan", g)
		case p := <-p.chanPlay:
			fmt.Println("Receive From Play", p)
		}
	}

	fmt.Println("exit DoChanMessage", p.name, p.ID)
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

	s += fmt.Sprintf("+Action(%d)\n", r.actions)
	s += fmt.Sprintf("+Buy(%d)\n", r.buys)
	s += fmt.Sprintf("+Coin(%d)\n", r.coins)
	s += fmt.Sprintf("+Deck")
	s += fmt.Sprintf("%s\n", r.deck)
	s += fmt.Sprintf("+Hand")
	s += fmt.Sprintf("%s\n", r.handCards)

	s += fmt.Sprintf("+CardPlayingArea")
	s += fmt.Sprintf("%s\n", r.cardPlayingArea)

	s += fmt.Sprintf("+DiscardPile")
	s += fmt.Sprintf("%s\n", r.discardPile)

	// test call ... GetLogInstance().Println(s)

	return s
}

func (r Player) StringHand() string {
	s := fmt.Sprintf("+Hand")
	s += fmt.Sprintf("%s", r.handCards)

	return s
}

func (r *Player) InitForNextTurn() {
	r.coins = 0
	r.buys = 1
	r.actions = 1

	r.deck.Shuffle()
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
	rand.Seed(time.Now().UTC().UnixNano())
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

func (r *Player) PutOnToTopDeck(id CardID) {
	tmpdeck := r.deck
	r.deck = r.deck[:0]

	r.deck = append(r.deck, id)
	r.deck = append(r.deck, tmpdeck...)
}

// Artisan : SpecialAbility
func (r *Player) PutCardFromHandToTopDeck(index int) error {
	if index >= len(r.handCards) {
		return errors.New(fmt.Sprintf("NOTE:Hand isn't %d index card.", index))
	}

	cardID := r.handCards[index]

	front := r.handCards[0:index]
	end := r.handCards[index+1 : len(r.handCards)]

	r.handCards = front
	r.handCards = append(r.handCards, end...)

	r.PutOnToTopDeck(cardID)

	return nil
}

// DrawCard is draw cards from deck to hand
func (r *Player) DrawCard(cnt int) (CardIDs, error) {
	if len(r.deck) < cnt {
		// add to deck
		r.AddDiscardPileToDeck()
	}

	if len(r.deck) < cnt {
		return CardIDs{}, errors.New(fmt.Sprintf("NOTE:not enough deck. deck is %d < %d", len(r.deck), cnt))
	}

	tmpCards := r.deck[0:cnt]
	r.deck = r.deck[cnt:len(r.deck)]
	r.handCards = append(r.handCards, tmpCards...)

	return tmpCards, nil
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

	// 한번 샀으면 구매 횟수 차감
	r.buys--
	r.GainCard(card, ToDiscardPile)

	return nil
}

func (r *Player) BuyCardGM(card CardID) {
	r.GainCard(card, ToDiscardPile)
}
func (r *Player) GainCardGM(card CardID) {
	r.GainCard(card, ToHand)
}

func (r *Player) DiscardFromHand(index int) error {
	if index >= len(r.handCards) {
		return errors.New("Invaild Hand Cards Index")
	}

	id := r.handCards[index]

	front := r.handCards[0:index]
	end := r.handCards[index+1 : len(r.handCards)]

	r.handCards = front
	r.handCards = append(r.handCards, end...)

	r.GainCard(id, ToDiscardPile)

	return nil
}

/*
// TranshCard is trash card to trash
func (r *Player) TrashCardFromHand(index int) error {
	if cardID, err := r.handCards.RemoveCard(int); err != nil {
		return err
	}

}
*/

func (r *Player) PlayCardFromHand(index int, gman *GameMan) error {
	if index >= len(r.handCards) {
		return errors.New("Invaild Hand Cards Index")
	}
	if r.actions < 1 {
		return errors.New("Player's actions is 0")
	}

	cardID := r.handCards[index]

	front := r.handCards[0:index]
	end := r.handCards[index+1 : len(r.handCards)]

	r.handCards = front
	r.handCards = append(r.handCards, end...)

	r.actions--
	card, exist := gman.cards[cardID]
	if exist == false {
		return errors.New(fmt.Sprintf("%s is not registed.", cardID))
	}

	card.DoAbility(r)
	card.DoSpecialAbility(r, gman)

	return nil
}

func (r *Player) RevealTopCardFromDeck(cnt int) (CardIDs, error) {
	if len(r.deck) < cnt {
		// add to deck
		r.AddDiscardPileToDeck()
	}

	if len(r.deck) < cnt {
		return CardIDs{}, errors.New(fmt.Sprintf("not enough deck. deck is %d < %d", len(r.deck), cnt))
	}

	cards := r.deck[0:cnt]

	return cards, nil
}

func (r *Player) PopTopCardFromDeck(cnt int) (CardIDs, error) {
	if len(r.deck) < cnt {
		// add to deck
		r.AddDiscardPileToDeck()
	}

	if len(r.deck) < cnt {
		return CardIDs{}, errors.New(fmt.Sprintf("not enough deck. deck is %d < %d", len(r.deck), cnt))
	}

	cards := r.deck[0:cnt]
	r.deck = r.deck[cnt:len(r.deck)]

	return cards, nil
}
