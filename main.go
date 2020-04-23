package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ipreferwater/netflikss-golang/api"
	"github.com/ipreferwater/netflikss-golang/configuration"
	"github.com/ipreferwater/netflikss-golang/di"
	"github.com/ipreferwater/netflikss-golang/graph"
	"github.com/ipreferwater/netflikss-golang/graph/generated"
	"github.com/rs/cors"
)

func main() {
	configuration.InitUserVariable()
	configuration.InitGlobalVariable()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{di.Configuration.ServerConfiguration.AllowedOrigin, "*"},
	})
	http.Handle("/playground", c.Handler(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", c.Handler(srv))
	http.Handle("/usb", http.FileServer(http.Dir("/dev")))
	http.Handle("/", http.FileServer(http.Dir(di.Configuration.FileServerPath)))
	http.Handle("/stockpath", c.Handler(http.HandlerFunc(api.StockPath)))
	http.Handle("/directorieslist", c.Handler(http.HandlerFunc(api.DirectoriesList)))

	log.Printf("connect to http://localhost:%s/", di.Configuration.ServerConfiguration.Port)
	log.Fatal(http.ListenAndServe(":"+di.Configuration.ServerConfiguration.Port, nil))
}
