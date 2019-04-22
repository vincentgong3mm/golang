package doriginal

import "fmt"

type CardUnique int

type CardType int

const (
	CardTypeAction = 0 + iota
	CardTypeTreasure
	CardTypeVictory
)

var CardTypeString = [...]string{
	"Action",
	"Treasure",
	"Victory",
}

func (r CardType) String() string {
	return CardTypeString[r%3]
}

type Card struct {
	name       string
	cardUnique CardUnique
	cardType   []CardType
	cost       int
}

func (r Card) String() string {
	return fmt.Sprintf("%s %d %s %d", r.name, r.cardUnique, r.cardType, r.cost)
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
