package main

import (
	"./files"
	"./graphql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	/*var files []string

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
	}*/



	filess, err := ioutil.ReadDir("./stock")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range filess {
		if f.IsDir() {
			fmt.Println(f.Name())
			content, err := ioutil.ReadFile("./stock/"+f.Name()+"/info.json")

			if err != nil {
				fmt.Print("Error:", err)
				continue
			}
			result := files.Info{}
			err = json.Unmarshal(content, &result)
			if err != nil {
				fmt.Print("Error:", err)
				fmt.Print("Failed to unmarshal content %s, the error is %v", string(content), err)

				continue
			}
			fmt.Print("ok:", result.Seasons)
		}
	}


}

