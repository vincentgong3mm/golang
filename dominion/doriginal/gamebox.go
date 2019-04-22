package doriginal

import "fmt"

type GameBox struct {
	cards map[CardUnique]Card
}

func init() {
	fmt.Println("import d_original/gamebox")
}

func CreateNewGameBox() *GameBox {
	n := GameBox{}
	n.cards = make(map[CardUnique]Card)

	return &n
}

func (r *GameBox) CreateCard(c Card) {
	r.cards[c.cardUnique] = c
}
