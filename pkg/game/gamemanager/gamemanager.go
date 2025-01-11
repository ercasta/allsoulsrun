package gamemanager

import (
	"encoding/json"
	"log"
	"path/filepath"

	engine "github.com/ercasta/allsoulsrun/pkg/engine"
	monsters "github.com/ercasta/allsoulsrun/pkg/game/monsters"
)

// Loads all data and keeps references to it: adventure, monsters, items, etc.

type GameManager struct {
	monsterDatabase *monsters.MonsterCompendium
}

// Declare listeners for adventure start, fight, etc.
type AdventureStartListener struct {
	gamemanager *GameManager
	adventure   *Adventure
}

func (asl *AdventureStartListener) On(ev engine.Eventer, phase engine.EventSequencePhase, t *engine.Timeline) {

}

func readRunConfig(runConfig string) *RunConfig {
	var runConfigLoaded RunConfig
	err := json.Unmarshal([]byte(runConfig), &runConfigLoaded)
	if err != nil {
		log.Fatalf("Failed to load run config: %v", err)
	}
	return &runConfigLoaded
}

func (gm *GameManager) Init(runConfig string, dataFolder string) {
	// Load all databases
	compendiumPath := filepath.Join(dataFolder, "monsters/compendium.json")

	//runConfigLoaded := readRunConfig(runConfig)

	var err error
	//adventure, err := readAdventure(dataFolder, runConfigLoaded.Adventure)

	gm.monsterDatabase, err = monsters.LoadCompendiumFromFile(compendiumPath)
	if err != nil {
		log.Fatalf("Failed to load compendium: %v", err)
	}

}
