package monsters

import (
	"encoding/json"
	"log"
	"os"

	engine "github.com/ercasta/allsoulsrun/pkg/engine"
	common "github.com/ercasta/allsoulsrun/pkg/game/common"
)

type MonsterTemplate struct {
	Strength     int
	Dexterity    int
	Intelligence int
	Constitution int
	Health       int
	Mana         int
}

type MonsterCompendium map[string]MonsterTemplate

func LoadCompendiumFromFile(filename string) (*MonsterCompendium, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var compendium = MonsterCompendium{}
	err = json.Unmarshal(data, &compendium)
	if err != nil {
		return nil, err
	}

	return &compendium, nil
}

var mc MonsterCompendium = MonsterCompendium{}

func InitCompendium(filename string) {

	compendium, err := LoadCompendiumFromFile(filename)
	if err != nil {
		log.Fatalf("Failed to load compendium: %v", err)
	}

	mc = *compendium

	//for name, monster := range compendium {
	//		fmt.Printf("Monster: %s, Stats: %+v\n", name, monster)
	//	}

}

func NewMonster(g *engine.Game, monsterType string, name string) engine.EntityID {
	id := g.CreateEntity()

	mt := mc[monsterType]

	g.SetComponent(id, common.CharacterStats{
		Name:         name,
		Strength:     mt.Strength,
		Dexterity:    mt.Dexterity,
		Intelligence: mt.Intelligence,
		Constitution: mt.Constitution,
	})

	g.SetComponent(id, common.CharacterEnergyLevels{
		Health:    mt.Health,
		Mana:      mt.Mana,
		MaxHealth: mt.Health,
		MaxMana:   mt.Mana,
	})

	return id
}
