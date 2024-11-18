package common

import (
	"fmt"
	rand "math/rand"

	e "github.com/ercasta/allsoulsrun/pkg/engine"
	f "github.com/ercasta/allsoulsrun/pkg/game/common"
)

const (
	DAMAGE = "Damage"
)

type Damage struct {
	Damaged   *e.Character
	Fight     *f.Fight
	cancelled bool
}

func (d *Damage) GetType() e.EffectType {
	return DAMAGE
}

func (d *Damage) Apply(es *e.EffectStack) {
	if !d.cancelled {
		damageAmount := rand.Intn(5) + 1
		d.Damaged.Energy.Health -= damageAmount
		fmt.Printf("%s has %d health left after taking %d damage\n", d.Damaged.Name, d.Damaged.Energy.Health, damageAmount)
		if d.Damaged.Energy.Health <= 0 {
			es.StackEffect(&Die{Dead: d.Damaged, Fight: d.Fight})
		}
	}
}

func (d *Damage) Cancel() {
	// Implementation for canceling the effect
	d.cancelled = true
}
