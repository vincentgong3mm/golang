package dominion

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
type MessageType int

const (
	MsgNone        MessageType = 0 + iota
	MsgGameMan                 // Player <- GameMan
	MsgPlayer                  // Player <- Player
	MsgOtherPlayer             // Player <- OtherPlayer
)
*/

type MessageTypeGameMan int

const (
	MsgPrepareGame MessageTypeGameMan = 0 + iota // 게임 시작
	MsgFinishGame                                // 플레이 종료. Victory 점수 계산 단계
	MsgCloseGame                                 // 모두 종료
	MsgYourTurn                                  // player turn
)

type MessageGameMan struct {
	Msg MessageTypeGameMan
}

type MessageTypePlay int

const (
	MsgPlayCard      MessageTypePlay = 0 + iota // 자신의 turn에 카드 플레이
	MsgOtherPlayCard                            // 다른 플레이가 카드 플레이, 나에게 영향을 미치는 경우 받음.
)

type ActionState int

const (
	DoAction ActionState = 0 + iota
	DoneAction
)

type MessagePlay struct {
	Msg    MessageTypePlay
	ThisID PlayerID
	CardID CardID

	Step   int
	IsDone ActionState // DoAction : step의 action 해야함, DoneAction : step의 action 완료됨
}

type GameMan struct {
	//cards map[CardID]Card

	//cards2 map[CardID]Actioner
	cards map[CardID]Actioner

	supply    *Supply
	trashPile *TrashPile

	players     map[PlayerID]*Player
	genPlayerID func() PlayerID

	input io.Reader
	inbuf bytes.Buffer
}

func init() {
	// Create and Get Log Instance
	GetLogInstance()

	logger.Println("import dominion/GameMan")
	fmt.Println("import dominion/GameMan")
}

func CreateNewGameMan() *GameMan {
	n := GameMan{}
	//n.cards = make(map[CardID]Card)
	//n.cards2 = make(map[CardID]Actioner)
	n.cards = make(map[CardID]Actioner)

	n.supply = CreateNewSupply()
	n.trashPile = CreateNewTrashPile()
	n.players = make(map[PlayerID]*Player)
	n.genPlayerID = NewPlayerIDGenerator()

	n.input = bufio.NewReader(os.Stdin)

	return &n
}

func (r *GameMan) SetInputFromBuffer() {
	r.input = bufio.NewReader(&r.inbuf)
}

func (r *GameMan) SetInputFromStdin() {
	r.input = os.Stdin
}

func (r *GameMan) WriteInBuffer(s string) {
	r.inbuf.Write([]byte(s))
}

/*
func (r *GameMan) createCard(cardID CardID, cardType []CardType, cost int, ability []Ability) Card {
	r.cards[cardID] = Card{name: "", CardID: cardID, cardType: cardType, cost: cost, Ability: ability}

	return r.cards[cardID]
}


func (r *GameMan) createCard2(cardID CardID, cardType []CardType, cost int, ability []Ability) Actioner {
	switch cardID {
	case Bandit:
		r.cards2[cardID] = &CardBandit{Card: Card{name: "", CardID: cardID, cardType: cardType, cost: cost, Ability: ability}}
	default:
		r.cards2[cardID] = &Card{name: "", CardID: cardID, cardType: cardType, cost: cost, Ability: ability}
	}

	return r.cards2[cardID]
}
*/

func (r *GameMan) createCardByID(cardID CardID) Actioner {
	// 카드의 상세 데이터 초기화는 각 카드의 struct의 Init을 호출해서 해야함.
	// 각 카드의 데이터는 별도로 분리하기 위함. 여기 두면 코드가 너무 복잡해짐.
	return r.createCard(cardID, []CardType{}, 0, []Ability{})
}

func (r *GameMan) createCard(cardID CardID, cardType []CardType, cost int, ability []Ability) Actioner {
	// 이미 생성되어 있으면 에러 또는 등록처리 하지 않아야함.
	//
	// -----------------------------------------------

	/*
		switch cardID {
		case Bandit:
			n := &CardBandit{}
			n.InitCard()
			r.cards[cardID] = n
		case Upgrade:
			n := &CardUpgrade{}
			n.InitCard()
			r.cards[cardID] = n
		default:
			r.cards[cardID] = &Card{name: "", CardID: cardID, cardType: cardType, cost: cost, Ability: ability}
		}

		return r.cards[cardID]
	*/
	return nil
}

