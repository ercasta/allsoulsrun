# Internal Design

## Design Principles

- Use entity - component from E-C-S. Avoid explicit pointers to objects to avoid a big net of objects. Use UUID, and helper functions for retrieving. 
- Code is mainly in listeners. Listening structure is created mainly at startup time. Listeners have no data. All data is in components
- Component objects are immutable. Every change produces a copy, that is stored in the component history log. This helps tracking history.
- Use Trackers to track relevant run data
- Use Reports to analyzed tracked data
- Don't put code in components (except getters and setters)

## Technological principles

- Use Json to exchange data with UI: it's one of the most universal formats; moreover not much data should be exchanged with front end
- Use Json to manage content
- Future version: use Apache Beam to generate summary reports
- Summary reports should consist in small amount of data
- Use plain logging for debugging. To store logs, use any suitable format (es. JSON Objects in Base64, one per line)
- Individual summary reports can be exposed at different urls
- Use external calibration files to manage balancing
- Leverage an extensive test suite to test game balance on different runs, using range validation.

## Technical notes: Observer pattern

Event Timeline is based on the Observer pattern. All changes to these structure is notified to listener: this allows creating new events that trigger in specific conditions


## References

[Entity Component System](https://en.wikipedia.org/wiki/Entity_component_system)

[Wiki on Entity Component System](http://entity-systems.wikidot.com/)

[Why isn't Godot an ECS-based game engine?](https://godotengine.org/article/why-isnt-godot-ecs-based-game-engine/)

[State History](https://forum.heroiclabs.com/t/storing-match-state-history/3877)

[Event Sourcing Pattern](https://martinfowler.com/eaaDev/EventSourcing.html)

[Memento](https://refactoring.guru/design-patterns/memento)
