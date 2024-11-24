package run

import (
	"fmt"
	rand "math/rand"
	"time"

	"os"
	"path/filepath"

	"github.com/ercasta/allsoulsrun/pkg/engine"
	"github.com/ercasta/allsoulsrun/pkg/engine/utils"
	game "github.com/ercasta/allsoulsrun/pkg/game/common"
	ev "github.com/ercasta/allsoulsrun/pkg/game/events/common"
	el "github.com/ercasta/allsoulsrun/pkg/game/events/listeners"
	trackers "github.com/ercasta/allsoulsrun/pkg/game/trackers"
	"github.com/gin-gonic/gin"
)

func randstring(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewRun() {

	var newgame engine.Game = engine.Game{}
	newgame.Init()

	nonce := fmt.Sprintf("%d - %s", time.Now().UnixNano()/int64(time.Millisecond), randstring(12))
	newgame.UUID = nonce

	rundataPath := filepath.Join("rundata", nonce)
	err := os.MkdirAll(rundataPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	var world = newgame.CreateEntity()
	worldcomp := game.World{}
	hero := game.NewCharacter(&newgame, "Lufvd", 1, 0, 100, 10, 20, 5, 10, 100, 50)
	newgame.SetComponent(world, worldcomp)

	var fightevent = ev.FightEvent{}
	fightevent.Fight = newgame.CreateEntity()
	fight := game.Fight{}
	fight.AddFighter(hero, game.SIDE_CHARACTERS)

	fight.AddFighter(game.NewCharacter(&newgame, "Goblin", 1, 0, 50, 5, 5, 5, 5, 15, 0), game.SIDE_MONSTERS)
	fight.AddFighter(game.NewCharacter(&newgame, "Orc", 1, 0, 150, 15, 10, 10, 10, 20, 0), game.SIDE_MONSTERS)
	fight.AddFighter(game.NewCharacter(&newgame, "Slime", 1, 0, 150, 3, 1, 1, 2, 234, 0), game.SIDE_MONSTERS)

	newgame.SetComponent(fightevent.Fight, fight)

	// Whoever wants to mod the game, will need to add new listeners and events.

	newgame.Timeline.AddEventListener(ev.AttackEvent{}.GetType(), engine.OnEvent, el.AttackScheduler{})
	newgame.Timeline.AddEventListener(ev.AttackEvent{}.GetType(), engine.After, el.AttackScheduler{})
	newgame.Timeline.AddEventListener(fightevent.GetType(), engine.OnEvent, el.OnFight{})
	newgame.Timeline.AddEventListener(ev.Die{}.GetType(), engine.OnEvent, ev.DieListener{})
	newgame.Timeline.AddEventListener(ev.Damage{}.GetType(), engine.OnEvent, ev.DamageListener{})

	//newgame.Timeline.AddAnalyzer(ev.Damage{}.GetType(), engine.OnEvent, el.FightAnalyzer{})

	avroRecoder := utils.AvroRecoder{Basepath: rundataPath}

	newgame.Timeline.SetRecorder(&avroRecoder)
	newgame.Timeline.AddTracker(ev.CreateCharacterEvent{}.GetType(), engine.OnEvent, trackers.CharacterRecorder{})

	newgame.Timeline.ScheduleEvent(ev.CreateCharacterEvent{World: world, CharacterID: hero}, 0)
	newgame.Timeline.ScheduleEvent(fightevent, 1)

	newgame.Run()
	newgame.Terminate()

	avroRecoder.Close()

}

func Rungame(c *gin.Context) {
	go NewRun()

}
