package skills

import "github.com/ercasta/allsoulsrun/pkg/engine"

type Skill interface {
	Activate(id engine.EntityID)
	Deactivate()
}

type SkillRegistry struct{}
