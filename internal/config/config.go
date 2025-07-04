package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg Config

func LoadConfig() {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
