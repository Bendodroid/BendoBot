package util

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Bendodroid/BendoBot/errors"
)

// LoadJSON loads the json file into target
func LoadJSON(filename string, target interface{}) {
	// Read []byte from file
	dat, err := ioutil.ReadFile(filename)
	errors.Check(err, "Error reading from file")

	// Parse json
	err = json.Unmarshal(dat, &target)
	errors.Check(err, "Failed to parse json")
}
