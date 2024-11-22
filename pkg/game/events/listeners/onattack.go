package listeners

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
	ef "github.com/ercasta/allsoulsrun/pkg/game/effects/common"
	a "github.com/ercasta/allsoulsrun/pkg/game/events/common"
)

type AttackScheduler struct{}

func (oar AttackScheduler) scheduleNewAttackToFirstOpponent(attacker e.EntityID, fight e.EntityID, t *e.Timeline) {
	g := t.Game
	fighcomponent := g.GetComponent(fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	var opponents = fighcomponent.GetOpponents(attacker)
	if len(opponents) > 0 {
		t.AddEvent(a.AttackEvent{Attacker: attacker, Attacked: opponents[0], Fight: fight}, t.CurrentTime+2000)
	}

}

func (oar AttackScheduler) reschedule(e e.Eventer, t *e.Timeline) {
	attackevent := e.(a.AttackEvent)
	oar.scheduleNewAttackToFirstOpponent(attackevent.Attacker, attackevent.Fight, t)
}

func (oar AttackScheduler) OnEvent(e e.Eventer, t *e.Timeline) {
	attackevent := e.(a.AttackEvent)
	t.Game.EffectStack.StackEffect(ef.Damage{Damaged: attackevent.Attacked, Fight: attackevent.Fight})
}

func (oar AttackScheduler) OnCancel(e e.Eventer, t *e.Timeline) {
	oar.reschedule(e, t)
}

func (oar AttackScheduler) OnScheduled(e e.Eventer, t *e.Timeline) {
	// Do nothing
}

func (oar AttackScheduler) After(e e.Eventer, t *e.Timeline) {
	oar.reschedule(e, t)
}
