package dominion

import (
	"fmt"
	"strings"
)

// Card Dumy in supply
type Supply struct {
	cards map[CardID]int
}

func (r Supply) String() string {
	s := "|"
	for cardID, cnt := range r.cards {
		s += fmt.Sprintf("%s#%d|", cardID.String(), cnt)
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
	n.cards = make(map[CardID]int)

	return &n
}

func (r *Supply) RegistCard(c CardID, cnt int) {
	r.cards[c] = cnt
}

func (r *Supply) Pop(id CardID) bool {
	cnt, exist := r.cards[id]
	if exist == true && cnt > 0 {
		cnt--
		r.cards[id] = cnt
	}

	return exist
}
