package events

import "github.com/ercasta/allsoulsrun/pkg/engine"

type AttackEvent struct {
}

func (a *AttackEvent) GetType() engine.EventType { return engine.AttackEvent }

func (a *AttackEvent) Stack() {}

func (a *AttackEvent) Apply() {}
