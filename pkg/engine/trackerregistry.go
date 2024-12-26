package engine

type TrackerRegistryEntry struct {
	EventType          EventType
	EventTracker       EventTracker
	EventSequencePhase EventSequencePhase
}

type TrackerRegistry map[string]TrackerRegistryEntry
