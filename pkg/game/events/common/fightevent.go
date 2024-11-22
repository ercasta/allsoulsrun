package common

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
)

const (
	SIDE_CHARACTERS = "Characters"
	SIDE_MONSTERS   = "Monsters"
)

const FIGHT_EVENT = "Fight"

type FightEvent struct {
	Fight e.EntityID
}

type FightEventListener struct{}

func (f FightEvent) GetType() e.EventType { return FIGHT_EVENT }
