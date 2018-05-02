package main

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "fmt"
)

type Settings struct {
	TelegramToken   string `yaml:"telegramToken"`
	TransmissionUrl string `yaml:"transmissionUrl"`
	SeriePath       string `yaml:"seriePath"`
	MoviePath       string `yaml:"moviePath"`
	MasterChatId    string `yaml:"masterChatId"`
}

func (settings *Settings) Read() *Settings {
	yamlFile, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		panic(fmt.Sprintf("err opening settings: #%v ", err))
	}
	err = yaml.Unmarshal(yamlFile, settings)
	if err != nil {
		panic(fmt.Sprintf("err parsing settings: %v", err))
	}

	return settings
}
