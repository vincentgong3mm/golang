package dominion

import (
	"fmt"
	"strings"
)

type DumyIndex int

func NewDumyIDGenerator() func() DumyIndex {
	var next int
	return func() DumyIndex {
		next++
		return DumyIndex(next)
	}
}

type CardDumy struct {
	cardID CardID
	cnt    int
}

// Card Dumy in supply
type Supply struct {
	cards        []CardDumy
	genDumyIndex func() DumyIndex
}

func (r Supply) String() string {
	s := "|"
	for i, v := range r.cards {
		s += fmt.Sprintf("%d#%s(%d)|", i, v.cardID, v.cnt)
	}

	sline := strings.Repeat("-", len(s))

	return "+Supply\n" + sline + "\n" + s + "\n" + sline
}

type SupplySet int

const (
	SetFirstGame SupplySet = 0 + iota
	SetBigMoney
	SetInteraction
)

func init() {
	fmt.Println("import dominion/supply")
}

func CreateNewSupply() *Supply {
	n := Supply{}
	n.cards = make([]CardDumy, 0, 10)

	//n.cards = make(map[CardID]CardDumy)
	//n.genDumyIndex = NewDumyIDGenerator()

	return &n
}

func (r *Supply) RegistCard(c CardID, cnt int) {
	r.cards = append(r.cards, CardDumy{cardID: c, cnt: cnt})
	//r.cards[c] = CardDumy{index: r.genDumyIndex(), cnt: cnt}
}

func (r *Supply) Pop(id CardID) bool {
	return r.PopByCardID(id)
}

func (r *Supply) PopByCardID(id CardID) bool {
	for i, v := range r.cards {
		if v.cardID == id {
			v.cnt--
			r.cards[i] = v
			return true
		}
	}

	return false
}

func (r *Supply) PopByIndex(index int) (CardID, bool) {
	if index >= len(r.cards) {
		return 0, false
	}

	r.cards[index].cnt--

	return r.cards[index].cardID, true
}

func (r *Supply) GetCard(index int) (CardID, bool) {
	if index >= len(r.cards) {
		return -1, false
	}

	return r.cards[index].cardID, true
}
