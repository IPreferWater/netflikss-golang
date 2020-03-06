package main

import (
	"./graphql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)


func main() {
	testDirectory()

	go func() {
		graphql.StartServerGraphQL()
	}()

		fs := http.FileServer(http.Dir("./stock"))
		http.Handle("/", fs)

		log.Println("Listening on :8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
}

func testDirectory(){
	var files []string

	root := "./stock"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

