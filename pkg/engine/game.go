package engine

import (
	"slices"
)

type Game struct {
	UUID             string
	componentManager componentManager
	Timeline         Timeline
	entityID         EntityID
}

func (g *Game) Init() {
	g.componentManager = componentManager{}
	g.Timeline = Timeline{Game: g}
}

func (g *Game) Terminate() {
	g.componentManager.Done()
	g.SaveHistory()
}

func (g *Game) Run() {
	for !g.Timeline.isFinished() {
		g.Timeline.RunNextEvent()
	}
}

func (g *Game) GetGameUUID() string {
	return g.UUID
}

func (g *Game) CreateEntity() EntityID {
	g.entityID++
	return g.entityID
}

func (g *Game) GetComponent(entityID EntityID, componentType ComponentType) Componenter {
	return g.componentManager.GetComponent(entityID, componentType)
}

func (g *Game) SetComponent(entityID EntityID, component Componenter) {
	g.componentManager.SetComponent(entityID, component, g.Timeline.CurrentSequence)
}

func (g *Game) GetHistoryLen() int {
	count := 0
	for _, h := range g.componentManager.history {
		if h != (ComponentHistory{}) {
			count++
		}
	}
	return count
}

func (g *Game) SaveHistory() {
	types := make([]ComponentType, 1000)
	for _, h := range g.componentManager.history {
		if h.Component != nil && !slices.Contains(types, h.Component.GetComponentType()) {
			types = append(types, h.Component.GetComponentType())
		}
	}
	for _, t := range types {
		for _, h := range g.componentManager.history {
			if h.Component != nil && h.Component.GetComponentType() == t {
				// h.Component.PersistAll(g.componentManager.history)
				break
			}
		}
	}
}
