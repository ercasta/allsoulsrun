package common

import (
	"encoding/json"
	"fmt"

	engine "github.com/ercasta/allsoulsrun/pkg/engine"
)

var uuid uint64 = 0

const CharacterExperienceType engine.ComponentType = "CharacterExperience"
const CharacterStatsType engine.ComponentType = "CharacterStats"
const CharacterEnergyLevelsType engine.ComponentType = "CharacterEnergyLevels"

type CharacterEnergyLevels struct {
	Health int
	Mana   int
}

func (c CharacterEnergyLevels) GetComponentType() engine.ComponentType {
	return CharacterEnergyLevelsType
}

func (c CharacterEnergyLevels) PersistAll(ch []engine.ComponentHistory) {
	for _, component := range ch {
		if value, ok := (component.Component).(CharacterEnergyLevels); ok {
			// TODO Write all to file
			fmt.Printf("CharacterEnergyLevels for entity %d: Health %d, Mana %d\n", component.EntityID, value.Health, value.Mana)
		}
	}
}

func GetName(id engine.EntityID, g *engine.Game) string {
	stats := g.GetComponent(id, CharacterStatsType).(CharacterStats)
	return stats.Name
}

type CharacterStats struct {
	Name         string
	Strength     int
	Dexterity    int
	Intelligence int
	Constitution int
	Attack       int
	Defense      int
}

func (c CharacterStats) GetComponentType() engine.ComponentType {
	return CharacterStatsType
}

func (c CharacterStats) PersistAll(ch []engine.ComponentHistory) {
	for _, component := range ch {
		if value, ok := (component.Component).(CharacterStats); ok {
			// TODO Write all to file
			fmt.Printf("CharacterStats for entity %d: Strength %d", component.EntityID, value.Strength)
		}
	}
}

type CharacterExperience struct {
	Level        int
	Exp          int
	NextLevelExp int
}

func (c CharacterExperience) GetComponentType() engine.ComponentType {
	return CharacterExperienceType
}

func (c CharacterExperience) PersistAll(ch []engine.ComponentHistory) {
	for _, component := range ch {
		if value, ok := (component.Component).(CharacterExperience); ok {
			// TODO Write all to file
			fmt.Printf("CharacterStats for entity %d: Exp %d\n", component.EntityID, value.Exp)
		}
	}
}

// Test
func (ce *CharacterExperience) fromData(jsonData string) {
	json.Unmarshal([]byte(jsonData), ce)
}

func NewCharacter(g *engine.Game, name string, level, exp, nextLevelExp, strength, dexterity, intelligence, constitution, health, mana int) engine.EntityID {
	id := g.CreateEntity()

	expcomponent := CharacterExperience{
		Level:        level,
		Exp:          exp,
		NextLevelExp: nextLevelExp,
	}

	g.SetComponent(id, expcomponent)

	g.SetComponent(id, CharacterStats{
		Name:         name,
		Strength:     strength,
		Dexterity:    dexterity,
		Intelligence: intelligence,
		Constitution: constitution,
	})

	g.SetComponent(id, CharacterEnergyLevels{
		Health: health,
		Mana:   mana,
	})

	return id
}
