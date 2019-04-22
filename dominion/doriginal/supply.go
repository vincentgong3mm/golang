package doriginal

import "fmt"

// Card Dumy in supply
type Supply struct {
	cards map[CardId]int
}

func init() {
	fmt.Println("import d_original/supply")
}

func CreateNewSupply(name string) *Supply {
	n := Supply{}
	n.cards = make(map[CardId]int)

	return &n
}

func (r *Supply) RegistCarad(c Card, cnt int) {
	r.cards[c.cardId] = cnt
}
