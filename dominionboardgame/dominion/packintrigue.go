package dominion

type CardUpgrade struct {
	Card
}

func (r *CardUpgrade) InitCard() {
	r.CardID = Upgrade
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{{AbilityAddAction, 1}, {AbilityAddCard, 1}}
}

func (r *CardUpgrade) DoAbility(p *Player) {
	r.Card.DoAbility(p)
}

/*
func (r *CardUpgrade) AddCard(p *Player) {
	cnt, _ := r.GetAbilityCount(AbilityAddCard)

	p.DrawCard(cnt)
}

func (r *CardUpgrade) AddBuy(p *Player) {
	cnt, _ := r.GetAbilityCount(AbilityAddBuy)
	p.buys += cnt
}

func (r *CardUpgrade) AddAction(p *Player) {
	cnt, _ := r.GetAbilityCount(AbilityAddAction)
	p.actions += cnt
}
*/

func (r *CardUpgrade) String() string {
	return r.Card.String()
}
