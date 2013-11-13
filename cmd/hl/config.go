package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ToolConfig struct {
	Url   string `json:"api_url"`
	Key   string `json:"api_key"`
	Debug bool   `json:"debug"`
}

// Get the path prefix for the tool
func GetPath() string {
	prefix := os.Getenv("HLPATH")

	if prefix == "" {
		prefix = filepath.Join(os.Getenv("HOME"), ".hl")
	}

	return prefix
}

func GetConfig() (ToolConfig, error) {
	var conf ToolConfig
	confPath := filepath.Join(GetPath(), "config.json")

	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		return conf, err
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		return conf, fmt.Errorf("can't decode %q into ToolConfig: %s", confPath, err)
	}

	return conf, nil
}
