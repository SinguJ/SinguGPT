package store

import "SinguGPT/models"

const configPath = "./config.yml"

var Config = &models.Config{}

func init() {
	err := loadYaml(configPath, Config)
	if err != nil {
		panic(err)
	}
}
