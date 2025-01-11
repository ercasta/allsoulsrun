package gamemanager

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Adventure struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Levels      []Level `json:"levels"`
}

type Level struct {
	Name          string         `json:"name"`
	RetreatOnLoss bool           `json:"retreat_on_loss"`
	NextLevel     string         `json:"next_level,omitempty"`
	NumEncounters EncounterRange `json:"num_encounters,omitempty"`
	Encounters    []Encounter    `json:"encounters"`
}

type EncounterRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Encounter struct {
	Name     string    `json:"name"`
	Monsters []Monster `json:"monsters"`
}

type Monster struct {
	Enemy string `json:"enemy"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
}

func readAdventure(dataFolder string, adventure string) (*Adventure, error) {
	data, err := os.ReadFile(filepath.Join(dataFolder, adventure))
	if err != nil {
		return nil, err
	}

	var adventureLoaded = Adventure{}
	err = json.Unmarshal(data, &adventureLoaded)
	if err != nil {
		return nil, err
	}

	return &adventureLoaded, nil
}
