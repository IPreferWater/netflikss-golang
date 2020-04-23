package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ipreferwater/netflikss-golang/api"
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

	configuration.InitUserVariable()
	configuration.InitGlobalVariable()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:64594", "*"},
	})
	http.Handle("/playground", c.Handler(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", c.Handler(srv))
	http.Handle("/usb", http.FileServer(http.Dir("/dev")))
	http.Handle("/", http.FileServer(http.Dir(organizer.FileServerPath)))
	http.Handle("/stockpath", c.Handler(http.HandlerFunc(api.StockPath)))
	http.Handle("/directorieslist", c.Handler(http.HandlerFunc(api.DirectoriesList)))

	log.Printf("connect to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}




