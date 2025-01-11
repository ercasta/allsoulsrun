package engine

import (
	"math"
)

type GameTime uint64

type GameEventSequence uint64

type TimedEvent struct {
	StartTime GameTime
	Event     Eventer
}

type TimelineHistory struct {
	EventSequence      GameEventSequence
	EventSequencePhase EventSequencePhase
	Eventer            Eventer
}

type EventSequencePhase int

const (
	OnSchedule    EventSequencePhase = 0
	OnCancel      EventSequencePhase = 1
	OnStack       EventSequencePhase = 2
	OnPop         EventSequencePhase = 3
	OnStackCancel EventSequencePhase = 4
	OnEvent       EventSequencePhase = 5
	After         EventSequencePhase = 6
)

type Timeline struct {
	CurrentSequence GameEventSequence
	CurrentTime     GameTime
	events          []TimedEvent
	eventStack      []Eventer
	eventListeners  map[EventType]map[EventSequencePhase][]EventListener
	eventTrackers   map[EventType]map[EventSequencePhase][]EventTracker
	Game            *Game
	recorder        InitializableRecorder
}

func (t *Timeline) SetRecorder(r InitializableRecorder) {
	t.recorder = r
}

func (t *Timeline) ScheduleEvent(e Eventer, time GameTime) {
	// TODO: sort events by time, consider adding unique id
	t.events = append(t.events, TimedEvent{time, e})
	t.callListeners(e, OnSchedule)
}

func (t *Timeline) AddEventListener(e EventType, p EventSequencePhase, l EventListener) {
	if t.eventListeners == nil {
		t.eventListeners = make(map[EventType]map[EventSequencePhase][]EventListener, 1000)
	}
	if t.eventListeners[e] == nil {
		t.eventListeners[e] = make(map[EventSequencePhase][]EventListener, 1000)
	}
	t.eventListeners[e][p] = append(t.eventListeners[e][p], l)
}

func (t *Timeline) RemoveEventListener(e EventType, p EventSequencePhase, l EventListener) {
	if t.eventListeners == nil || t.eventListeners[e] == nil || t.eventListeners[e][p] == nil {
		return
	}
	listeners := t.eventListeners[e][p]
	for i, listener := range listeners {
		if listener == l {
			t.eventListeners[e][p] = append(listeners[:i], listeners[i+1:]...)
			break
		}
	}
}

func (t *Timeline) AddTracker(e EventType, p EventSequencePhase, l EventTracker) {
	if t.eventTrackers == nil {
		t.eventTrackers = make(map[EventType]map[EventSequencePhase][]EventTracker, 1000)
	}
	if t.eventTrackers[e] == nil {
		t.eventTrackers[e] = make(map[EventSequencePhase][]EventTracker, 1000)
	}
	t.eventTrackers[e][p] = append(t.eventTrackers[e][p], l)
	t.recorder.Init(l.GetType(), l.GetSchema())
}

// func (t *Timeline) AddAnalyzer(e EventType, p EventSequencePhase, l EventListener) {
// 	if t.eventListeners == nil {
// 		t.eventListeners = make(map[EventType]map[EventSequencePhase][]EventListener, 1000)
// 	}
// 	if t.eventListeners[e] == nil {
// 		t.eventListeners[e] = make(map[EventSequencePhase][]EventListener, 1000)
// 	}
// 	t.eventListeners[e][p] = append(t.eventListeners[e][p], l)
// }

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

func (t *Timeline) callListeners(e Eventer, p EventSequencePhase) {
	if t.eventListeners != nil && t.eventListeners[e.GetType()] != nil && t.eventListeners[e.GetType()][p] != nil {
		for _, l := range t.eventListeners[e.GetType()][p] {
			t.CurrentSequence++
			// t.saveHistory(e, p)
			l.On(e, p, t)
		}
	}
	if t.eventTrackers != nil && t.eventTrackers[e.GetType()] != nil && t.eventTrackers[e.GetType()][p] != nil {
		for _, tracker := range t.eventTrackers[e.GetType()][p] {
			tracker.Track(t.CurrentSequence, e, p, t.Game, (t.recorder).(Recorder))
		}
	}
}

func (t *Timeline) StackEvent(e Eventer) {
	t.eventStack = append(t.eventStack, e)
	t.callListeners(e, OnStack)
}

func (t *Timeline) PopEvent() Eventer {
	if len(t.eventStack) == 0 {
		return nil
	}
	e := t.eventStack[len(t.eventStack)-1]
	t.eventStack = t.eventStack[0 : len(t.eventStack)-1]
	t.callListeners(e, OnPop)
	return e
}

func (t *Timeline) Resolve() {
	for len(t.eventStack) != 0 {
		e := t.PopEvent()
		t.callListeners(e, OnEvent)
		t.callListeners(e, After)
	}
}

func (t *Timeline) RunNextEvent() {
	if len(t.events) == 0 {
		return
	}
	var nextId = t.findNextIdx()
	e := t.events[nextId].Event
	t.CurrentTime = t.events[nextId].StartTime
	t.events = append(t.events[:nextId], t.events[nextId+1:]...)

	t.StackEvent(e)
	t.Resolve()
}
