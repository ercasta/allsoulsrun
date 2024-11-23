package engine

type EventType string

type Eventer interface {
	GetType() EventType
	//	Happen(t *Timeline)
}

type EventListener interface {
	On(ev Eventer, phase EventSequencePhase, t *Timeline)
}
