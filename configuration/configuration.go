package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Configuration todo doc of configuration
type Configuration struct {
	StockPath string `json:"stockPath"`
}

func GetConfigurationByteFormat() []byte {

	content, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	return content
}
func ReadConfigurationFile() []byte {

	content, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	/*configuration := Configuration{}
	err = json.Unmarshal(content, &configuration)
	if err != nil {
		fmt.Printf("Failed to unmarshal content %s, the error is %v", string(content), err)
	}*/

	return content
}

func SetConfiguration(configuration Configuration) {

	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = ioutil.WriteFile("configuration.json", file, 0644)
}
