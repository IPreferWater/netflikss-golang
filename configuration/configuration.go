package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"github.com/ipreferwater/netflikss-golang/graph/model"
)

func GetFileAndStockPath() string {
	return filepath.Join(Configuration.FileServerPath, Configuration.StockPath)
}

//InitUserVariable init the user to have the directory Path
func InitUserVariable() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	User = user
}

func InitGlobalVariable() {
	//init the path from config
	configuration := ReadConfigurationFile()
	Configuration.ServerConfiguration = &model.ServerConfiguration{}

	Configuration.StockPath = configuration.StockPath
	Configuration.ServerConfiguration.Port = configuration.ServerConfiguration.Port
	Configuration.ServerConfiguration.AllowedOrigin = configuration.ServerConfiguration.AllowedOrigin

	if configuration.FileServerPath == "" {
		Configuration.FileServerPath = User.HomeDir

		print("set " + Configuration.FileServerPath)
	} else {
		Configuration.FileServerPath = configuration.FileServerPath
		print("set " + Configuration.FileServerPath)
	}

}

func ReadConfigurationFile() model.Configuration {

	configurationByte, err := ioutil.ReadFile("configuration.json")
	if err != nil {
		fmt.Println("error can't parse tbhe configuration.js")
		log.Fatal(err)
	}
	configuration := model.Configuration{}
	err = json.Unmarshal(configurationByte, &configuration)
	if err != nil {
		fmt.Printf("Failed to unmarshal content %s, the error is %v", string(configurationByte), err)
	}

	return configuration
}

func SetConfiguration(configuration model.Configuration) {
	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = ioutil.WriteFile("configuration.json", file, 0644)
}
