package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Hostname string          `yaml:"hostname"`
	Port     uint16          `yaml:"port"`
	Discord  DiscordConfig   `yaml:"discord"`
	Services []ServiceConfig `yaml:"services"`
}

type DiscordConfig struct {
	Token string `yaml:"token"`
}

type ServiceConfig struct {
	Name    string `yaml:"name"`
	Channel string `yaml:"channel"`
}

func (c *Config) Read(filepath string) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("configuration file not found: %s", filepath))
	}

	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return errors.New(fmt.Sprintf("error reading configuration file: %v", err))
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return errors.New(fmt.Sprintf("error parsing YAML configuration: %v", err))
	}

	return nil
}
