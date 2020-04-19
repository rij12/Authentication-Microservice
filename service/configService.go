package service

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rij12/Authentication-Microservice/models"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var cfg *models.Config

type ConfigurationService struct{}

func (config *ConfigurationService) getEnv(cfg *models.Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func (config *ConfigurationService) getFromFile(filename string, cfg *models.Config) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func (config *ConfigurationService) GetConfig() models.Config {
	if cfg == nil {
		cfg = new(models.Config)
		// TODO Pass in config YML as a flag
		config.getFromFile("config.yml", cfg)
		config.getEnv(cfg)
	}
	return *cfg
}
