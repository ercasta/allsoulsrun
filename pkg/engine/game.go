package engine

type Game struct {
	componentManager componentManager
	Timeline         Timeline
	EffectStack      EffectStack
	entityID         EntityID
}

func (g *Game) Init() {
	g.EffectStack = EffectStack{Game: g}
	g.Timeline = Timeline{Game: g}
}

func (g *Game) Run() {
	for !g.Timeline.isFinished() {
		g.Timeline.RunNextEvent()
	}
}

func (g *Game) CreateEntity() EntityID {
	g.entityID++
	return g.entityID
}

func (g *Game) GetComponent(entityID EntityID, componentType ComponentType) Componenter {
	return g.componentManager.GetComponent(entityID, componentType)
}

func (g *Game) SetComponent(entityID EntityID, component Componenter) {
	g.componentManager.SetComponent(entityID, component, g.Timeline.CurrentTime)
}
