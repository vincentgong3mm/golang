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
