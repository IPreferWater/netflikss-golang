package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"github.com/ipreferwater/netflikss-golang/di"
	"github.com/ipreferwater/netflikss-golang/graph/model"
)

func GetFileAndStockPath() string {
	return filepath.Join(di.Configuration.FileServerPath, di.Configuration.StockPath)
}

//InitUserVariable init the user to have the directory Path
func InitUserVariable() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	di.User = user
}

func InitGlobalVariable() {
	//init the path from config
	configuration := ReadConfigurationFile()
	di.Configuration.ServerConfiguration = &model.ServerConfiguration{}

	di.Configuration.StockPath = configuration.StockPath
	di.Configuration.ServerConfiguration.Port = configuration.ServerConfiguration.Port
	di.Configuration.ServerConfiguration.AllowedOrigin = configuration.ServerConfiguration.AllowedOrigin

	if configuration.FileServerPath == "" {
		di.Configuration.FileServerPath = di.User.HomeDir

		print("set " + di.Configuration.FileServerPath)
	} else {
		di.Configuration.FileServerPath = configuration.FileServerPath
		print("set " + di.Configuration.FileServerPath)
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
