package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	config := Configuration{}
	err := gonfig.GetConf("echo-rest/config/config.json", &config)
	if err != nil{
		panic(err)
	}

	return config
}