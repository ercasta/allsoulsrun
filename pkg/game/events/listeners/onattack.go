package listeners

import (
	"fmt"
	rand "math/rand"

	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
	a "github.com/ercasta/allsoulsrun/pkg/game/events/common"
	strategies "github.com/ercasta/allsoulsrun/pkg/game/strategies"
)

type AttackScheduler struct{}

func (oar AttackScheduler) scheduleNewAttackToFirstOpponent(attacker e.EntityID, fight e.EntityID, t *e.Timeline) {
	g := t.Game
	fighcomponent := g.GetComponent(fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
	var opponents = fighcomponent.GetOpponents(attacker)
	if len(opponents) > 0 {
		stats := g.GetComponent(attacker, gamecommon.CharacterStats{}.GetComponentType()).(gamecommon.CharacterStats)
		waitTime := int(1000 * (20 / stats.Dexterity))
		newTime := t.CurrentTime + e.GameTime(waitTime)
		opponent := strategies.ChooseStrongestOpponent(attacker, fight, t)
		t.ScheduleEvent(a.AttackEvent{Attacker: attacker, Attacked: opponent, Fight: fight, SecondAttack: false}, newTime)
		if rand.Float64() > 0.9 {
			t.ScheduleEvent(a.AttackEvent{Attacker: attacker, Attacked: opponents[0], Fight: fight, SecondAttack: true}, newTime)
		}
	}

}

func (oar AttackScheduler) reschedule(e e.Eventer, t *e.Timeline) {
	attackevent := e.(a.AttackEvent)
	oar.scheduleNewAttackToFirstOpponent(attackevent.Attacker, attackevent.Fight, t)
}

func (oar AttackScheduler) On(ev e.Eventer, phase e.EventSequencePhase, t *e.Timeline) {
	switch phase {
	case e.OnEvent:
		attackevent := ev.(a.AttackEvent)
		stats := t.Game.GetComponent(attackevent.Attacker, gamecommon.CharacterStats{}.GetComponentType()).(gamecommon.CharacterStats)
		damage := stats.Strength/10 + rand.Intn(3) + 1
		fight := t.Game.GetComponent(attackevent.Fight, gamecommon.Fight{}.GetComponentType()).(gamecommon.Fight)
		if fight.IsInFight(attackevent.Attacked) && fight.IsInFight(attackevent.Attacker) {
			opponentstats := t.Game.GetComponent(attackevent.Attacked, gamecommon.CharacterStats{}.GetComponentType()).(gamecommon.CharacterStats)
			var issecondattack string
			if attackevent.SecondAttack {
				issecondattack = "It's a double attack!!!"
			} else {
				issecondattack = ""
			}
			fmt.Printf("%s attacks %s dealing %d damage. %s \n", stats.Name, opponentstats.Name, damage, issecondattack)
			t.StackEvent(a.Damage{Damaged: attackevent.Attacked, Fight: attackevent.Fight, Damageamount: damage})
		}
	case e.OnCancel:
		oar.reschedule(ev, t)
	case e.After:
		attackevent := ev.(a.AttackEvent)
		if !attackevent.SecondAttack {
			oar.reschedule(ev, t)
		}
	}
}

func (oar AttackScheduler) After(e e.Eventer, t *e.Timeline) {

}
