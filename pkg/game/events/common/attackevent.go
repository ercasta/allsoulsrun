package common

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
)

const (
	AttackEventId = "Attack"
)

type AttackEvent struct {
	Attacker     engine.EntityID
	Attacked     engine.EntityID
	Fight        engine.EntityID
	SecondAttack bool
}

func (a AttackEvent) GetType() engine.EventType { return AttackEventId }
