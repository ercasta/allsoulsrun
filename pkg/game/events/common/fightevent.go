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
	sides map[c.Side][]engine.Character
}

func (f *FightEvent) GetType() engine.EventType { return FIGHT_EVENT }

func (f *FightEvent) Happen(t *engine.Timeline) {
	if f.sides == nil {
		f.sides = make(map[c.Side][]engine.Character)
	}

	f.sides[SIDE_CHARACTERS] = t.World.Characters
	f.sides[SIDE_MONSTERS] =
		[]engine.Character{
			engine.NewCharacter("Goblin", 1, 0, 50, 5, 5, 5, 5, 30, 0),
			engine.NewCharacter("Orc", 1, 0, 150, 15, 10, 10, 10, 70, 0),
		}

	// Roll for initiative and create initial events.
	t.AddEvent(&AttackEvent{Attacker: &f.sides[SIDE_CHARACTERS][0], Attacked: &f.sides[SIDE_MONSTERS][0], Fight: f.Fight}, 2000)
	t.AddEvent(&AttackEvent{Attacker: &f.sides[SIDE_MONSTERS][0], Attacked: &f.sides[SIDE_CHARACTERS][0], Fight: f.Fight}, 2500)
}

func Scheduled(t *engine.Timeline) {}

func (f *FightEvent) Cancel(t *engine.Timeline) {
	// Does nothing
}
