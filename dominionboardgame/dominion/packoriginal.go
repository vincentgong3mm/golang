package dominion

import (
	"fmt"
)

type CardOriBase struct {
	Card
}

func (r *CardOriBase) InitCard() {
	r.name = ""
	// CardID is init when struct init.
	//r.CardID = xxxx

	switch r.CardID {
	case Copper:
		r.cardType = []CardType{CardTypeTreasure}
		r.cost = 0
		r.Ability = []Ability{{AbilityAddCoin, 1}}
	case Silver:
		r.cardType = []CardType{CardTypeTreasure}
		r.cost = 3
		r.Ability = []Ability{{AbilityAddCoin, 2}}
	case Gold:
		r.cardType = []CardType{CardTypeTreasure}
		r.cost = 6
		r.Ability = []Ability{{AbilityAddCoin, 3}}
	case Estate:
		r.cardType = []CardType{CardTypeVictory}
		r.cost = 2
		r.Ability = []Ability{{AbilityAddVictory, 1}}
	case Duchy:
		r.cardType = []CardType{CardTypeVictory}
		r.cost = 5
		r.Ability = []Ability{{AbilityAddVictory, 5}}
	case Province:
		r.cardType = []CardType{CardTypeVictory}
		r.cost = 8
		r.Ability = []Ability{{AbilityAddVictory, 6}}
	case Curse:
		r.cardType = []CardType{CardTypeVictory}
		r.cost = 0
		r.Ability = []Ability{{AbilityAddVictory, -1}}
	case Village:
		r.cardType = []CardType{CardTypeAction}
		r.cost = 4
		r.Ability = []Ability{{AbilityAddAction, 2}, {AbilityAddCard, 1}}
	case Festival:
		r.cardType = []CardType{CardTypeAction}
		r.cost = 5
		r.Ability = []Ability{{AbilityAddAction, 2}, {AbilityAddBuy, 1}, {AbilityAddCoin, 2}}
	case Smithy:
		r.cardType = []CardType{CardTypeAction}
		r.cost = 4
		r.Ability = []Ability{{AbilityAddCard, 3}}
	case Market:
		r.cardType = []CardType{CardTypeAction}
		r.cost = 4
		r.Ability = []Ability{{AbilityAddBuy, 1}, {AbilityAddAction, 1}, {AbilityAddCard, 1}, {AbilityAddCoin, 1}}
	case Laboratory:
		r.cardType = []CardType{CardTypeAction}
		r.cost = 5
		r.Ability = []Ability{{AbilityAddAction, 1}, {AbilityAddCard, 2}}
		/*
			case Village:
				r.cardType = []CardType{CardTypeAction}
				r.cost = 4
				r.Ability = []Ability{{AbilityAddAction, 2}, {AbilityAddCard, 1}}
		*/
	default:
	}
}

func (r *CardOriBase) DoAbility(p *Player) {
	r.Card.DoAbility(p)
}

func (r *CardOriBase) Draw(p *Player) {
}

func (r *CardOriBase) AddBuy(p *Player) {
}

func (r *CardOriBase) AddAction(p *Player) {
}

func (r *CardOriBase) String() string {
	return r.Card.String()
}

type CardArtisan struct {
	Card
}

func (r *CardArtisan) InitCard() {
	r.CardID = Artisan
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{}
}

func (r *CardArtisan) DoSpecialAbility(p *Player, g *GameMan) {
	for {
		fmt.Println(">>>>", p.StringHand())
		fmt.Println(">>>>", g.StringSupply())
		index, _ := g.ReadInput(r.CardID.String(), ": Gain a card to your hand consting up to 5, choose supply's index #")
		if err := g.GainCardFromSupplyToHandByIndex(index, p, 5); err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	for {
		fmt.Println(">>>>", p.StringHand())
		cardIndexInHand, _ := g.ReadInput(r.CardID.String(), ": Put a card from your hand onto your deck, choose hand's index #")
		if err := p.PutCardFromHandToTopDeck(cardIndexInHand); err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

}

func (r *CardArtisan) String() string {
	return r.Card.String()
}

type CardChapel struct {
	Card
}

func (r *CardChapel) InitCard() {
	r.CardID = Chapel
	r.cardType = []CardType{CardTypeAction}
	r.cost = 2
	r.Ability = []Ability{}
}

func (r *CardChapel) DoSpecialAbility(p *Player, g *GameMan) {
	for i := 0; i < 4; {
		fmt.Println(">>>>", p.StringHand())
		index, err := g.ReadInput(r.CardID.String(), ": Trash up to 4 cards from your hand, choose card's index #")

		// input '' enter is that don't trash card
		if err != nil {
			break
		}

		if err := g.TrashCardFromHand(p, index); err != nil {
			fmt.Println(err)
		} else {
			i++
		}
	}
}

func (r *CardChapel) String() string {
	return r.Card.String()
}

type CardCellar struct {
	Card
}

func (r *CardCellar) InitCard() {
	r.CardID = Cellar
	r.cardType = []CardType{CardTypeAction}
	r.cost = 2
	r.Ability = []Ability{{AbilityAddAction, 1}}
}

func (r *CardCellar) DoSpecialAbility(p *Player, g *GameMan) {
	for i := 0; i < 4; {
		fmt.Println(">>>>", p.StringHand())
		index, err := g.ReadInput(r.CardID.String(), ": Discard any number of cards.+1 Card per card discarded. Choose card's index to discard#")
		// input '' enter is that don't trash card
		if err != nil {
			break
		}

		if err := p.DiscardFromHand(index); err != nil {
			fmt.Println(err)
		} else {
			i++
		}
	}
}

func (r *CardCellar) String() string {
	return r.Card.String()
}

/*
type CardBandit struct {
	Card
}

func (r *CardBandit) InitCard() {
	r.name = ""
	r.CardID = Bandit
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{}
}

func (r *CardBandit) DoAbility(p *Player) {
}

func (r *CardBandit) Draw(p *Player) {
}

func (r *CardBandit) AddBuy(p *Player) {
}

func (r *CardBandit) AddAction(p *Player) {
}

func (r *CardBandit) String() string {
	return r.Card.String()
}
*/
