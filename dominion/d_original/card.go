package d_original

import "fmt"

type CardUnique int

type CardType int

const (
	CardTypeAction = 0 + iota
	CardTypeTreasure
	CardTypeVictory
)

type Card struct {
	name       string
	cardUnique CardUnique
	cardType   []CardType
	cost       int
}

func init() {
	fmt.Println("import d_original/card")
}

/*
func CreateNewCard(name string) *Card {
	nc := Card{name: name}

	return &nc
}
*/

func CreateNewCard(name string, i CardUnique, cardType []CardType, cost int) *Card {
	n := Card{name: name, cardUnique: i, cardType: cardType, cost: cost}

	return &n
}
