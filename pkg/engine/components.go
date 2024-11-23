package engine

type ComponentType string

type ComponentHistory struct {
	HistoryId uint64
	Time      GameTime
	EntityID  EntityID
	Component Componenter
}

type Componenter interface {
	GetComponentType() ComponentType
	// PersistAll([]ComponentHistory)
}

type ComponentWithIdAndType struct {
	ComponentType ComponentType
	Component     Componenter
}

type componentManager struct {
	Components map[EntityID]map[ComponentType]Componenter
	history    []ComponentHistory
	historyId  uint64
	saveChan   chan ComponentHistory
	doneChan   chan bool
}

func (cm *componentManager) Done() {
	close(cm.saveChan)
	<-cm.doneChan
}

func (ch *componentManager) addComponentHistory(entityID EntityID, component Componenter, time GameTime) {
	ch.historyId++
	if ch.history == nil {
		ch.history = make([]ComponentHistory, 10000)
		ch.saveChan = make(chan ComponentHistory, 1000)
		ch.doneChan = make(chan bool)
		go SaveComponent(&ch.history, ch.saveChan, ch.doneChan)
	}
	ch.saveChan <- ComponentHistory{HistoryId: ch.historyId, Time: time, EntityID: entityID, Component: component}
	//ch.history = append(ch.history, ComponentHistory{historyId: ch.historyId, time: time, entityID: entityID, component: component})
}

func ComponentTypeMap() map[string]uint64 {
	return map[string]uint64{
		"CharacterExperience": 1,
		"CharacterStats":      2,
		"CharacterEnergy":     3,
	}
}

func (cm *componentManager) GetComponent(entityID EntityID, componentType ComponentType) Componenter {
	if components, ok := cm.Components[entityID]; ok {
		if component, ok := components[componentType]; ok {
			return component
		}
	}
	return nil
}

func (cm *componentManager) SetComponent(entityID EntityID, componenter Componenter, time GameTime) {
	if cm.Components == nil {
		cm.Components = make(map[EntityID]map[ComponentType]Componenter)
	}
	if _, ok := cm.Components[entityID]; !ok {
		cm.Components[entityID] = make(map[ComponentType]Componenter)
	}
	cm.Components[entityID][componenter.GetComponentType()] = componenter
	cm.addComponentHistory(entityID, componenter, time)
}

// EntityListener is an interface that defines methods for responding to changes in entities and their components.
// Implementers of this interface can handle events when entities are added or removed, and when components are added or removed from entities.
type EntityListener interface {
	// OnEntityAdded is called when an entity is added.
	// entityID is the unique identifier of the added entity.
	OnEntityAdded(entityID EntityID)

	// OnEntityRemoved is called when an entity is removed.
	// entityID is the unique identifier of the removed entity.
	OnEntityRemoved(entityID EntityID)

	// OnComponentAdded is called when a component is added to an entity.
	// entityID is the unique identifier of the entity.
	// componentTypeID is the unique identifier of the component type.
	// component is the instance of the added component.
	OnComponentAdded(entityID EntityID, componentType ComponentType, component any)

	// OnComponentRemoved is called when a component is removed from an entity.
	// entityID is the unique identifier of the entity.
	// componentTypeID is the unique identifier of the component type.
	OnComponentRemoved(entityID EntityID, componentType ComponentType)
}
