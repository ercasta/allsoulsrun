package common

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
)

type CreateCharacterEvent struct {
	World       e.EntityID
	CharacterID e.EntityID
}

type CreateCharacterEventListener struct{}

func (c CreateCharacterEvent) GetType() e.EventType { return "Create Character" }

func (c CreateCharacterEventListener) On(ev e.Eventer, phase e.EventSequencePhase, t *e.Timeline) {
	create_ev := ev.(CreateCharacterEvent)
	world := t.Game.GetComponent(create_ev.World, gamecommon.World{}.GetComponentType()).(gamecommon.World)
	world.AddCharacter(create_ev.CharacterID)
	t.Game.SetComponent(create_ev.World, world)
}
