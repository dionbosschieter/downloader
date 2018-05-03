package bot

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "fmt"
    "os"
)

type Settings struct {
	TelegramToken   string   `yaml:"telegramToken"`
	TransmissionUrl string   `yaml:"transmissionUrl"`
	SeriePath       string   `yaml:"seriePath"`
	MoviePath       string   `yaml:"moviePath"`
	MasterChatId    int64    `yaml:"masterChatId"`
    SearchProviders []string `yaml:"searchProviders"`
    SearchPostfixes []string `yaml:"searchPostfixes"`
}

func (settings *Settings) FileExists() bool {
    if _, err := os.Stat("settings.yaml"); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

func (settings *Settings) Parse() *Settings {
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
