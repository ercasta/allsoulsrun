package common

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
)

const (
	DIE = "Die"
)

type Die struct {
	Dead  e.EntityID
	Fight e.EntityID
}

func (d Die) GetType() e.EffectType {
	return DIE
}

type DieListener struct{}

func (d DieListener) OnApply(e e.Effecter, es *e.EffectStack) {
	var dievent = e.(*Die)

	println(gamecommon.GetName(dievent.Dead, es.Game), "is dead")
	fight := es.Game.GetComponent(dievent.Fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	fight.RemoveFighter(dievent.Dead)
	es.Game.SetComponent(dievent.Fight, fight)

}

func (d DieListener) OnStack(e e.Effecter, es *e.EffectStack) {}

func (d DieListener) OnPop(e e.Effecter, es *e.EffectStack) {}

func (d DieListener) OnCancel(e e.Effecter, es *e.EffectStack) {}
