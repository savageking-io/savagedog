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

type ClientConfig struct {
	Dest    string `yaml:"dest"`
	From    string `yaml:"from"`
	Header  string `yaml:"header"`
	Content string `yaml:"content"`
	Sender  string `yaml:"sender"`
	Fields  string `yaml:"fields"`
	Color   string `yaml:"color"`
}

type DiscordConfig struct {
	Token string `yaml:"token"`
}

type ServiceConfig struct {
	Name        string `yaml:"name"`
	Channel     string `yaml:"channel"`
	Author      string `yaml:"author"`
	AuthorURL   string `yaml:"author_url"`
	AuthorImage string `yaml:"author_image"`
}

func ReadConfig(filepath string, out interface{}) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("configuration file not found: %s", filepath))
	}

	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return errors.New(fmt.Sprintf("error reading configuration file: %v", err))
	}

	err = yaml.Unmarshal(yamlFile, out)
	if err != nil {
		return errors.New(fmt.Sprintf("error parsing YAML configuration: %v", err))
	}

	return nil
}
