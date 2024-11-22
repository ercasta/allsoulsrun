package common

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
)

const (
	AttackEventId = "Attack"
)

type AttackEvent struct {
	Attacker engine.EntityID
	Attacked engine.EntityID
	Fight    engine.EntityID
}

func (a AttackEvent) GetType() engine.EventType { return AttackEventId }

func (a AttackEvent) Scheduled(t *engine.Timeline) {}

func (a AttackEvent) Happen(t *engine.Timeline) {

}

func (a AttackEvent) Cancel(t *engine.Timeline) {
}
