package allsouls

import (
	"fmt"
	"math/rand"
)

type CharacterEnergyLevels struct {
	Health int
	Mana   int
}

type CharacterStats struct {
	Strength     int
	Dexterity    int
	Intelligence int
	Constitution int
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

	var archer Character = NewCharacter("Legolas", 1, 0, 100, 10, 20, 5, 10, 100, 50)
	var skeleton Character = NewCharacter("Skeleton", 1, 0, 100, 5, 5, 5, 5, 50, 0)
	var characters [2]Character = [2]Character{archer, skeleton}
	var initiative [2]Character
	initiative[0] = characters[0]
	initiative[1] = characters[1]

	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(10) + 1
		if randomNumber%2 == 0 {
			initiative[0] = initiative[1]
			initiative[1] = initiative[0]
		}
		for j := 0; j < len(initiative); j++ {
			hitRoll := rand.Intn(10) + 1
			fmt.Println(initiative[j].Name, "attacks", initiative[j^1].Name, "with a hit roll of", hitRoll)
			initiative[j^1].Energy.Health -= hitRoll
			fmt.Println(initiative[j^1].Name, "has", initiative[j^1].Energy.Health, "health left")
		}
	}
}
