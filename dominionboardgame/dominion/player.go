package dominion

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 이렇게 까지는 필요하지 않음.
// player의 turn에 해야할 순서
// type Actioner interface {
// 	action() error
// 	buy() error
// 	clean() error
// 	draw() error
// }

type PlayerStatus int

const (
	PlayerNone                  PlayerStatus = 0 + iota
	PlayerStandBy                            // 대기 중, 다른 플레이어가 진행 중
	PlayerAction                             // 플레이, 액션카드 사용 가능
	PlayerActionWaitOtherPlayer              // 플레이, 액션카드 사용 후 다른 플레이어 효과처리 기다리는 중
	PlayerBuy                                // 두매단계
	PlayerCleanUp                            // 마지막 단계, Hand와 PlayingArea에 있는 카드 모두 DiscardPile로 이동
	PlayerDoAbilityOtherPlayer               // 플레이어로 인해서 대기 중인 플레이어의 효과 적용 중
	MaxPlayerStatusID
)

var PlayerStatusString = [...]string{
	"None",
	"StandBy",
	"Action",
	"ActionWaitOtherPlayer",
	"Buy",
	"CleanUp",
	"DoAbilityOtherPlayer",
}

func (r PlayerStatus) String() string {
	return PlayerStatusString[r%MaxPlayerStatusID]
}

// add counter -> Player.index
type Player struct {
	name            string
	ID              PlayerID
	deck            CardIDs // index 0 is the top card
	handCards       CardIDs
	HandCards       Cards
	cardPlayingArea CardIDs
	discardPile     CardIDs

	actions int
	buys    int
	coins   int

	status      PlayerStatus
	chanGameMan chan MessageGameMan
	chanPlay    chan MessagePlay
	returnCnt   int // 다른 유저로 부터 받아야 할 메시지 수
}

func init() {
	fmt.Println("import dominon/player")

}

func (p *Player) InitChan(g *GameMan) {
	p.chanGameMan = make(chan MessageGameMan)
	p.chanPlay = make(chan MessagePlay)

	go p.DoChanMessage(g)
}

func (p *Player) SendGameManMessage(msg *MessageGameMan) {
	p.chanGameMan <- *msg
}

func (p *Player) SendPlayMessage(msg *MessagePlay) {
	p.chanPlay <- *msg
}

func (p *Player) DoChanMessage(g *GameMan) {
	fmt.Println("Run DoChanMessage", p.name, p.ID)
	for {
		select {
		case msg := <-p.chanGameMan:
			fmt.Println("Receive From chanGameMan", g)
			p.DoGameManMessage(g, &msg)
		case msg := <-p.chanPlay:
			fmt.Println("Receive From Play", p)
			p.DoPlayMessage(g, &msg)
		}
	}

	fmt.Println("exit DoChanMessage", p.name, p.ID)
}

func (p *Player) DoGameManMessage(g *GameMan, msg *MessageGameMan) {

}

func (p *Player) DoPlayMessage(g *GameMan, msg *MessagePlay) {
	fmt.Println("DoPlayMessage", msg)

	card := g.cards[msg.CardID]

	switch msg.Msg {
	case MsgPlayCard:
		// 카드의 능력 실행
		card.DoAbility(p)
		card.DoSpecialAbility(p, g, msg)
	case MsgOtherPlayCard:
		// A->B : 공격 메시지를 상대방에게 보낸 경우, 상대방 처리
		if msg.IsDone == DoAction {
			card.DoOtherPlayer(p, g, msg)
			// B->A : 공격 메시지 처리 후 공격을 보낸 플레이어에게 보낸 결과 처리
		} else if msg.IsDone == DoneAction {
			card.AfterDoSpecialAbility(p, g, msg)
		}
	}
}

type PlayerID int

func NewPlayerIDGenerator() func() PlayerID {
	var next int
	return func() PlayerID {
		next++
		return PlayerID(next)
	}
}

