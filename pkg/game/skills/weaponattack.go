package skills

import "github.com/ercasta/allsoulsrun/pkg/engine"

type WeaponAttack struct {
	Game *engine.Game
}

func (wa WeaponAttack) Activate(id engine.EntityID) {
	// Choose target, then attack

}

func (wa WeaponAttack) Deactivate() {

}
