package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	engine "github.com/ercasta/allsoulsrun/pkg/engine"
	ocf "github.com/hamba/avro/v2/ocf"
)

type AvroRecoder struct {
	Basepath string
	types    map[engine.TrackerType]*ocf.Encoder
}

func (rec *AvroRecoder) Init(trackertype engine.TrackerType, schema engine.AvroSchema) {
	if rec.types == nil {
		rec.types = make(map[engine.TrackerType]*ocf.Encoder, 20)
	}

	avrofolder := filepath.Join(rec.Basepath, string(trackertype))
	err := os.MkdirAll(avrofolder, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	f, err := os.Create(filepath.Join(avrofolder, "tracked.avro"))

	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}

	enc, err := ocf.NewEncoder(string(schema), f)
	if err != nil {
		fmt.Printf("Error creating encoder: %v\n", err)
		return
	}

	rec.types[trackertype] = enc
}

func (rec *AvroRecoder) Record(trackertype engine.TrackerType, trackeddata any) {
	enc := rec.types[trackertype]
	print("Writing to file\n")
	err := enc.Encode(trackeddata)
	if err != nil {
		fmt.Printf("Error encoding data: %v\n", err)
	}
}

func (rec *AvroRecoder) Close() {
	for _, enc := range rec.types {
		enc.Close()
		if err := enc.Flush(); err != nil {
			log.Fatal(err)
		}

		//TODO
		// if err := f.Sync(); err != nil {
		// 	log.Fatal(err)
		// }
	}
}
