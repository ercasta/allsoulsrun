package common

import (
	"encoding/json"

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

type CharacterExperience struct {
	Level        int
	Exp          int
	NextLevelExp int
}

func (c CharacterExperience) GetComponentType() engine.ComponentType {
	return CharacterExperienceType
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
