package doriginal

import "fmt"

type GameMan struct {
	cards map[CardID]Card

	cards2 map[CardID]Actioner

	supply *Supply

	players     map[PlayerID]*Player
	genPlayerID func() PlayerID
}

func init() {
	// Create and Get Log Instance
	GetLogInstance()

	logger.Println("import d_original/GameMan")
	fmt.Println("import d_original/GameMan")
}

func CreateNewGameMan() *GameMan {
	n := GameMan{}
	n.cards = make(map[CardID]Card)

	n.cards2 = make(map[CardID]Actioner)

	n.supply = CreateNewSupply()
	n.players = make(map[PlayerID]*Player)
	n.genPlayerID = NewPlayerIDGenerator()

	return &n
}

func (r *GameMan) createCard(cardID CardID, cardType []CardType, cost int, ability []Ability) Card {
	r.cards[cardID] = Card{name: "", CardID: cardID, cardType: cardType, cost: cost,
		Ability: ability}

	return r.cards[cardID]
}

func (r *GameMan) createCard2(cardID CardID, cardType []CardType, cost int, ability []Ability) Actioner {
	switch cardID {
	case Thief:
		r.cards2[cardID] = &CardThief{comment: 10, Card: Card{name: "", CardID: cardID, cardType: cardType, cost: cost, Ability: ability}}
	default:
		r.cards2[cardID] = &Card{name: "", CardID: cardID, cardType: cardType, cost: cost, Ability: ability}
	}

	return r.cards2[cardID]
}

func (r *GameMan) RegistCardToSuppy(t SupplySet, players int) {
	estate := 8
	duchy := 8
	province := 8

	switch players {
	case 2:
		estate = 8 + players*3 // 3 coper per player
		duchy = 8
		province = 8
	case 3:
		estate = 10 + players*3 // 3 coper per player
		duchy = 10
		province = 10
	default:
		estate = 12 + players*3 // 3 coper per player
		duchy = 12
		province = 12
	}

	switch t {
	case SetFirstGame:
		r.supply.RegistCard(Copper, 50)
		r.supply.RegistCard(Silver, 40)
		r.supply.RegistCard(Gold, 30)
		r.supply.RegistCard(Estate, estate)
		r.supply.RegistCard(Duchy, duchy)
		r.supply.RegistCard(Province, province)
		r.supply.RegistCard(Market, 10)
		r.supply.RegistCard(Festival, 10)
		r.supply.RegistCard(Smithy, 10)
	case SetBigMoney:
		r.supply.RegistCard(Copper, 50)
	}

}

func (r *GameMan) CreateNewPlayer(name string) *Player {
	playerID := r.genPlayerID()
	player := Player{name: name, ID: playerID}

	// inser Pplayer Point to map
	r.players[playerID] = &player

	t, _ := r.players[playerID]

	r.gainBeginHandCard(t)

	return r.players[playerID]
}

func (r *GameMan) GetPlayer(id PlayerID) *Player {
	p, exist := r.players[id]

	if exist == true {
		return p
	}

	return nil
}

func (r *GameMan) gainPlayerFromSupply(id CardID, player *Player) bool {
	if r.supply.Pop(id) == true {
		player.GainCard(id, ToDeck)
		return true
	}

	return false
}

func (r *GameMan) gainBeginHandCard(player *Player) {
	// draw 7 copper`
	for i := 0; i < 7; i++ {
		r.gainPlayerFromSupply(Copper, player)
	}

	// draw 3 estate
	for i := 0; i < 3; i++ {
		r.gainPlayerFromSupply(Estate, player)
	}

	// first shuffle deck
	player.deck.Shuffle()

	// test ..
	/*
		player.GainCard(Village, ToDiscardPile)
		player.GainCard(Market, ToDiscardPile)
		player.GainCard(Smithy, ToDiscardPile)
		player.GainCard(Smithy, ToDiscardPile)
	*/
}

/*
func (r *GameMan) GetCard(cardID CardID) *Card {
	c, exist := r.cards[cardID]

	if exist == true {
		return &c
	}
	return nil
}
*/

func (r *GameMan) String() string {
	s := "GameMan Info\n"
	s += "Card List\n"
	for _, v := range r.cards {
		s += v.String()
	}

	s += "Supply List\n"
	s += r.supply.String()

	s += "Player List\n"
	for _, v := range r.players {
		s += v.String()
	}

	return s
}

func (r GameMan) StringSupply() string {
	return r.supply.String()
}

func (r *GameMan) CreateAllCard() error {
	r.createCard2(Thief, []CardType{CardTypeAction}, 5,
		[]Ability{{AbilityAddAction, 2}, {AbilityAddBuy, 1}, {AbilityAddCoin, 2}})

	r.createCard2(Smithy, []CardType{CardTypeAction}, 4,
		[]Ability{{AbilityAddAction, 2}, {AbilityAddBuy, 1}, {AbilityAddCoin, 2}})

	fmt.Println("create card2+++")
	fmt.Println(r.cards2)

	for _, v := range r.cards2 {
		v.Draw(nil)
		fmt.Println(v)
	}
	fmt.Println("create card2---")
	fmt.Println(r.cards2)
	fmt.Println("create card2---000000")

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

	fmt.Println(r.cards2)
	return nil
}

func (r *GameMan) Cards2String() string {

	return fmt.Sprintf("%s", r.cards2)
}

func (r *GameMan) GMPlayAllCard(player *Player) {
	for _, v := range r.cards {
		v.Play(player)
	}
}
