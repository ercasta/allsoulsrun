package common

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	c "github.com/ercasta/allsoulsrun/pkg/game/common"
)

const (
	DIE = "Die"
)

type Die struct {
	Dead      *e.Character
	Fight     *c.Fight
	cancelled bool
}

func (d *Die) GetType() e.EffectType {
	return DIE
}

func (d *Die) Apply() {
	if !d.cancelled {
		d.Fight.RemoveFighter(d.Dead)
	}
}

func (d *Die) Cancel() {
	// Implementation for canceling the effect
	d.cancelled = true
}
