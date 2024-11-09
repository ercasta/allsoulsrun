package allsouls

import (
	"fmt"
	"math/rand"

	"github.com/ercasta/allsoulsrun/pkg/engine"
	"github.com/ercasta/allsoulsrun/pkg/game/events"
)

type IntRange struct {
	Min, Max int
}

type Equipment struct {
	Name                     string
	PhysicalDamage           IntRange
	IceDamage                IntRange
	FireDamage               IntRange
	LightningDamage          IntRange
	DexterityBonus           IntRange
	StrengthBonus            IntRange
	IntelligenceBonus        IntRange
	ConstitutionBonus        IntRange
	DexterityBonusPercent    int
	StrengthBonusPercent     int
	IntelligenceBonusPercent int
	ConstitutionBonusPercent int
}

type CharacterEnergyLevels struct {
	Health int
	Mana   int
}

type CharacterStats struct {
	Strength     int
	Dexterity    int
	Intelligence int
	Constitution int
	Attack       int
	Defense      int
}

type CharacterExperience struct {
	Level        int
	Exp          int
	NextLevelExp int
}

type Character struct {
	Name   string
	Exp    CharacterExperience
	Stats  CharacterStats
	Energy CharacterEnergyLevels
}

func NewCharacter(name string, level, exp, nextLevelExp, strength, dexterity, intelligence, constitution, health, mana int) Character {
	return Character{
		Name: name,
		Exp: CharacterExperience{
			Level:        level,
			Exp:          exp,
			NextLevelExp: nextLevelExp,
		},
		Stats: CharacterStats{
			Strength:     strength,
			Dexterity:    dexterity,
			Intelligence: intelligence,
			Constitution: constitution,
		},
		Energy: CharacterEnergyLevels{
			Health: health,
			Mana:   mana,
		},
	}
}

func Rungame() {

	// Create a new character

	var eventstack engine.EventStack = engine.EventStack{}
	eventstack.StackEvent(&events.AttackEvent{})

	var archer Character = NewCharacter("Legolas", 1, 0, 100, 10, 20, 5, 10, 100, 50)
	var skeleton Character = NewCharacter("Skeleton", 1, 0, 100, 5, 5, 5, 5, 50, 0)
	var characters [2]Character = [2]Character{archer, skeleton}
	var initiative [2]Character
	initiative[0] = characters[0]
	initiative[1] = characters[1]

	fmt.Println("Start")
	for i := 0; i < 1000000; i++ {
		for z := 0; z < 100; z++ {
			randomNumber := rand.Intn(10) + 1
			if randomNumber%2 == 0 {
				initiative[0] = initiative[1]
				initiative[1] = initiative[0]
			}
			for j := 0; j < len(initiative); j++ {
				hitRoll := rand.Intn(10) + 1
				//fmt.Println(initiative[j].Name, "attacks", initiative[j^1].Name, "with a hit roll of", hitRoll)
				initiative[j^1].Energy.Health -= hitRoll
				//fmt.Println(initiative[j^1].Name, "has", initiative[j^1].Energy.Health, "health left")
			}
		}
	}
	fmt.Println("Done")
}
