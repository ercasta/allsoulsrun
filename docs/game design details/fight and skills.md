# Fight

It actually starts from the end of the previous fight. A listener is hooked to the Fight End event; it looks up the next fight in the adventure, and schedules a Fight Start Event (with a fight id).

Another listener is hooked to the Fight Start Event. It looks up the details of the fight in the adventure, and:
- Adds Heroes and Monsters to the Fight (looking up details in the adventure manager, and monsters from the Monsters Compendium)
- Puts on the stack the skill activation event for Heroes and Monsters
- The listener for the skill activation event (connected to the Skill Management System) activates the Heroes and Monster Skills (looking up Skills in the Skills Registry). 
- From this moment on, fight is driven by skills.

## Fight progress

- Firs


# Skills Manager

Handles Skill Activate, Schedules skill deactivation (at fight end).

Make skill development easy; move as much as possibile into skill manager (e.g. when to deactivate)

Skills attribute:
- 


# Examples of Skills

# Simple attack

Performs an attack with the weapon. The skill also chooses a suitable target.

Implementation: a TargetChosen event is stacked (to allow other listeners to act on this). On TargetChosen Event, an attack is scheduled. 

TargetChosen event:
- id source
- id target
GetSource()
GetTarget()

Notes: for efficiency reasons, a listener for the specific entity is added on skill activation, and removed on skill deactivation

Notes: this skill is parametric. As a parameter, it takes the target selection strategy (strongest, healthiest).

Also, after the attack completes, or if it is cancelled, a new TargetChosen event is stacked, and so on.

## Double strike
On **first** strike, has a chance to immediately 

Implementation: on the Attack Event.

Notes: for efficiency, on activation, skill registers a listener

## Defend

When standing on front line, and a friend on second line is attacked from an enemy, the character has a chance to block the attack (taking the damage itself)

Implementation: on stacking the Attack, a TargetHijack event is stacked on top. The skill can change the target.

Notes: on activation, this skill registers listeners for Attack and TargetHijack events.