func (r Player) String() string {
	s := ""
	s += fmt.Sprintf("@Player:%s(ID:%d)\n", r.name, r.ID)

	s += fmt.Sprintf("+Status(%s)\n", r.status)
	s += fmt.Sprintf("+Action(%d)\n", r.actions)
	s += fmt.Sprintf("+Buy(%d)\n", r.buys)
	s += fmt.Sprintf("+Coin(%d)\n", r.coins)
	s += fmt.Sprintf("+Deck")
	s += fmt.Sprintf("%s\n", r.deck)
	s += fmt.Sprintf("+Hand")
	s += fmt.Sprintf("%s\n", r.handCards)

	s += fmt.Sprintf("+CardPlayingArea")
	s += fmt.Sprintf("%s\n", r.cardPlayingArea)

	s += fmt.Sprintf("+DiscardPile")
	s += fmt.Sprintf("%s\n", r.discardPile)

	// test call ... GetLogInstance().Println(s)

	return s
}

func (r Player) StringHand() string {
	s := fmt.Sprintf("+Hand")
	s += fmt.Sprintf("%s", r.handCards)

	return s
}

func (r *Player) InitForNextTurn() {
	r.coins = 0
	r.buys = 1
	r.actions = 1

	r.status = PlayerStandBy

	// 처음 받을 때만 해야함.
	//r.deck.Shuffle()
}

func (r *Player) AddDiscardPileToDeck() {
	// shuffle discard pile
	r.discardPile.Shuffle()

	// add iscard pile to deck
	r.addCardsToDeckBottom(&r.discardPile)

	// empty discard pile
	r.discardPile = r.discardPile[0:0]
}

func (r *Player) addCardsToDeckBottom(cards *CardIDs) {
	r.deck = append(r.deck, *cards...)
}

func (r CardIDs) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
}

// Gain a card to
type GainedCard int

const (
	ToDiscardPile GainedCard = 0 + iota
	ToDeck
	ToHand
)

// GainCard is gain a card from Supply
func (r *Player) GainCard(id CardID, to GainedCard) {
	switch to {
	case ToDiscardPile:
		r.discardPile = append(r.discardPile, id)
	case ToDeck:
		r.deck = append(r.deck, id)
	case ToHand:
		r.handCards = append(r.handCards, id)
	}
}

func (r *Player) PutOnToTopDeck(id CardID) {
	tmpdeck := r.deck
	r.deck = r.deck[:0]

	r.deck = append(r.deck, id)
	r.deck = append(r.deck, tmpdeck...)
}

// Artisan : SpecialAbility
func (r *Player) PutCardFromHandToTopDeck(index int) error {
	if index >= len(r.handCards) {
		return errors.New(fmt.Sprintf("NOTE:Hand isn't %d index card.", index))
	}

	cardID := r.handCards[index]

	front := r.handCards[0:index]
	end := r.handCards[index+1 : len(r.handCards)]

	r.handCards = front
	r.handCards = append(r.handCards, end...)

	r.PutOnToTopDeck(cardID)

	return nil
}

// DrawCard is draw cards from deck to hand
func (r *Player) DrawCard(cnt int) (CardIDs, error) {
	if len(r.deck) < cnt {
		// add to deck
		r.AddDiscardPileToDeck()
	}

	if len(r.deck) < cnt {
		return CardIDs{}, errors.New(fmt.Sprintf("NOTE:not enough deck. deck is %d < %d", len(r.deck), cnt))
	}

	tmpCards := r.deck[0:cnt]
	r.deck = r.deck[cnt:len(r.deck)]
	r.handCards = append(r.handCards, tmpCards...)

	return tmpCards, nil
}

func (r *Player) PlayActionCard(index int) {
}

func (r *Player) PlayTreasureCard(index int) {
}

func (r *Player) CleanUp() {
	// empty hand cards to discardpile
	r.discardPile = append(r.discardPile, r.handCards...)

	// emptly hand cards
	r.handCards = r.handCards[:0]
}

