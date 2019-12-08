package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Bendodroid/BendoBot/errors"
)

// loadConfig reads the Config from a JSON file
func Load(filename string) Config {
	// Read []byte from file
	dat, err := ioutil.ReadFile(filename)
	errors.Check(err, "Error reading from configfile")

	// Parse json
	var config Config
	err = json.Unmarshal(dat, &config)
	errors.Check(err, "Failed to parse json")
	return config
}
