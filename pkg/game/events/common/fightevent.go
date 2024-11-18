package common

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
	c "github.com/ercasta/allsoulsrun/pkg/game/common"
)

const (
	SIDE_CHARACTERS = "Characters"
	SIDE_MONSTERS   = "Monsters"
)

const FIGHT_EVENT = "Fight"

type FightEvent struct {
	Fight *c.Fight
}

func (f *FightEvent) GetType() engine.EventType { return FIGHT_EVENT }

func (f *FightEvent) Happen(t *engine.Timeline) {
	f.Fight = &c.Fight{}
	for _, fighter := range t.World.Characters {
		f.Fight.AddFighter(&fighter, SIDE_CHARACTERS)
	}

	var opp = engine.NewCharacter("Goblin", 1, 0, 50, 5, 5, 5, 5, 15, 0)
	f.Fight.AddFighter(&opp, SIDE_MONSTERS)
	var opp2 = engine.NewCharacter("Orc", 1, 0, 150, 15, 10, 10, 10, 20, 0)
	f.Fight.AddFighter(&opp2, SIDE_MONSTERS)

}

func Scheduled(t *engine.Timeline) {}

func (f *FightEvent) Cancel(t *engine.Timeline) {
	// Does nothing
}
