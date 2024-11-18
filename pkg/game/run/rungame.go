package run

import (
	"github.com/ercasta/allsoulsrun/pkg/engine"
	ev "github.com/ercasta/allsoulsrun/pkg/game/events/common"
	a "github.com/ercasta/allsoulsrun/pkg/game/events/listeners"
	"github.com/gin-gonic/gin"
)

func NewRun() {
	var newgame engine.Game = engine.Game{}
	newgame.Init()
	var archer engine.Character = engine.NewCharacter("Legolas", 1, 0, 100, 10, 20, 5, 10, 100, 50)
	newgame.World.AddCharacter(archer)
	var fightevent = ev.FightEvent{}
	newgame.Timeline.AddEvent(&fightevent, 0)
	newgame.Timeline.AddEventListener(ev.AttackEventId, &a.AttackScheduler{})
	newgame.Timeline.AddEventListener(ev.FIGHT_EVENT, &a.AttackScheduler{})
	newgame.Run()
}

func Rungame(c *gin.Context) {
	NewRun()
	// var effectstack engine.EffectStack = engine.EffectStack{}
	// effectstack.StackEffect(&events.AttackEvent{})

	// var archer engine.Character = engine.NewCharacter("Legolas", 1, 0, 100, 10, 20, 5, 10, 100, 50)
	// var skeleton engine.Character = engine.NewCharacter("Skeleton", 1, 0, 100, 5, 5, 5, 5, 50, 0)
	// var characters [2]engine.Character = [2]engine.Character{archer, skeleton}
	// var initiative [2]engine.Character
	// initiative[0] = characters[0]
	// initiative[1] = characters[1]

	// fmt.Println("Start")
	// for i := 0; i < 1000000; i++ {
	// 	for z := 0; z < 100; z++ {
	// 		randomNumber := rand.Intn(10) + 1
	// 		if randomNumber%2 == 0 {
	// 			initiative[0] = initiative[1]
	// 			initiative[1] = initiative[0]
	// 		}
	// 		for j := 0; j < len(initiative); j++ {
	// 			hitRoll := rand.Intn(10) + 1
	// 			//fmt.Println(initiative[j].Name, "attacks", initiative[j^1].Name, "with a hit roll of", hitRoll)
	// 			initiative[j^1].Energy.Health -= hitRoll
	// 			//fmt.Println(initiative[j^1].Name, "has", initiative[j^1].Energy.Health, "health left")
	// 		}
	// 	}
	// }
	// fmt.Println("Done")
}
