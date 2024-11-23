# Internal Design

## Design Principles

- Use entity - component from E-C-S. Avoid explicit pointers to objects to avoid a big net of objects. Use UUID, and helper functions for retrieving. 
- Code is mainly in listeners. Listening structure is created mainly at startup time. Listeners have no data. All data is in components
- Component objects are immutable. Every change produces a copy, that is stored in the component history log. This helps tracking history.

## Technological principles

- Use Json to exchange data with UI: it's one of the most universal formats; moreover not much data should be exchanged with front end
- Use Json to manage content
- Use Arrow for in-memory analytics and data persistence
- Do not expose Arrow specifics to content creators.
- Modding must comply with usage of Arrow
- It should be possibile to swap storage backends for deployment e.g. to cloud, impacting only the engine.


## References

[Entity Component System](https://en.wikipedia.org/wiki/Entity_component_system)

[Wiki on Entity Component System](http://entity-systems.wikidot.com/)

[Why isn't Godot an ECS-based game engine?](https://godotengine.org/article/why-isnt-godot-ecs-based-game-engine/)

[State History](https://forum.heroiclabs.com/t/storing-match-state-history/3877)

[Event Sourcing Pattern](https://martinfowler.com/eaaDev/EventSourcing.html)

[Memento](https://refactoring.guru/design-patterns/memento)
