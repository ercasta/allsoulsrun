package common

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
)

type Side string

const (
	SIDE_CHARACTERS = "Characters"
	SIDE_MONSTERS   = "Monsters"
)

type Fight struct {
	sides map[Side][]engine.Character
}

func (fight *Fight) RemoveFighter(fighter *engine.Character) {
	for side, fighters := range fight.sides {
		for i, f := range fighters {
			if &f == fighter {
				fighters = append(fighters[:i], fighters[i+1:]...)
				fight.sides[side] = fighters
				return
			}
		}
	}
}
