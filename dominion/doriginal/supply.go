package doriginal

import "fmt"

// Card Dumy in supply
type Supply struct {
	cards map[CardID]int
}

func init() {
	fmt.Println("import d_original/supply")
}

func CreateNewSupply(name string) *Supply {
	n := Supply{}
	n.cards = make(map[CardID]int)

	return &n
}

func (r *Supply) RegistCarad(c Card, cnt int) {
	r.cards[c.CardID] = cnt
}
