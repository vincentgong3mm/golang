package doriginal

import "fmt"

// Card Dumy in supply
type Supply struct {
	cards map[CardID]int
}

type SupplySet int

const (
	SetFirstGame SupplySet = 0 + iota
	SetBigMoney
	SetInteraction
)

func init() {
	fmt.Println("import d_original/supply")
}

func CreateNewSupply() *Supply {
	n := Supply{}
	n.cards = make(map[CardID]int)

	return &n
}

func (r *Supply) RegistCard(c CardID, cnt int) {
	r.cards[c] = cnt
}
