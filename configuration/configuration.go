package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

	"github.com/ipreferwater/netflikss-golang/organizer"
)

//the configuration will be handle by graphql in the futur,
// the object used will be the generated one
type Configuration struct {
	StockPath      string `json:"stockPath"`
	FileServerPath string `json:"fileServerPath"`
	Port           string `json:"port"`
	AllowedOrigin  string `json:"allowedOrigin"`
}

//InitUserVariable init the user to have the directory Path
func InitUserVariable() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	organizer.User = user
}

func InitGlobalVariable() {
	//init the path from config
	configuration := ReadConfigurationFile()
	organizer.StockPath = configuration.StockPath
	organizer.Port = configuration.Port
	organizer.AllowedOrigin = configuration.AllowedOrigin

	if configuration.FileServerPath == "" {
		organizer.FileServerPath = organizer.User.HomeDir

		print("set " + organizer.FileServerPath)
	} else {
		organizer.FileServerPath = configuration.FileServerPath
		print("set " + organizer.FileServerPath)
	}

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
