package dominion

import (
	"fmt"
	"time"
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

func (r *CardArtisan) DoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
	for {
		fmt.Println(">>>>", p.StringHand())
		fmt.Println(">>>>", g.StringSupply())
		index, _ := g.ReadInput(r.CardID.String(), ": Gain a card to your hand consting up to 5, choose supply's index #")
		if err := g.GainCardFromSupplyToHandByIndex(index, p, 5, CardTypeNone); err != nil {
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

func (r *CardChapel) DoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
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

func (r *CardCellar) DoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
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

type CardWorkshop struct {
	Card
}

func (r *CardWorkshop) InitCard() {
	r.CardID = Workshop
	r.cardType = []CardType{CardTypeAction}
	r.cost = 3
	r.Ability = []Ability{}
}

func (r *CardWorkshop) DoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
	for {
		fmt.Println(">>>>", g.StringSupply())
		index, err := g.ReadInput(r.CardID.String(), ": Gain a card costing up to 4. Choose card's index in supply#")
		// input '' enter is that don't trash card
		if err != nil {
			break
		}
		if err := g.GainCardFromSupplyByIndex(index, p, 4); err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
}

func (r *CardWorkshop) String() string {
	return r.Card.String()
}

type CardWitch struct {
	Card
}

func (r *CardWitch) String() string {
	return r.Card.String()
}

func (r *CardWitch) InitCard() {
	r.CardID = Witch
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{{AbilityAddCard, 2}}
}

func (r *CardWitch) DoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
	fmt.Println(">>>>", g.StringSupply())
	fmt.Println("Each other player gains a Curse card.")

	//otherMsg := MessagePlay{Msg: MsgOtherPlayCard, CardID: r.CardID, Step: 0, IsDone: DoAction}
	otherMsg := *msg
	otherMsg.Msg = MsgOtherDoEffect

	p.returnCnt = len(g.players) - 1

	// 다른 플레이어로 부터 처리 결과 메시지 오기를 기다리는 상태로 변경
	p.status = PlayerActionWaitOtherPlayer

	g.SendMessageToOtherPlayer(p, &otherMsg)
}

func (r *CardWitch) AfterDoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
	p.returnCnt--

	// 다른 플레이어로 부터 Curse 획득 모두 받았으면
	if p.returnCnt == 0 {
		p.status = PlayerAction
	}
}

func (r *CardWitch) DoOtherPlayer(p *Player, g *GameMan, msg *MessagePlay) {
	fmt.Println(">>>>", g.StringSupply())
	fmt.Println("Gain a Curse card.")

	if err := g.GainCardFromSupplyToDiscardPile(Curse, p); err != nil {
		fmt.Println(err)
	}

	if thisPlayer := g.GetPlayer(msg.ThisID); thisPlayer != nil {
		// 구조체 복사해서 후 데이터 변경 후
		thisMsg := *msg
		thisMsg.Msg = MsgOtherDoneEffect
		thisPlayer.SendPlayMessage(&thisMsg)
	}
}

type CardCouncilRoom struct {
	Card
}

func (r *CardCouncilRoom) String() string {
	return r.Card.String()
}

func (r *CardCouncilRoom) InitCard() {
	r.CardID = CouncilRoom
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{{AbilityAddCard, 4}, {AbilityAddBuy, 1}}
}

func (r *CardCouncilRoom) DoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
	fmt.Println(">>>>", g.StringSupply())
	fmt.Println("Each other player draw a card.")

	//otherMsg := MessagePlay{Msg: MsgOtherPlayCard, CardID: r.CardID, Step: 0, IsDone: DoAction}
	otherMsg := *msg
	otherMsg.Msg = MsgOtherDoEffect

	p.returnCnt = len(g.players) - 1

	// 다른 플레이어로 부터 처리 결과 메시지 오기를 기다리는 상태로 변경
	p.status = PlayerActionWaitOtherPlayer

	g.SendMessageToOtherPlayer(p, &otherMsg)
}

func (r *CardCouncilRoom) AfterDoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
	p.returnCnt--

	// 다른 플레이어로 부터 Curse 획득 모두 받았으면
	if p.returnCnt == 0 {
		p.status = PlayerAction
	}
}

func (r *CardCouncilRoom) DoOtherPlayer(p *Player, g *GameMan, msg *MessagePlay) {
	fmt.Println(">>>>", g.StringSupply())
	fmt.Println("Draw a card.")

	if _, err := p.DrawCard(1); err != nil {
		fmt.Println(err)
	}

	if thisPlayer := g.GetPlayer(msg.ThisID); thisPlayer != nil {
		// 구조체 복사해서 후 데이터 변경 후
		thisMsg := *msg
		thisMsg.Msg = MsgOtherDoneEffect
		thisPlayer.SendPlayMessage(&thisMsg)
	}
}

type CardMine struct {
	Card
}

func (r *CardMine) String() string {
	return r.Card.String()
}

func (r *CardMine) InitCard() {
	r.CardID = Mine
	r.cardType = []CardType{CardTypeAction}
	r.cost = 5
	r.Ability = []Ability{}
}

func (r *CardMine) DoSpecialAbility(p *Player, g *GameMan, msg *MessagePlay) {
	fmt.Println(">>>>", g.StringSupply())

	for {
		fmt.Println(">>>>", p.StringHand())
		fmt.Println("Trash a Treasure card from you hand. Gain a Treasure card costing up to 3 more; put it into your hand.")
		index, err := g.ReadInput(r.CardID.String(), ": choose hand's index #")

		if err == nil {
			cardID, err := p.handCards.GetCardID(index)

			if err == nil {
				card := g.cards[cardID]

				// Treasure Card만 버릴 수 있음
				if card.IsType(CardTypeTreasure) == true {
					if err := g.TrashCardFromHand(p, index); err != nil {
						fmt.Println(err)
					} else {
						upto := card.GetCost()
						upto += 3

						// supply index input 받기
						supplyIndex, err := g.ReadInput(r.CardID.String(), fmt.Sprintf(": Gain a treasure card to your hand up to %d, choose supply's index #", upto))

						// input '' enter is that don't trash card
						if err != nil {
							break
						}

						// trasnCardCost 이하의 treasure 카드 hand로 가져오기
						if err := g.GainCardFromSupplyToHandByIndex(supplyIndex, p, upto, CardTypeTreasure); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(">>>>", p)
							fmt.Println(">>>>", r.GetCardID(), "Complete.")
							//fmt.Printf(">>>> %s Complete.\n", r.GetCardID())
							break
						}
					}
				} else {
					fmt.Printf("NOTE:%s is not treasure card.", card)
				}
			}
		} else {
			// 버릴 treasure 카드가 없어 enter입력하면 광산 액션 종료
			fmt.Printf("NOTE:didn't choose card!")
			break
		}
	}

	for i := 0; i < 3; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Mine Sleep")
	}

	//g.Done()
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
