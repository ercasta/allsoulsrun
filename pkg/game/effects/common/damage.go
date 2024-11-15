package common

import (
	"fmt"

	e "github.com/ercasta/allsoulsrun/pkg/engine"
)

const (
	DAMAGE = "Damage"
)

type Damage struct {
	Damaged   *e.Character
	cancelled bool
}

func (d *Damage) GetType() e.EffectType {
	return DAMAGE
}

func (d *Damage) Apply() {
	if !d.cancelled {
		d.Damaged.Energy.Health -= 1
		fmt.Printf("%s has %d health left\n", d.Damaged.Name, d.Damaged.Energy.Health)
	}
}

func (d *Damage) Cancel() {
	// Implementation for canceling the effect
	d.cancelled = true
}
