package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ipreferwater/netflikss-golang/graph"
	"github.com/ipreferwater/netflikss-golang/graph/generated"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	c := cors.New(cors.Options{
		//AllowedOrigins: []string{"Access-Control-Allow-Origin", "*"},
		AllowedOrigins: []string{"http://localhost:64594", "*"},
	})
	http.Handle("/playground", c.Handler(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", c.Handler(srv))
	http.Handle("/usb", http.FileServer(http.Dir("/dev")))
	http.Handle("/", http.FileServer(http.Dir(user.HomeDir)))
	http.HandleFunc("/stock", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
