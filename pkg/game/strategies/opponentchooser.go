package strategies

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
)

func ChooseHealtiestOpponent(attacker e.EntityID, fight e.EntityID, t *e.Timeline) e.EntityID {
	g := t.Game
	targetHealth := 0
	targetOpponent := e.EntityID(0)
	fighcomponent := g.GetComponent(fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	var opponents = fighcomponent.GetOpponents(attacker)
	if len(opponents) > 0 {
		for _, opponent := range opponents {
			energy := g.GetComponent(opponent, gamecommon.CharacterEnergyLevels{}.GetComponentType()).(gamecommon.CharacterEnergyLevels)
			if energy.Health > targetHealth {
				targetHealth = energy.Health
				targetOpponent = opponent
			}
		}
	}
	return targetOpponent
}

func ChooseStrongestOpponent(attacker e.EntityID, fight e.EntityID, t *e.Timeline) e.EntityID {
	g := t.Game
	targetStrength := 0
	targetOpponent := e.EntityID(0)
	fighcomponent := g.GetComponent(fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	var opponents = fighcomponent.GetOpponents(attacker)
	if len(opponents) > 0 {
		for _, opponent := range opponents {
			stats := g.GetComponent(opponent, gamecommon.CharacterStats{}.GetComponentType()).(gamecommon.CharacterStats)
			if stats.Strength > targetStrength {
				targetStrength = stats.Strength
				targetOpponent = opponent
			}
		}
	}
	return targetOpponent
}
