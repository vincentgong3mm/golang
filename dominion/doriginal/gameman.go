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

func (r *GameBox) createCard(name string, cardType []CardType, cost int, ability []Ability) Card {
	CardID := r.genCardID()
	r.cards[CardID] = Card{name: name, CardID: CardID, cardType: cardType, cost: cost,
		Ability: ability}

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

func (r *GameBox) CreateAllCard() error {
	r.createCard("Festival", []CardType{CardTypeAction}, 5,
		[]Ability{{AbilityAddAction, 2}, {AbilityAddBuy, 1}, {AbilityAddCoin, 2}})
	r.createCard("Villiage", []CardType{CardTypeAction}, 3,
		[]Ability{{AbilityAddAction, 2}, {AbilityAddCard, 1}})
	r.createCard("Smithy", []CardType{CardTypeAction}, 4,
		[]Ability{{AbilityAddCard, 3}})
	r.createCard("Market", []CardType{CardTypeAction}, 5,
		[]Ability{{AbilityAddAction, 1}, {AbilityAddBuy, 1}, {AbilityAddCard, 1}, {AbilityAddCoin, 1}})
	r.createCard("Gold", []CardType{CardTypeTreasure}, 6,
		[]Ability{{AbilityAddCoin, 3}})
	r.createCard("Silver", []CardType{CardTypeTreasure}, 3,
		[]Ability{{AbilityAddCoin, 2}})
	r.createCard("Copper", []CardType{CardTypeTreasure}, 0,
		[]Ability{{AbilityAddCoin, 1}})

	return nil
}

func (r *GameBox) GMPlayAllCard() {
	for _, v := range r.cards {
		v.Play()
	}
}
