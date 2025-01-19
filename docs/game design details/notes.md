# Notes

- typical run length: 10K encounters? (avg 100 encounters per player level)
- how many areas per world? 10? **Subdivide areas by acts, each act having a final boss before unlocking next area (decided)**
- how to progress within act? Visit different areas?
- how many character levels? **100 (decided)**
- How long should a play last i.e. how long should a player spend on an adventure? A few minutes? hours? Days? weeks? **Hours, maybe over several days, ideally not more than a week or month (depending on the result the player wants to achieve)**. The pattern should be players having their "favourite" parties (player builds and party formations), and configuring runs, running them, looking at results, adjusting either runs or builds / formations. These are not monoliths, they are composable elements. Some of these might be referenced by URIs, so people can publish their builds / formations, and players can just configure runs.
- Unbalance by default: world and rules parameters might vary (e.g. exp required per level; these params are collectively known as "the ruleset" - it will be fun to play with different / "extreme" rulesets)
- Areas have defaults. Of nothing is configured, players are "attracted" to areas according to a criteria (actually, a progression of areas, determined by the adventure, with default rules for retry / advance / fall back e.g. num fall back losses. So player specifies these rules)
- Acts can also be played independently. Each act has a starting level, which determines the skill points the players can use.
- Each area is divided in levels. Each area contains a series of encounters; If the character survives all of them, he checkpoints (can start from the next area, if he dies).
- Defeat 2h recover. Std interval 60 sec; full health recover
- Level goals: 
    - best accel
    - Best overall time
    - Best consistency on N runs
- Next level strategy: decide whether to go to next level or not; it's a threshold e.g. 3 encounters at min 60% health left. Rationale: slower is safer but might miss opp.

What's the simplest, first version of the game? One in which nothing changes but levelling up, no skills, only the strategy changes (next on: x encounters at y% remaining.). And I want to know what happened, so I want reporting, even in textual form. 


**Important** All modifications to components happen within the "current event", and tracked in this way. Listeners cannot modify game components, they can only modify the stack; also, modifications to game components only take place in the "happen" phase, and tied to the event. Events cannot modify the stack, changes are done only through listeners. Do events have components too? And do listeners listen to event components? This would allow creating a flexible event taxonomy.

The idea is each event does something and then something else happens. e.g.

Attack (or attack abort)-> Hit or Miss -> Damage -> (potentially) Death, then respawn, or something else.

The thing should be: events only care about themselves, and something else glues events together. This creates flexibility but also complexity in debugging.
It would be easier to:
- do everything in one function, one thing leading to another; but this would create a huge "chain of ifs": where to stop?
- make an event directly trigger other events. But this creates coupling.
- Listeners should be the "glue" between events. Triggered by one event; apply some rule and fire another event. As they are just "glue", they cannot change the game world. 
- This would lead to an easier understanding. If an event "happens" there is no doubt about what it means, in terms of effects on the world; this would help reconstructing what happened during the game.
- Also, "rules" for skills and everything else should be in listeners. 
- Events change the state of the world; Listeners ("rules") change the scheduling and happening of events.

Events
- charging (what?)
- start burning
- poison
- death
- spawn
- eat  Rule: when eat, heal. ("Lufvd eats an apple". "Lufvd is healed by 5 HP".)
- kill. Rule: when kill, add exp (add exp event? No? Yes! Exp gained event. Tied to levelup )
- levelup:

It seems overkill, for every minuscule bit of world modification i need an event?!?! An then a listener? (a "system"?)

or maybe yes? It depends on how cumbersome it is to create one.

systems react to events, and fire other events. Every change happens within the context of an event.

- probably we won't trace every minuscule component change, but... who cares?!

- We need events to trigger other systems. That's why we need lots of them. And also to allow "interruptions" and "reactions". Anyway, we can split events afterwards










