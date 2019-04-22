package doriginal

import "fmt"

type GameBox struct {
	cards     map[CardId]Card
	genCardId func() CardId
}

func init() {
	fmt.Println("import d_original/gamebox")
}

func CreateNewGameBox() *GameBox {
	n := GameBox{}
	n.cards = make(map[CardId]Card)
	n.genCardId = NewCardIdGenerator()

	return &n
}

func (r *GameBox) CreateCard(name string, cardType []CardType, cost int) Card {
	cardId := r.genCardId()
	r.cards[cardId] = Card{name: name, cardId: cardId, cardType: cardType, cost: cost}

	return r.cards[cardId]
}

func (r *GameBox) String() string {
	s := "Card\n"
	for _, v := range r.cards {
		s += v.String()
	}

	return s
}
