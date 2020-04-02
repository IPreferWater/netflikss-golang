package organizer

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readAllInside(directory string){
	filess, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range filess {
		fmt.Println(f.Name())
		guessedNumber := guessNumber(f.Name())
		fmt.Printf("guessed number = %s \n", guessedNumber)
	}
}

func guessNumber(fileName string) string{
	lastIndexOfPoint := strings.LastIndex(fileName, ".")
	if lastIndexOfPoint<0 {
		return ""
	}

	removedFileExtension := fileName[0:lastIndexOfPoint]
	fmt.Println(removedFileExtension)

	lastCharacter := removedFileExtension[lastIndexOfPoint-1:lastIndexOfPoint]

	_, err :=strconv.Atoi(lastCharacter)
	if err != nil{
		return ""
	}

	twoLastCharacter := removedFileExtension[lastIndexOfPoint-2:lastIndexOfPoint]

	_, err2 :=strconv.Atoi(twoLastCharacter)
	if err2 != nil{
		return lastCharacter
	}

	return twoLastCharacter
}
