# Game Design

## Character

### Stats

- Strength
- Dexterity
- Intelligence
- Constitution

### Levels

- Each level requires 25% more than previous to level up

### Strategy

Strategy means setting a party configuration, in terms of skills and front / back line configuration.

Besides a default strategy, there can be specific strategies for specific levels or acts

Most Skills are parametric. For example the attack skill has an opponent choosing parameter. Default is strongest, but it can be set to healthiest.

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


