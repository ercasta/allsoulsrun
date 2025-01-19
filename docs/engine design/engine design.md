# Game Design

## Game Sequence

Game is continuous in game time. A Character starts in town. He can then teleport to a level. Within the level, there can be several monster fights. Fights can have arbitrary number of participants on each of the two side (characters and monsters). Fights ends when one of there are no more participants in one of the sides. If characters are dead, they go back to town. If they win, they can move to next level fight, or decide to go back to town.

## Future improvements

Acts, Stash, Inventory, Skill Tree

## Event Timeline

Since the game does not have a UI the continually displays what's happening, and for performance purpose, instead of a loop with defined time steps, we have an event timeline. Events are placed on the timeline at specific (future) timestamps (we're talking about game time, not wall clock time). This allows to "fast forward" and save useless, "empty" time steps. The timeline can be altered: future events might shift, or they might be cancelled, or some new events might appear before some others.

Example:
- "Attack" for the main character is put on the event timeline in the future, according to character attack cooldown time.
- "Attack" for a foe is put on the event timeline too.
- "Explode" is put on the event timeline when a trap is placed.


## Event Stack
Once an event "happens", it is puy on a "resolution stack". This might trigger other events, that are put on top of the stack. When no events effects are triggered, the event on the top is "popped" and it happens. Again, this might trigger other events, which are stacked. This goes on until the stack is empty. Note that stack resolution happens virtually in the exact game time instant, but each listener gets its own unique game sequence id; this is needed to correctly track states associated to different events happening.

Example:
- "Damage" event is put on stack. This triggers a defensive "magic wall" ability from the opponent, so
- "Magic Wall" event is put on the stack. This ability prevents the damage to be inflicted, so
- "Damage cancel" event is put on stack. 

Then stack is resolved:
- "Damage cancel" is applied. This cancels the "Damage" from the stack
- "Magic Wall" applies. Actually, this event does nothing when applied. Its only reason to exist is to put other events on the stack, to allow counter-effects


## Character Build

A character build is composed of:
- Stats points progression
- Skill progression
- Item wishlist

## Party

A party is composed of one or more heroes

## Adventures

Runs follow a specific adventure.

An adventure has a series of levels. Each level has a series of encounters. Some levels make you retreat to previous level if you lose (e.g. boss fights)

Each adventure has its own party restrictions (e.g. number of heroes, classes, etc.)

## Worlds

Adventures take place in worlds. Worlds are simply collections of predefined monsters, locations, items. Adventures can reuse them, but also define new ones.
