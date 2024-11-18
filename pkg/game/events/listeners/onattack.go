package listeners

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	f "github.com/ercasta/allsoulsrun/pkg/game/common"
	a "github.com/ercasta/allsoulsrun/pkg/game/events/common"
)

type AttackScheduler struct {
	Attacker *e.Character
	Fight    *f.Fight
}

func scheduleNewAttackToFirstOpponent(attacker *e.Character, fight *f.Fight, t *e.Timeline) {
	var opponents = fight.GetOpponents(*attacker)
	if len(opponents) > 0 {
		t.AddEvent(&a.AttackEvent{Attacker: attacker, Attacked: opponents[0], Fight: fight}, t.CurrentTime+2000)
	}
}

func (oar *AttackScheduler) reschedule(_ e.Eventer, t *e.Timeline) {
	// if oar.AttackEvent != nil {
	// 	var fight = oar.AttackEvent.Fight
	// 	if fight.IsInFight(oar.AttackEvent.Attacked) {
	// 		t.AddEvent(&a.AttackEvent{Attacker: oar.AttackEvent.Attacker, Attacked: oar.AttackEvent.Attacked, Fight: fight}, t.CurrentTime+2000)
	// 		return
	// 	}
	// }
	if oar.Fight != nil {
		// Choose new target
		var fight = oar.Fight
		if oar.Attacker == nil {
			// Attacker might be nil... should schedule a new attack for all fighters
			fighters := fight.GetFighters()
			for _, fighter := range fighters {
				scheduleNewAttackToFirstOpponent(fighter, fight, t)
			}
		} else {
			scheduleNewAttackToFirstOpponent(oar.Attacker, fight, t)
		}
	}
}

func (oar *AttackScheduler) OnEvent(e e.Eventer, t *e.Timeline) {
	switch e.(type) {
	case *a.FightEvent:
		var FightEvent = e.(*a.FightEvent)
		oar.Fight = FightEvent.Fight
	case *a.AttackEvent:
		// Do nothing
		var atk = e.(*a.AttackEvent)
		oar.Fight = atk.Fight
		oar.Attacker = atk.Attacker
	}
	oar.reschedule(e, t)
}

func (oar *AttackScheduler) OnCancel(e e.Eventer, t *e.Timeline) {
	oar.reschedule(e, t)
}

func (oar *AttackScheduler) OnScheduled(e e.Eventer, t *e.Timeline) {
	// Do nothing
}
