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

func (d Die) GetType() e.EventType {
	return DIE
}

type DieListener struct{}

func (d DieListener) On(ev e.Eventer, phase e.EventSequencePhase, t *e.Timeline) {
	var dievent = ev.(*Die)

	println(gamecommon.GetName(dievent.Dead, t.Game), "is dead")
	fight := t.Game.GetComponent(dievent.Fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	fight.RemoveFighter(dievent.Dead)
	t.Game.SetComponent(dievent.Fight, fight)

}
