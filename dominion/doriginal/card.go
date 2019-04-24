package doriginal

import (
	"fmt"
)

// CardID is
type CardID int

// CardType is	Action, Treasure, Victory
type CardType int

const (
	CardTypeAction = 0 + iota
	CardTypeTreasure
	CardTypeVictory
	CardTypeCurse
)

const (
	InvaildCardID = -1
)

const (
	CardWidth = 20
)

var CardTypeString = [...]string{
	"Action",
	"Treasure",
	"Victory",
	"Curse",
}

func (r CardType) String() string {
	return CardTypeString[r%3]
}

type Card struct {
	name     string
	CardID   CardID
	cardType []CardType
	cost     int
}

func (r Card) String() string {
	return fmt.Sprintf("%s(ID:%d)\n\tcost(%d)\n\tType%s\t\n", r.name, r.CardID, r.cost, r.cardType)
}

func (r Card) TermString() string {

	ct := fmt.Sprintf("%s", r.cardType)
	//ct = strings.TrimLeft(ct, "[")
	//ct = strings.TrimRight(ct, "]")

	return ConvertTermString(CardWidth, r.name) + "\n" + ConvertTermString(CardWidth, ct) + "\n"

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

func NewCardIDGenerator() func() CardID {
	var next int
	return func() CardID {
		next++
		return CardID(next)
	}
}

// CreateNewCarad is
//func CreateNewCard(name string, cardType []CardType, cost int) *Card {
//	n := Card{name: name, CardID: InvaildCardID, cardType: cardType, cost: cost}
//	return &n
//}
