package utils

import (
	"log"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	DNSServers []string `yaml:"dns_servers"`
	Cache      struct {
		Enabled bool `yaml:"enabled"`
		TTL     int  `yaml:"ttl"`
	} `yaml:"cache"`
	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
	GRPC struct {
		Port int `yaml:"port"`
	} `yaml:"grpc"`
}

func LoadConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
	return &config
}