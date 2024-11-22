package engine

import (
	"math"
)

type GameTime uint64

type TimedEvent struct {
	StartTime GameTime
	Event     Eventer
}

type Timeline struct {
	CurrentTime    GameTime
	events         []TimedEvent
	eventListeners map[EventType][]EventListener
	Game           *Game
}

func (t *Timeline) AddEvent(e Eventer, time GameTime) {
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
	var smallest GameTime
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
	var nextId = t.findNextIdx()
	e := t.events[nextId].Event
	t.CurrentTime = t.events[nextId].StartTime
	t.events = t.events[1:]
	//e.Happen(t)
	if t.eventListeners != nil {
		for _, l := range t.eventListeners[e.GetType()] {
			l.OnEvent(e, t)
		}
	}
	t.Game.EffectStack.Resolve()

}

func (t *Timeline) NextEvent() Eventer {
	if len(t.events) == 0 {
		return nil
	}
	e := t.events[0].Event
	t.events = t.events[1:]
	return e
}
