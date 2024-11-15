package engine

type EventType string

type Eventer interface {
	GetType() EventType
	Happen(t *Timeline)
}
