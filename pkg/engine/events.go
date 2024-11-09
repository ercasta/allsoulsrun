package engine

type EventType int

const (
	DamageEvent = 1
	AttackEvent = 2
)

type Event interface {
	GetType() EventType
	Stack()
	Apply()
}

type EventStack struct {
	Events []Event
}

func (es *EventStack) StackEvent(e Event) {
	es.Events = append(es.Events, e)
}

func (es *EventStack) PopEvent() Event {
	if len(es.Events) == 0 {
		return nil
	}
	e := es.Events[len(es.Events)-1]
	es.Events = es.Events[0 : len(es.Events)-1]
	return e
}
