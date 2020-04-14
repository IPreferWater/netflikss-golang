package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ipreferwater/netflikss-golang/configuration"
	"github.com/ipreferwater/netflikss-golang/graph"
	"github.com/ipreferwater/netflikss-golang/graph/generated"
	"github.com/ipreferwater/netflikss-golang/organizer"
	"github.com/rs/cors"
)

const defaultPort = "7171"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	//init the path from config
	configuration := configuration.ReadConfigurationFile()
	if configuration.StockPath == "" {
		organizer.StockPath = user.HomeDir
	} else {
		organizer.StockPath = configuration.StockPath
	}

	if configuration.FileServerPath == "" {
		organizer.FileServerPath = user.HomeDir
	} else {
		organizer.FileServerPath = configuration.FileServerPath
	}


	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:64594", "*"},
	})
	http.Handle("/playground", c.Handler(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", c.Handler(srv))
	http.Handle("/usb", http.FileServer(http.Dir("/dev")))
	http.Handle("/", http.FileServer(http.Dir(configuration.FileServerPath)))
	http.HandleFunc("/stockpath", stockPath)
	http.HandleFunc("/directorieslist", directoriesList)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:64594")
}

func stockPath(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)

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

func directoriesList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	var s string
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join(organizer.StockPath, s)
	println(path)

	 listDirectoriesName := organizer.GetAllDirectoriesName(path)
	 jsonListDirectoriesName, err := json.Marshal(listDirectoriesName)

	 if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	 }
	 w.Write(jsonListDirectoriesName)
}
