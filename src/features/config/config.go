package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host       string
	Port       string
	DBUser     string
	DBName     string
	DBPassword string
}

func LoadConfig(filePath string) (config Config) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	}
	return
}
