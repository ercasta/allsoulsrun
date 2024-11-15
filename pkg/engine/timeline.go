package engine

import (
	"math"
)

type TimedEvent struct {
	StartTime uint64
	Event     Eventer
}

type Timeline struct {
	events         []TimedEvent
	eventListeners map[EventType][]EventListener
	World          *World
	Game           *Game
}

type EventListener interface {
	// Called after events happen
	OnEvent(e Eventer, t *Timeline)
	OnScheduled(e Eventer, t *Timeline)
	OnCancel(e Eventer, t *Timeline)
}

func (t *Timeline) AddEvent(e Eventer, time uint64) {
	// TODO: sort events by time
	t.events = append(t.events, TimedEvent{time, e})
}

func (t *Timeline) AddEventListener(e EventType, l EventListener) {
	if t.eventListeners == nil {
		t.eventListeners = make(map[EventType][]EventListener)
	}
	t.eventListeners[e] = append(t.eventListeners[e], l)
}

func (t *Timeline) isFinished() bool {
	return len(t.events) == 0
}

func (t *Timeline) findNextIdx() int {
	var smallest uint64
	var idx int
	smallest = math.MaxUint64
	for i, e := range t.events {
		if e.StartTime < smallest {
			smallest = e.StartTime
			idx = i
		}
	}
	return idx
}

func (t *Timeline) RunNextEvent() {
	if len(t.events) == 0 {
		return
	}
	e := t.events[t.findNextIdx()].Event
	t.events = t.events[1:]
	e.Happen(t)
	t.Game.EffectStack.Resolve()
	if t.eventListeners != nil {
		for _, l := range t.eventListeners[e.GetType()] {
			l.OnEvent(e, t)
		}
	}
}

func (t *Timeline) NextEvent() Eventer {
	if len(t.events) == 0 {
		return nil
	}
	e := t.events[0].Event
	t.events = t.events[1:]
	return e
}
