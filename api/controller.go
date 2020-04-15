package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/ipreferwater/netflikss-golang/configuration"
	"github.com/ipreferwater/netflikss-golang/organizer"
)

func StockPath(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		/*configuration := configuration.GetConfigurationByteFormat()
		w.Write(configuration)*/
		w.Write([]byte(organizer.StockPath))

	case "POST":
		//TODO: we update just the config.json, not the global variable
		newConfiguration := configuration.Configuration{}
		err := json.NewDecoder(r.Body).Decode(&newConfiguration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		configuration.SetConfiguration(newConfiguration)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func DirectoriesList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var b bytes.Buffer
	b.ReadFrom(r.Body)
	pathToExplore := b.String()
	path := filepath.Join(organizer.FileServerPath, pathToExplore)

	listDirectoriesName := organizer.GetAllDirectoriesName(path)
	jsonListDirectoriesName, err := json.Marshal(listDirectoriesName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.Write(jsonListDirectoriesName)
}