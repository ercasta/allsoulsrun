package gamemanager

type PartyMember struct {
	Name  string `json:"name"`
	Build string `json:"build"`
	Line  string `json:"line"`
}

type RunConfig struct {
	Party     []PartyMember `json:"party"`
	Adventure string        `json:"adventure"`
	Trackers  []string      `json:"trackers"`
}
