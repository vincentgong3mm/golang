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
	index DumyIndex
	cnt   int
}

// Card Dumy in supply
type Supply struct {
	cards        map[CardID]CardDumy
	genDumyIndex func() DumyIndex
}

func (r Supply) String() string {
	s := "|"
	for cardID, dumy := range r.cards {
		s += fmt.Sprintf("%d#%s(%d)|", dumy.index, cardID.String(), dumy.cnt)
	}

	sline := strings.Repeat("-", len(s))

	return "#Supply\n" + sline + "\n" + s + "\n" + sline + "\n"
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
	n.cards = make(map[CardID]CardDumy)

	n.genDumyIndex = NewDumyIDGenerator()

	return &n
}

func (r *Supply) RegistCard(c CardID, cnt int) {
	r.cards[c] = CardDumy{index: r.genDumyIndex(), cnt: cnt}
}

func (r *Supply) Pop(id CardID) bool {
	dumy, exist := r.cards[id]
	if exist == true && dumy.cnt > 0 {
		dumy.cnt--
		r.cards[id] = CardDumy{index: dumy.index, cnt: dumy.cnt}
	}

	return exist
}
