package listeners

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
	a "github.com/ercasta/allsoulsrun/pkg/game/events/common"
)

type OnFight struct{}

func (oar OnFight) scheduleNewAttackToFirstOpponent(attacker e.EntityID, fight e.EntityID, t *e.Timeline) {
	g := t.Game
	fighcomponent := g.GetComponent(fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	var opponents = fighcomponent.GetOpponents(attacker)
	if len(opponents) > 0 {
		t.ScheduleEvent(a.AttackEvent{Attacker: attacker, Attacked: opponents[0], Fight: fight, SecondAttack: false}, t.CurrentTime+2000)
	}

}

func (oar OnFight) On(ev e.Eventer, phase e.EventSequencePhase, t *e.Timeline) {
	fight := ev.(a.FightEvent)
	fightcomponent := t.Game.GetComponent(fight.Fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	fighters := fightcomponent.GetFighters()
	for _, fighter := range fighters {
		oar.scheduleNewAttackToFirstOpponent(fighter, fight.Fight, t)
	}
}
