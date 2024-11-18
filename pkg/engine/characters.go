package engine

var uuid uint64 = 0

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
	UUID   uint64
	Name   string
	Exp    CharacterExperience
	Stats  CharacterStats
	Energy CharacterEnergyLevels
}

func NewCharacter(name string, level, exp, nextLevelExp, strength, dexterity, intelligence, constitution, health, mana int) Character {
	uuid++
	return Character{
		UUID: uuid,
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
