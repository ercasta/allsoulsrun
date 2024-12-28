package monsters

import (
	"testing"

	"github.com/ercasta/allsoulsrun/pkg/engine"
	common "github.com/ercasta/allsoulsrun/pkg/game/common"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	filename := "../../../examples/data/monsters/compendium.json"

	InitCompendium(filename)

	g := &engine.Game{}
	g.Init()

	id := NewMonster(g, "Goblin", "My Goblin")
	cs := g.GetComponent(id, common.CharacterStatsType).(common.CharacterStats)
	if cs.Name != "My Goblin" {
		t.Fatalf(`Unmatched name`)
	}
	if cs.Strength != 5 {
		t.Fatalf(`Unmatched strength`)
	}
}
