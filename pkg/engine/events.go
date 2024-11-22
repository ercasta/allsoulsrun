package engine

type EventType string

type Eventer interface {
	GetType() EventType
	//	Happen(t *Timeline)
}

type EventListener interface {
	// Called after events happen
	OnEvent(e Eventer, t *Timeline)
	OnScheduled(e Eventer, t *Timeline)
	OnCancel(e Eventer, t *Timeline)
	After(e Eventer, t *Timeline)
}
