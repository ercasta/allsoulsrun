package trackers

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
	gamecommon "github.com/ercasta/allsoulsrun/pkg/game/common"
	events "github.com/ercasta/allsoulsrun/pkg/game/events/common"
)

// Tracks basic character info

type CharacterName struct {
	Seq  e.GameEventSequence `avro:"seq"`
	Id   e.EntityID          `avro:"id"`
	Name string              `avro:"name"`
}

var Schema = `
{
    "type": "record",
    "name": "charactername",
    "namespace": "com.github.ercasta.allsoulsrun.common",
    "fields" : [
        {"name": "seq", "type": {
        "name": "myFixedSeq",
        "type": "fixed",
        "size": 8
      }},
        {"name": "id", "type": {
        "name": "myFixedId",
        "type": "fixed",
        "size": 8
      }},
		{"name": "name", "type": "string"}
    ]
}`

type CharacterRecorder struct{}

func (dc CharacterRecorder) GetType() e.TrackerType {
	return "CharacterName"
}

func (dc CharacterRecorder) GetSchema() e.AvroSchema {
	return e.AvroSchema(Schema)
}

func (dc CharacterRecorder) Track(seq e.GameEventSequence, ev e.Eventer, phase e.EventSequencePhase, gdv e.GameDataView, rec e.Recorder) {
	switch phase {
	case e.OnEvent:
		ce := ev.(events.CreateCharacterEvent)
		stats := gdv.GetComponent(ce.CharacterID, gamecommon.CharacterStats{}.GetComponentType()).(gamecommon.CharacterStats)
		record := CharacterName{Seq: seq, Id: ce.CharacterID, Name: stats.Name}
		rec.Record(dc.GetType(), record)
	}
}
