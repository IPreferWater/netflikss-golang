package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

	"github.com/ipreferwater/netflikss-golang/organizer"
)

//Configuration todo doc of configuration
type Configuration struct {
	StockPath string `json:"stockPath"`
	FileServerPath string `json:"fileServerPath"`
}

func InitGlobalVariable(){

	//TODO: we should init this only once
	user, err := user.Current()
	if err != nil {
		panic(err)
	}//end TODO
	
		//init the path from config
		configuration := ReadConfigurationFile()
		organizer.StockPath = configuration.StockPath
	
		if configuration.FileServerPath == "" {
			organizer.FileServerPath = user.HomeDir
			print("set " + organizer.FileServerPath)
		} else {
			organizer.FileServerPath = configuration.FileServerPath
			print("set " + organizer.FileServerPath)
		}

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
