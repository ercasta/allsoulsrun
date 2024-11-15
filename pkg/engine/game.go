package engine

type Game struct {
	Timeline    Timeline
	EffectStack EffectStack
	World       World
}

func (g *Game) Init() {
	g.EffectStack = EffectStack{}
	g.World = World{}
	g.Timeline = Timeline{Game: g, World: &g.World}
	g.Timeline.Game = g
}

func (g *Game) Run() {
	for !g.Timeline.isFinished() {
		g.Timeline.RunNextEvent()
	}
}
