package dominion

type CardUpgrade struct {
	Card
}

func (r *CardUpgrade) InitCard() {
	logger.Println("InitCard CardUpgrade")

	r.CardID = Upgrade
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{{AbilityAddAction, 1}, {AbilityAddCard, 1}}
}

func (r *CardUpgrade) DoAbility(p *Player) {
	r.Card.DoAbility(p)
	r.DoSpecialAbility()
}

func (r *CardUpgrade) DoSpecialAbility() {
}

func (r *CardUpgrade) String() string {
	return r.Card.String()
}
