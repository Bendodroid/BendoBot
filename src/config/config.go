package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Bendodroid/BendoBot/src/errors"
)

// loadConfig reads the Config from a JSON file
func LoadConfig(filename string, config *Config) {
	// Read []byte from file
	dat, err := ioutil.ReadFile(filename)
	errors.CheckErr(err, "Error reading from configfile")

	// Parse json
	err = json.Unmarshal(dat, config)
	errors.CheckErr(err, "Failed to parse json")
}
