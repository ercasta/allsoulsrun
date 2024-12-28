# Game Design

## Character

### Stats

- Strength
- Dexterity
- Intelligence
- Constitution

### Levels

- Each level requires 25% more than previous to level up


## Events & Listeners

List of all game events and listeners

- Create Character
    - Trackers:
        - Records character name
- Adventure Start

- Fight Start:
    - Listeners:
        - On: Spawns monsters, add them to the fight together with characters
    
- Die: 
    - On: removes the fighter from the fight
- Fight End
    - On: Add next encounter to the timeline
- Adventure End

## Trackers

List of all trackers

- Cha

### StartGame

Creates World, Character(s).

### Create Character

On character creation


### Enter Area

### Exit Area

### Start Encounter


