package dominion

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
