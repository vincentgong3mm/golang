package doriginal

import (
	"fmt"
)

type Actioner interface {
	Draw(p *Player)
	AddBuy(p *Player)
	AddAction(p *Player)
	//DoSpecailACtion()
}

type CardThief struct {
	comment int
	Card
}

func (r *CardThief) Draw(p *Player) {
	fmt.Println("CardThief:name", r)
	r.comment = 999
}

func (r *CardThief) AddBuy(p *Player) {
}

func (r *CardThief) AddAction(p *Player) {
}

func (r CardThief) String() string {
	return fmt.Sprintf("%d<<<%s", r.comment, r.Card.String())
}

/*
func (r CardThief) Draw(p *Player) {
	fmt.Println("CardThief:name", r)

	r.cost = 1000
	r.name = "aaaaaa"

	r.comment = 999

	fmt.Println(r)
}

func (r CardThief) AddBuy(p *Player) {
}

func (r CardThief) AddAction(p *Player) {
}
*/

/*
type Smithy struct {
	*Card
}
*/
