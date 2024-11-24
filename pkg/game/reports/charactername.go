package reports

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tracker "github.com/ercasta/allsoulsrun/pkg/game/trackers"
	"github.com/gin-gonic/gin"

	ocf "github.com/hamba/avro/v2/ocf"
)

func GetReportData(runId string) []tracker.CharacterName {
	// do nothing}

	avrofile := filepath.Join("rundata/"+runId, string(tracker.CharacterRecorder{}.GetType()), "tracked.avro")

	f, err := os.Open(avrofile)

	if err != nil {
		fmt.Printf("Error Opening file: %v\n", err)
		return nil
	}
	defer f.Close()

	decoder, err := ocf.NewDecoder(f)
	if err != nil {
		fmt.Printf("Error creating encoder: %v\n", err)
		return nil
	}

	simple := tracker.CharacterName{}

	for decoder.HasNext() {

		err = decoder.Decode(&simple)
		if err != nil {
			log.Fatal(err)
		}

		// Do something with the data
	}

	return append(make([]tracker.CharacterName, 0), simple)
}

func GetCharacterNameReport(c *gin.Context) {
	runId := c.Param("runId")
	println(runId)
	characters := GetReportData(runId)
	name := characters[0].Name
	c.JSON(200, gin.H{"name": name})
}