/*
func (r *Player) BuyCard(card CardID) error {
	if r.buys <= 0 {
		return errors.New(fmt.Sprintf("can't buy. buy count is %d", r.buys))
	}

	// 한번 샀으면 구매 횟수 차감
	r.buys--
	r.GainCard(card, ToDiscardPile)

	return nil
}
*/

func (r *Player) BuyCardGM(card CardID) {
	r.GainCard(card, ToDiscardPile)
}
func (r *Player) GainCardGM(card CardID) {
	r.GainCard(card, ToHand)
}

func (r *Player) DiscardFromHand(index int) error {
	if index >= len(r.handCards) {
		return errors.New("NOTE:Invaild Hand Cards Index")
	}

	id := r.handCards[index]

	front := r.handCards[0:index]
	end := r.handCards[index+1 : len(r.handCards)]

	r.handCards = front
	r.handCards = append(r.handCards, end...)

	r.GainCard(id, ToDiscardPile)

	return nil
}

/*
// TranshCard is trash card to trash
func (r *Player) TrashCardFromHand(index int) error {
	if cardID, err := r.handCards.RemoveCard(int); err != nil {
		return err
	}

}
*/

func (r *Player) PlayCardFromHand(index int, gman *GameMan) error {
	if index >= len(r.handCards) {
		return errors.New("NOTE:Invaild Hand Cards Index")
	}
	if r.actions < 1 {
		return errors.New("NOTE:Player's actions is 0")
	}
	if r.status != PlayerAction {
		return errors.New(fmt.Sprintf("NOTE:Player status is not Action, current Status is %s", r.status))
	}

	// 핸드에서 카드 찾고
	cardID := r.handCards[index]

	_, exist := gman.cards[cardID]
	if exist == false {
		return errors.New(fmt.Sprintf("NOTE : %s is not registed.", cardID))
	}

	r.status = PlayerAction

	// 핸드에서 카드제거
	front := r.handCards[0:index]
	end := r.handCards[index+1 : len(r.handCards)]
	r.handCards = front
	r.handCards = append(r.handCards, end...)

	// 카드를 PlayingArea에 마지막에 추가
	r.cardPlayingArea = append(r.cardPlayingArea, cardID)

	// 액션 카드 사용할 횟수 차감
	r.actions--

	// 카드의 액션 실행하기 위해서 go routine으로 메시지 보냄
	// go routine에서 하는 이유
	//	- 다른 플레이어에게 영향을 주는 액션의 경우 나의 액션을 한 후 다른 플레이어 액션 완료를 기다린 후 진행해야함.
	//  - 예) Witch를 사용 후 +2 Card 후 다른 플레이어에게 메시지 보내고, 다른 플레이어가 Curse를 모두 받은 후 내가 다음 진행
	msg := MessagePlay{Msg: MsgPlayCard, ThisID: r.ID, CardID: cardID,
		Step: 0, IsDone: DoAction}
	r.SendPlayMessage(&msg)

	return nil
}

func (r *Player) RevealTopCardFromDeck(cnt int) (CardIDs, error) {
	if len(r.deck) < cnt {
		// add to deck
		r.AddDiscardPileToDeck()
	}

	if len(r.deck) < cnt {
		return CardIDs{}, errors.New(fmt.Sprintf("NOTE:not enough deck. deck is %d < %d", len(r.deck), cnt))
	}

	cards := r.deck[0:cnt]

	return cards, nil
}

func (r *Player) PopTopCardFromDeck(cnt int) (CardIDs, error) {
	if len(r.deck) < cnt {
		// add to deck
		r.AddDiscardPileToDeck()
	}

	if len(r.deck) < cnt {
		return CardIDs{}, errors.New(fmt.Sprintf("NOTE:not enough deck. deck is %d < %d", len(r.deck), cnt))
	}

	cards := r.deck[0:cnt]
	r.deck = r.deck[cnt:len(r.deck)]

	return cards, nil
}
