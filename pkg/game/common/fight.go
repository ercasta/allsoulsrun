package common

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
)

type Side string

const FightComponentType engine.ComponentType = "Fight"
const EntityFightComponentType engine.ComponentType = "EntityFight"

const (
	SIDE_CHARACTERS = "Characters"
	SIDE_MONSTERS   = "Monsters"
)

type Fight struct {
	sides map[Side][]engine.EntityID
}

// This component is attached / detached to entities as they enter / leave a fight
type EntityFight struct {
	FightId engine.EntityID
}

func (f EntityFight) GetComponentType() engine.ComponentType { return EntityFightComponentType }

func (f Fight) GetComponentType() engine.ComponentType { return FightComponentType }

// func (c Fight) PersistAll(ch []engine.ComponentHistory) {
// 	for _, component := range ch {
// 		if _, ok := (component.Component).(CharacterExperience); ok {
// 			// TODO Write all to file
// 			fmt.Println("Write Placeholder")
// 		}
// 	}
// }

func (f *Fight) AddFighter(fighter engine.EntityID, side Side) {
	if f.sides == nil {
		f.sides = make(map[Side][]engine.EntityID)
	}
	f.sides[side] = append(f.sides[side], fighter)
}

func (fight *Fight) IsInFight(fighter engine.EntityID) bool {
	for _, fighters := range fight.sides {
		for _, f := range fighters {
			if f == fighter {
				return true
			}
		}
	}
	return false
}

func (f *Fight) GetFighters() []engine.EntityID {
	var fighters []engine.EntityID
	for _, side := range f.sides {
		fighters = append(fighters, side...)
	}
	return fighters
}

func (fight *Fight) GetOpponents(c engine.EntityID) []engine.EntityID {
	// TODO optimize using a map
	for side, fighters := range fight.sides {
		if side == SIDE_CHARACTERS {
			for _, f := range fighters {
				if f == c {
					return fight.sides[SIDE_MONSTERS]
				}
			}
		}
		if side == SIDE_MONSTERS {
			for _, f := range fighters {
				if f == c {
					return fight.sides[SIDE_CHARACTERS]
				}
			}

		}
	}
	return nil
}

func (fight *Fight) RemoveFighter(fighter engine.EntityID) {
	for side, fighters := range fight.sides {
		for i, f := range fighters {
			if f == fighter {
				fighters = append(fighters[:i], fighters[i+1:]...)
				fight.sides[side] = fighters
				return
			}
		}
	}
}
