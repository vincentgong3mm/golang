package d_original

import "fmt"

// Card Dumy in supply
type Supply struct {
	cards map[CardUnique]int
}

func init() {
	fmt.Println("import d_original/supply")
}

func CreateNewSupply(name string) *Supply {
	n := Supply{}
	n.cards = make(map[CardUnique]int)

	return &n
}

func (r *Supply) RegistCarad(c Card, cnt int) {
	r.cards[c.cardUnique] = cnt
}
