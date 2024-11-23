package common

import (
	"fmt"

	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
)

const (
	DAMAGE e.EventType = "Damage"
)

type Damage struct {
	Damaged      e.EntityID
	Damageamount int
	Fight        e.EntityID
}

func (d Damage) GetType() e.EventType {
	return DAMAGE
}

type DamageListener struct{}

func (dl DamageListener) On(ev e.Eventer, phase e.EventSequencePhase, t *e.Timeline) {
	var g = t.Game
	var elcomponent, statscomponent e.Componenter
	d := ev.(Damage)
	damageAmount := d.Damageamount
	elcomponent = g.GetComponent(d.Damaged, gamecommon.CharacterEnergyLevels{}.GetComponentType())
	statscomponent = g.GetComponent(d.Damaged, gamecommon.CharacterStats{}.GetComponentType())
	el := elcomponent.(gamecommon.CharacterEnergyLevels)
	name := statscomponent.(gamecommon.CharacterStats)
	el.Health -= damageAmount
	g.SetComponent(d.Damaged, el)
	fmt.Printf("%s has %d health left after taking %d damage\n", name.Name, el.Health, damageAmount)
	if el.Health <= 0 {
		t.StackEvent(&Die{Dead: d.Damaged, Fight: d.Fight})
	}

}
