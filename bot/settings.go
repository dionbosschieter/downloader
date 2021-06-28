package bot

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
)

type Settings struct {
    Path            string
	TelegramToken   string   `yaml:"telegramToken"`
	TransmissionUrl string   `yaml:"transmissionUrl"`
	SeriePath       string   `yaml:"seriePath"`
	MoviePath       string   `yaml:"moviePath"`
    SearchProviders []string `yaml:"searchProviders"`
    SearchPostfixes []string `yaml:"searchPostfixes"`
}

func (settings *Settings) FileExists(settingsPath string) bool {
    if _, err := os.Stat(settingsPath); err == nil {
        return true
    }

    return false
}

func (settings *Settings) Parse(settingsPath string) {
	yamlFile, err := ioutil.ReadFile(settingsPath)
	if err != nil {
		panic(fmt.Sprintf("err opening settings: #%v ", err))
	}
	err = yaml.Unmarshal(yamlFile, settings)
	if err != nil {
		panic(fmt.Sprintf("err parsing settings: %v", err))
	}
}
