package doriginal

import "fmt"

type GameBox struct {
	cards     map[CardID]Card
	genCardID func() CardID
}

func init() {
	fmt.Println("import d_original/gamebox")
}

func CreateNewGameBox() *GameBox {
	n := GameBox{}
	n.cards = make(map[CardID]Card)
	n.genCardID = NewCardIDGenerator()

	return &n
}

func (r *GameBox) CreateCard(name string, cardType []CardType, cost int) Card {
	CardID := r.genCardID()
	r.cards[CardID] = Card{name: name, CardID: CardID, cardType: cardType, cost: cost}

	return r.cards[CardID]
}

func (r *GameBox) String() string {
	s := "GameBox Info\n"
	s += "Card List\n"
	for _, v := range r.cards {
		s += v.String()
	}

	return s
}
