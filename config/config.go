package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Endpoints map[string]EndpointConfig `json:"endpoints"`
}

type EndpointConfig struct {
	Limit           int `json:"limit"`
	IntervalSeconds int `json:"intervalSeconds"`
}

func LoadConfig(filename string) (Config, error) {
	var conf Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return conf, fmt.Errorf("failed to read config file: %v", err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return conf, fmt.Errorf("failed to unmarshal config data: %v", err)
	}
	
	return conf, nil
}
