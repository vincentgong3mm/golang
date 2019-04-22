package doriginal

import "fmt"

type CardId int

type CardType int

const (
	CardTypeAction = 0 + iota
	CardTypeTreasure
	CardTypeVictory
)

const (
	InvaildCardId = -1
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
	name     string
	cardId   CardId
	cardType []CardType
	cost     int
}

func (r Card) String() string {
	return fmt.Sprintf("%s %d %s %d", r.name, r.cardId, r.cardType, r.cost)
}

func (r Card) TermString() string {

	ln := make([]byte, 0, 40)

	width := 20
	lenstr := len(r.name)
	lenspace := (width - lenstr) / 2

	ln = append(ln, '|')

	space := make([]byte, lenspace)
	space2 := make([]byte, lenspace-1)
	ln = append(ln, space...)

	tmpName := []byte(r.name)
	ln = append(ln, tmpName...)

	if lenstr%2 == 0 {
		ln = append(ln, space...)
	} else {
		ln = append(ln, space2...)
	}

	ln = append(ln, '|')

	return string(ln)

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

func NewCardIdGenerator() func() CardId {
	var next int
	return func() CardId {
		next++
		return CardId(next)
	}
}

// CreateNewCarad is
//func CreateNewCard(name string, cardType []CardType, cost int) *Card {
//	n := Card{name: name, cardId: InvaildCardId, cardType: cardType, cost: cost}
//	return &n
//}
