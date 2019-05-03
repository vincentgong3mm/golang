package doriginal

import "fmt"

type GameBox struct {
	cards map[CardID]Card

	supply *Supply

	players     map[PlayerID]*Player
	genPlayerID func() PlayerID
}

func init() {
	fmt.Println("import d_original/gamebox")
}

func CreateNewGameBox() *GameBox {
	n := GameBox{}
	n.cards = make(map[CardID]Card)
	n.supply = CreateNewSupply()
	n.players = make(map[PlayerID]*Player)
	n.genPlayerID = NewPlayerIDGenerator()

	return &n
}

func (r *GameBox) createCard(cardID CardID, cardType []CardType, cost int, ability []Ability) Card {
	//if r.cards[cardID]
	r.cards[cardID] = Card{name: "", CardID: cardID, cardType: cardType, cost: cost,
		Ability: ability}

	return r.cards[cardID]
}

func (r *GameBox) RegistCardToSuppy(t SupplySet) {
	switch t {
	case SetFirstGame:
		r.supply.RegistCard(Copper, 50)
		r.supply.RegistCard(Silver, 40)
		r.supply.RegistCard(Gold, 30)
		r.supply.RegistCard(Market, 10)
		r.supply.RegistCard(Festival, 10)
		r.supply.RegistCard(Smithy, 10)
	case SetBigMoney:
	}

}

// gong/regist supply...

func (r *GameBox) CreateNewPlayer(name string) *Player {
	playerID := r.genPlayerID()
	player := Player{name: name}

	// inser Pplayer Point to map
	r.players[playerID] = &player

	t, _ := r.players[playerID]

	r.gainBeginHandCard(t)

	return r.players[playerID]
}

func (r *GameBox) gainBeginHandCard(player *Player) {
	for i := 0; i < 7; i++ {
		player.GainCard(Copper, ToDeck)
	}

	for i := 0; i < 3; i++ {
		player.GainCard(Estate, ToDeck)
	}

	// first shuffle deck
	player.deck.Shuffle()

	// test ..
	player.GainCard(Village, ToDiscardPile)
	player.GainCard(Market, ToDiscardPile)
	player.GainCard(Smithy, ToDiscardPile)
	player.GainCard(Smithy, ToDiscardPile)
}

func (r *GameBox) GetCard(cardID CardID) *Card {
	c, exist := r.cards[cardID]

	if exist == true {
		return &c
	}
	return nil
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
	r.createCard(Festival, []CardType{CardTypeAction}, 5,
		[]Ability{{AbilityAddAction, 2}, {AbilityAddBuy, 1}, {AbilityAddCoin, 2}})
	r.createCard(Village, []CardType{CardTypeAction}, 3,
		[]Ability{{AbilityAddAction, 2}, {AbilityAddCard, 1}})
	r.createCard(Smithy, []CardType{CardTypeAction}, 4,
		[]Ability{{AbilityAddCard, 3}})
	r.createCard(Market, []CardType{CardTypeAction}, 5,
		[]Ability{{AbilityAddAction, 1}, {AbilityAddBuy, 1}, {AbilityAddCard, 1}, {AbilityAddCoin, 1}})
	r.createCard(Gold, []CardType{CardTypeTreasure}, 6,
		[]Ability{{AbilityAddCoin, 3}})
	r.createCard(Silver, []CardType{CardTypeTreasure}, 3,
		[]Ability{{AbilityAddCoin, 2}})
	r.createCard(Copper, []CardType{CardTypeTreasure}, 0,
		[]Ability{{AbilityAddCoin, 1}})
	r.createCard(Province, []CardType{CardTypeVictory}, 8,
		[]Ability{{AbilityAddVictory, 6}})
	r.createCard(Duchy, []CardType{CardTypeVictory}, 5,
		[]Ability{{AbilityAddVictory, 3}})
	r.createCard(Estate, []CardType{CardTypeVictory}, 2,
		[]Ability{{AbilityAddVictory, 1}})

	return nil
}

func (r *GameBox) GMPlayAllCard(player *Player) {
	for _, v := range r.cards {
		v.Play(player)
	}
}
