package common

import (
	"fmt"
	rand "math/rand"

	"github.com/ercasta/allsoulsrun/pkg/engine"
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
)

const (
	DAMAGE = "Damage"
)

type Damage struct {
	Damaged engine.EntityID
	Fight   engine.EntityID
}

func (d Damage) GetType() e.EffectType {
	return DAMAGE
}

type DamageListener struct{}

func (dl DamageListener) OnApply(ef e.Effecter, es *e.EffectStack) {
	var g = es.Game
	var elcomponent, statscomponent e.Componenter
	d := ef.(Damage)
	damageAmount := rand.Intn(5) + 1
	elcomponent = g.GetComponent(d.Damaged, gamecommon.CharacterEnergyLevels{}.GetComponentType())
	statscomponent = g.GetComponent(d.Damaged, gamecommon.CharacterStats{}.GetComponentType())
	el := elcomponent.(gamecommon.CharacterEnergyLevels)
	name := statscomponent.(gamecommon.CharacterStats)
	el.Health -= damageAmount
	g.SetComponent(d.Damaged, el)
	fmt.Printf("%s has %d health left after taking %d damage\n", name.Name, el.Health, damageAmount)
	if el.Health <= 0 {
		es.StackEffect(&Die{Dead: d.Damaged, Fight: d.Fight})
	}

}

func (dl DamageListener) OnCancel(ef e.Effecter, es *e.EffectStack) {
	// Implementation for canceling the effect
}

func (dl DamageListener) OnStack(ef e.Effecter, es *e.EffectStack) {}
func (dl DamageListener) OnPop(ef e.Effecter, es *e.EffectStack)   {}
