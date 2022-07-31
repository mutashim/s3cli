package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_FILE = "./s3cli.yml"

type Config struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Endpoint  string `yaml:"endpoint"`
	Region    string `yaml:"region"`
}

func LoadConfig() *Config {
	// check file
	_, err := os.Stat(CONFIG_FILE)
	if os.IsNotExist(err) {
		log.Fatal("configuration file not found")
	}

	// open file
	data, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal
	cfg := Config{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
