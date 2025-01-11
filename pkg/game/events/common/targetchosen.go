package common

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
)

const (
	TargetChosenEventId = "TargetChosen"
)

type TargetChosenEvent struct {
	Source engine.EntityID
	Target engine.EntityID
	Fight  engine.EntityID
}

func (a TargetChosenEvent) GetType() engine.EventType { return TargetChosenEventId }
