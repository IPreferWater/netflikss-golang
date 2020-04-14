package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//Configuration todo doc of configuration
type Configuration struct {
	StockPath string `json:"stockPath"`
	CurrentPath string `json:"currentPath"`
}

func GetConfigurationByteFormat() []byte {

	content, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func ReadConfigurationFile() Configuration {

	configurationByte, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	configuration := Configuration{}
	err = json.Unmarshal(configurationByte, &configuration)
	if err != nil {
		fmt.Printf("Failed to unmarshal content %s, the error is %v", string(configurationByte), err)
	}

	return configuration
}

func SetConfiguration(configuration Configuration) {

	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = ioutil.WriteFile("configuration.json", file, 0644)
}
