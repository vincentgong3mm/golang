package dominion

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
