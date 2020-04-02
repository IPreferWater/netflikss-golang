package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//testDirectory()
	//createInfoJson()

	go func() {
		//graphql.StartServerGraphQL()
	}()

	fs := http.FileServer(http.Dir("./stock"))
	http.Handle("/", fs)

	log.Println("Listening onn :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createInfoJson(fileName string) {
	filess, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read all directories from %s \n", fileName)
	for _, f := range filess {
		if f.IsDir() {
			//seasons directories
			fmt.Println(f.Name())
			//readAllInside(fileName + "/" + f.Name())
		}
	}
}

func testDirectory() { /*
		filess, err := ioutil.ReadDir("./stock")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range filess {
			if f.IsDir() {
				fmt.Println(f.Name())
				 infoJsonFile := "./stock/"+f.Name()+"/info.json"
				content, err := ioutil.ReadFile(infoJsonFile)

				if err != nil {
					fmt.Println("need to be created:", err)
					createInfoJson("./stock/"+f.Name())

					continue
				}
				result := organizer.Info{}
				err = json.Unmarshal(content, &result)
				if err != nil {
					fmt.Print("Error:", err)
					fmt.Print("Failed to unmarshal content %s, the error is %v", string(content), err)

					continue
				}
			}*/
}
