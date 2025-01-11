package common

import "github.com/ercasta/allsoulsrun/pkg/engine"

// SourceEventListener is a struct that associates an event listener with a specific source entity, to simplify activating listeners for specific entities.
// It contains the following fields:
// - Source: The ID of the entity that is the source of the events.
// - EventListener: The event listener that will handle events from the specified source entity.
type SourceEventListener struct {
	Source        engine.EntityID
	EventListener engine.EventListener
}

// SourcedEvent is an interface that represents an event with a source entity.
// It has the following method:
// - GetSource: Returns the ID of the entity that is the source of the event.
type SourcedEvent interface {
	GetSource() engine.EntityID
}

// On is a method of SourceEventListener that handles events during a specific phase of the event sequence.
// It takes the following parameters:
// - ev: The event to be handled.
// - phase: The phase of the event sequence during which the event is being handled.
// - t: The timeline in which the event is being processed.
// If the event implements the SourcedEvent interface and its source matches the SourceEventListener's source,
// the event is passed to the associated EventListener's On method.
func (s SourceEventListener) On(ev engine.Eventer, phase engine.EventSequencePhase, t *engine.Timeline) {
	if sev, ok := ev.(SourcedEvent); ok && s.Source == sev.GetSource() {
		s.EventListener.On(ev, phase, t)
	}
}
