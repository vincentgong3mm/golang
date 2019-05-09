package dominion

type CardUpgrade struct {
	Card
}

func (r *CardUpgrade) InitCard() {
	r.name = ""
	r.CardID = Upgrade
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{{AbilityAddAction, 1}, {AbilityAddCard, 1}}
}

func (r *CardUpgrade) Draw(p *Player) {
}

func (r *CardUpgrade) AddBuy(p *Player) {
}

func (r *CardUpgrade) AddAction(p *Player) {
}

func (r *CardUpgrade) String() string {
	return r.Card.String()
}
