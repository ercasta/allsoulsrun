package engine

type TrackerType string

type AvroSchema string

type GameDataView interface {
	GetGameUUID() string
	GetComponent(entityID EntityID, componentType ComponentType) Componenter
}

type Recorder interface {
	//	Init(trackertype TrackerType, schema AvroSchema)
	Record(trackertype TrackerType, entityID any)
}

type InitializableRecorder interface {
	Init(trackertype TrackerType, schema AvroSchema)
	Record(trackertype TrackerType, entityID any)
}

type EventTracker interface {
	GetType() TrackerType
	GetSchema() AvroSchema
	Track(seq GameEventSequence, ev Eventer, phase EventSequencePhase, gdv GameDataView, rec Recorder)
}
