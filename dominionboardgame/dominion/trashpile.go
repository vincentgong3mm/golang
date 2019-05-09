package dominion

import (
	"fmt"
	"strings"
)

type TrashPile struct {
	cards map[CardID]int
}

func (r TrashPile) String() string {
	s := "|"
	for cardID, cnt := range r.cards {
		s += fmt.Sprintf("%s#%d|", cardID.String(), cnt)
	}

	sline := strings.Repeat("-", len(s))

	return "#TrashPile\n" + sline + "\n" + s + "\n" + sline + "\n"
}

func (r *TrashPile) AddCard(id CardID) {
	if _, exist := r.cards[id]; exist == true {
		r.cards[id]++
	} else {
		r.cards[id] = 0
	}
}

func (r *TrashPile) PopCard(id CardID) bool {
	cnt, exist := r.cards[id]
	if exist == true && cnt > 0 {
		cnt--
		r.cards[id] = cnt
	}

	return exist
}
