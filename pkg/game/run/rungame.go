package run

import (
	"fmt"
	"log"
	rand "math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ercasta/allsoulsrun/pkg/engine"
	"github.com/ercasta/allsoulsrun/pkg/engine/utils"
	"github.com/ercasta/allsoulsrun/pkg/game/common"
	game "github.com/ercasta/allsoulsrun/pkg/game/common"
	ev "github.com/ercasta/allsoulsrun/pkg/game/events/common"
	el "github.com/ercasta/allsoulsrun/pkg/game/events/listeners"
	"github.com/ercasta/allsoulsrun/pkg/game/gamemanager"
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

func CreateNewRun() string {
	return fmt.Sprintf("%d-%s", time.Now().UnixNano()/int64(time.Millisecond), randstring(12))
}

func configureRun(newgame *engine.Game, runId string, runConfig string) utils.AvroRecoder {
	nonce := runId
	newgame.UUID = nonce

	rundataPath := filepath.Join("rundata", nonce)
	err := os.MkdirAll(rundataPath, os.ModePerm)
	if err != nil {
		log.Printf("Error creating directory: %v", err)
		return utils.AvroRecoder{}
	}

	avroRecoder := utils.AvroRecoder{Basepath: rundataPath}
	newgame.Timeline.SetRecorder(&avroRecoder)
	return avroRecoder
}

func getBaseGame() (*engine.Game, engine.TrackerRegistry) {
	var newgame engine.Game = engine.Game{}
	newgame.Init()

	// Whoever wants to mod the game will change listeners and events.
	newgame.Timeline.AddEventListener(ev.AttackEvent{}.GetType(), engine.OnEvent, el.AttackScheduler{})
	newgame.Timeline.AddEventListener(ev.AttackEvent{}.GetType(), engine.After, el.AttackScheduler{})
	newgame.Timeline.AddEventListener(ev.FightEvent{}.GetType(), engine.OnEvent, el.OnFight{})
	newgame.Timeline.AddEventListener(ev.Die{}.GetType(), engine.OnEvent, ev.DieListener{})
	newgame.Timeline.AddEventListener(ev.Damage{}.GetType(), engine.OnEvent, ev.DamageListener{})

	return &newgame, getTrackersMap()
}

func getTrackersMap() engine.TrackerRegistry {
	trackersmap := make(map[string]engine.TrackerRegistryEntry)
	trackersmap["CharacterRecorder"] = engine.TrackerRegistryEntry{EventType: ev.CreateCharacterEvent{}.GetType(), EventTracker: trackers.CharacterRecorder{}, EventSequencePhase: engine.OnEvent}
	return trackersmap
}

func addToFight(newgame *engine.Game, fight engine.EntityID, fightpointer *game.Fight, character engine.EntityID, side game.Side) {
	fightpointer.AddFighter(character, side)
	newgame.SetComponent(character, common.EntityFight{FightId: fight})
}

func runAdventure(newgame *engine.Game, trackerRegistry engine.TrackerRegistry, runConfig string) {
	gamemanager := gamemanager.GameManager{}
	gamemanager.Init(runConfig, "../../../examples/data")
	var world = newgame.CreateEntity()
	worldcomp := game.World{}

	hero := game.NewCharacter(newgame, "Lufvd", 1, 0, 100, 10, 20, 5, 10, 100, 50)
	newgame.SetComponent(world, worldcomp)

	var fightevent = ev.FightEvent{}
	fightevent.Fight = newgame.CreateEntity()
	fight := game.Fight{}

	addToFight(newgame, fightevent.Fight, &fight, hero, game.SIDE_CHARACTERS)

	addToFight(newgame, fightevent.Fight, &fight, game.NewCharacter(newgame, "Goblin", 1, 0, 50, 5, 5, 5, 5, 15, 0), game.SIDE_MONSTERS)
	addToFight(newgame, fightevent.Fight, &fight, game.NewCharacter(newgame, "Orc", 1, 0, 150, 15, 10, 10, 10, 20, 0), game.SIDE_MONSTERS)
	addToFight(newgame, fightevent.Fight, &fight, game.NewCharacter(newgame, "Slime", 1, 0, 150, 3, 1, 1, 2, 234, 0), game.SIDE_MONSTERS)

	newgame.SetComponent(fightevent.Fight, fight)

	// make dynamic, for all configured trackers in runConfig
	trackerRegistryEntry := trackerRegistry["CharacterRecorder"]
	newgame.Timeline.AddTracker(trackerRegistryEntry.EventType, trackerRegistryEntry.EventSequencePhase, trackerRegistryEntry.EventTracker)

	newgame.Timeline.ScheduleEvent(ev.CreateCharacterEvent{World: world, CharacterID: hero}, 0)
	newgame.Timeline.ScheduleEvent(fightevent, 1)

	newgame.Run()
	newgame.Terminate()
}

func NewRun(runId string, runConfig string) {

	newgame, trackerRegistry := getBaseGame()
	avroRecoder := configureRun(newgame, runId, runConfig)

	runAdventure(newgame, trackerRegistry, runConfig)

	avroRecoder.Close()

}

func Rungame(c *gin.Context) {
	runId := CreateNewRun()
	runConfig := c.PostForm("runconfig")
	go NewRun(runId, runConfig)
	c.JSON(200, gin.H{"runId": runId})
}