func (r *GameMan) RegistCardToSuppy(t SupplySet, players int) {
	estate := 8
	duchy := 8
	province := 8
	curse := 8

	switch players {
	case 2:
		estate = 8 + players*3 // 3 coper per player
		duchy = 8
		province = 8
		curse = 8
	case 3:
		estate = 10 + players*3 // 3 coper per player
		duchy = 10
		province = 10
		curse = 12
	default:
		estate = 12 + players*3 // 3 coper per player
		duchy = 12
		province = 12
		curse = 18
	}

	switch t {
	case SetFirstGame:
		r.supply.RegistCard(Copper, 50)
		r.supply.RegistCard(Silver, 40)
		r.supply.RegistCard(Gold, 30)
		r.supply.RegistCard(Estate, estate)
		r.supply.RegistCard(Duchy, duchy)
		r.supply.RegistCard(Province, province)
		r.supply.RegistCard(Curse, curse)

		r.supply.RegistCard(Market, 10)
		r.supply.RegistCard(Festival, 10)
		r.supply.RegistCard(Smithy, 10)
		r.supply.RegistCard(Upgrade, 10)
		r.supply.RegistCard(Laboratory, 10)
		r.supply.RegistCard(Artisan, 10)
		r.supply.RegistCard(Chapel, 10)
		r.supply.RegistCard(Witch, 10)
		r.supply.RegistCard(CouncilRoom, 10)
		r.supply.RegistCard(Mine, 10)
		//r.supply.RegistCard(Cellar, 10)
	case SetBigMoney:
		r.supply.RegistCard(Copper, 50)
	}

}

func (r *GameMan) CreateNewPlayer(name string) *Player {
	playerID := r.genPlayerID()
	player := Player{name: name, ID: playerID}

	// init channel and go routine
	player.InitChan(r)

	// inser Pplayer Point to map
	r.players[playerID] = &player

	t, _ := r.players[playerID]

	r.gainBeginHandCard(t)

	// 처음 생성한 플레이어가 먼저 할 수 있도록 설정
	if player.ID == 1 {
		player.status = PlayerAction
	}

	return r.players[playerID]
}

func (r *GameMan) SendMessageToOtherPlayer(p *Player, msg *MessagePlay) {
	for _, v := range r.players {
		if p.ID != v.ID {
			v.SendPlayMessage(msg)
		}
	}
}

func (r *GameMan) GetPlayer(id PlayerID) *Player {
	p, exist := r.players[id]

	if exist == true {
		return p
	}

	return nil
}

func (r *GameMan) gainFromSupplyToDeck(id CardID, player *Player) bool {
	if r.supply.Pop(id) == true {
		player.GainCard(id, ToDeck)
		return true
	}

	return false
}

// BuyCard is gain a card from supply to discard pile, -1 buys.
func (r *GameMan) BuyCard(id CardID, p *Player) error {
	if p.buys <= 0 {
		return errors.New(fmt.Sprintf("NOTE:can't buy. buy count is %d", p.buys))
	}

	if r.supply.Pop(id) == true {
		p.buys--
		p.GainCard(id, ToDiscardPile)
		return nil
	}

	return errors.New(fmt.Sprintf("NOTE:%s isn't in Supply.", id))
}

//GainCardFromSupplyToDiscardPile is gain a card from supply to discard pile. ex) witch, gain a curse
func (r *GameMan) GainCardFromSupplyToDiscardPile(id CardID, p *Player) error {
	if r.supply.Pop(id) == true {
		p.buys--
		p.GainCard(id, ToDiscardPile)
		return nil
	}

	return errors.New(fmt.Sprintf("NOTE:%s isn't in Supply.", id))
}

func (r *GameMan) GainCardFromSupplyToHand(id CardID, p *Player) bool {
	if r.supply.Pop(id) == true {
		p.GainCard(id, ToHand)
		return true
	}

	return false
}

// GainCardFromSupplyToHandByIndex is Artisan's SpecialAbility
// GainCardFromSupplyToHandByIndex is Mine's SpecialAbility, add param CardType
func (r *GameMan) GainCardFromSupplyToHandByIndex(index int, p *Player, uptocost int, ctype CardType) error {
	// choose card's cost is up to uptocst.
	if cardID, exist := r.supply.GetCard(index); exist == true {
		card := r.cards[cardID]
		if card.GetCost() > uptocost {
			return errors.New(fmt.Sprintf("NOTE:%s's cost is not up to %d.", cardID, uptocost))
		}

		if card.IsType(ctype) == false {
			return errors.New(fmt.Sprintf("NOTE:%s's type is not match %s.", cardID, ctype))
		}

	} else {
		return errors.New(fmt.Sprintf("NOTE:Supply isn't %d index card.", index))
	}

	if cardID, exist := r.supply.PopByIndex(index); exist == true {
		p.GainCard(cardID, ToHand)

		return nil
	}

	return errors.New(fmt.Sprintf("NOTE:Supply isn't %d index card.", index))
}

// GainCardFromSupplyByIndex is Workshop's SpecialAbility
func (r *GameMan) GainCardFromSupplyByIndex(index int, p *Player, uptocost int) error {
	// choose card's cost is up to uptocst.
	if cardID, exist := r.supply.GetCard(index); exist == true {
		card := r.cards[cardID]
		if card.GetCost() > uptocost {
			return errors.New(fmt.Sprintf("NOTE:%s's cost is not up to %d.", cardID, uptocost))
		}
	} else {
		return errors.New(fmt.Sprintf("NOTE:Supply isn't %d index card.", index))
	}

	if cardID, exist := r.supply.PopByIndex(index); exist == true {
		p.GainCard(cardID, ToDiscardPile)
		return nil
	}

	return errors.New(fmt.Sprintf("NOTE:Supply isn't %d index card.", index))
}

