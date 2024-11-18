package common

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
	c "github.com/ercasta/allsoulsrun/pkg/game/common"
	ef "github.com/ercasta/allsoulsrun/pkg/game/effects/common"
)

const (
	AttackEventId = "Attack"
)

type AttackEvent struct {
	Attacker *engine.Character
	Attacked *engine.Character
	Fight    *c.Fight
	canceled bool
}

func (a *AttackEvent) GetType() engine.EventType { return AttackEventId }

func (a *AttackEvent) Scheduled(t *engine.Timeline) {}

func (a *AttackEvent) Happen(t *engine.Timeline) {
	if !a.canceled {
		t.Game.EffectStack.StackEffect(&ef.Damage{Damaged: a.Attacked, Fight: a.Fight})
	}
}

func (a *AttackEvent) Cancel(t *engine.Timeline) {
	a.canceled = true
}
