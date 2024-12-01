package engine

import (
	"encoding/json"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	sample := `{
    "sections": 
        [
            {
                "name": "section1",
                "type": "enemychooser",
                "rules": [
                    {"name":"target weakest", "condition": null, "action": [
                        {
                            "key": "target",
                            "value": "weakest"
                        }       
                ]}
                ]
            }
        ]
	}`

	var ruleList RuleList
	err := json.Unmarshal([]byte(sample), &ruleList)
	if err != nil {
		t.Fatalf(`Unable to unmarshal`)
	}

}