//func (r *GameMan)

func (r *GameMan) gainBeginHandCard(player *Player) {
	// draw 7 copper`
	for i := 0; i < 7; i++ {
		r.gainFromSupplyToDeck(Copper, player)
	}

	// draw 3 estate
	for i := 0; i < 3; i++ {
		r.gainFromSupplyToDeck(Estate, player)
	}

	// for next turn init data and shuffle deck
	player.InitForNextTurn()

	// 처음 받음 deck을 섞어야함.
	player.deck.Shuffle()
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
	s := "=======================================\n"
	s += "GameMan Info\n"
	s += "________________________________________\n"
	s += "Card List\n"
	for _, v := range r.cards {
		s += v.String()
	}

	s += "________________________________________\n"
	s += r.supply.String()
	s += "\n"

	s += "________________________________________\n"
	s += r.trashPile.String()
	s += "\n"

	s += "________________________________________\n"
	s += "Player List\n"
	for _, v := range r.players {
		s += v.String()
	}

	return s
}

func (r GameMan) StringSupply() string {
	return r.supply.String()
}
func (r *GameMan) createCardEx(card Actioner) {
	card.InitCard()
	r.cards[card.GetCardID()] = card
}

func (r *GameMan) CreateAllCard() error {
	// Original Dominion
	r.createCardEx(&CardOriBase{Card{CardID: Copper}})
	r.createCardEx(&CardOriBase{Card{CardID: Silver}})
	r.createCardEx(&CardOriBase{Card{CardID: Gold}})
	r.createCardEx(&CardOriBase{Card{CardID: Estate}})
	r.createCardEx(&CardOriBase{Card{CardID: Duchy}})
	r.createCardEx(&CardOriBase{Card{CardID: Province}})
	r.createCardEx(&CardOriBase{Card{CardID: Curse}})

	r.createCardEx(&CardOriBase{Card{CardID: Village}})
	r.createCardEx(&CardOriBase{Card{CardID: Market}})
	r.createCardEx(&CardOriBase{Card{CardID: Smithy}})
	r.createCardEx(&CardOriBase{Card{CardID: Festival}})
	r.createCardEx(&CardOriBase{Card{CardID: Village}})
	r.createCardEx(&CardOriBase{Card{CardID: Laboratory}})
	r.createCardEx(&CardArtisan{})
	r.createCardEx(&CardChapel{})
	r.createCardEx(&CardCellar{})
	r.createCardEx(&CardWorkshop{})
	r.createCardEx(&CardWitch{})
	r.createCardEx(&CardCouncilRoom{})
	r.createCardEx(&CardMine{})

	// Intregue
	r.createCardEx(&CardUpgrade{})

	//r.createCard(Festival, []CardType{CardTypeAction}, 5, []Ability{{AbilityAddAction, 2}, {AbilityAddBuy, 1}, {AbilityAddCoin, 2}})
	/*
			r.createCardByID(Bandit)
			r.createCardByID(Smithy)
			r.createCardByID(Upgrade)

		s	r.createCard(Festival, []CardType{CardTypeAction}, 5,
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
	*/

	return nil
}

/*
func (r *GameMan) buyCard(p *Player, id CardID) error {
	// supply에서 하나 제거하고
	if r.supply.Pop(id) == true {
		// player의 discard pile에 하나 추가
		p.BuyCard(id)
		return nil
	}

	return errors.New(fmt.Sprintf("Not enough %s card in supply", id))
}
*/

func (r *GameMan) TrashCardFromHand(p *Player, index int) error {
	if cardID, err := p.handCards.RemoveCard(index); err != nil {
		return err
	} else {
		r.trashPile.AddCard(cardID)
	}

	return nil
}

func (r *GameMan) TrashTopCardFromDeck(p *Player, cnt int) (CardIDs, error) {
	/*
		_, error := p.RevealTopCardFromDeck(cnt)
		if error != nil {
			return error
		}
	*/

	cards, error := p.PopTopCardFromDeck(cnt)
	if error != nil {
		return CardIDs{}, error
	}

	for _, v := range cards {
		r.trashPile.AddCard(v)
	}

	return cards, nil
}

func (r *GameMan) GMPlayAllCard(player *Player) {
	/*
		for _, v := range r.cards {
			v.Play(player)
		}
	*/
}

func (r *GameMan) ReadInput(s ...string) (int, error) {
	fmt.Print(s)
	reader := bufio.NewReader(r.input)
	str, err := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")

	n, err := strconv.Atoi(str)
	return n, err
}
