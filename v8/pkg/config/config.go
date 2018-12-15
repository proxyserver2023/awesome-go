package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Load - loads the configuration into a struct which are
// consist of different Categories  like server, jwt, application, Database
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct, %v", err)
	}

	return cfg, nil
}
