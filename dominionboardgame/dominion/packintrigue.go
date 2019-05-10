package dominion

import (
	"fmt"
)

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

func (r *CardUpgrade) DoAbility(p *Player) {
	fmt.Sprintf("DoAbility -> %s", r.String())
}

func (r *CardUpgrade) Draw(p *Player) {
}

func (r *CardUpgrade) AddBuy(p *Player) {
	cnt, _ := r.GetAbilityCount(AbilityAddBuy)
	p.buys += cnt
}

func (r *CardUpgrade) AddAction(p *Player) {
	cnt, _ := r.GetAbilityCount(AbilityAddAction)
	p.actions += cnt
}

func (r *CardUpgrade) String() string {
	//return r.Card.String()
	return "0000000000000000"
}
