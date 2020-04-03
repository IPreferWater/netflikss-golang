package organizer

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readAllInside(directory string) {
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

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func guessNumber(fileName string) string {
	lastIndexOfPoint := strings.LastIndex(fileName, ".")
	if lastIndexOfPoint < 0 {
		//the file has no extension
		lastIndexOfPoint = len(fileName)
	}

	noExtension := fileName[0:lastIndexOfPoint]

	for i := 1; i < len(noExtension); i++ {
		stringToTest := noExtension[lastIndexOfPoint-i : len(noExtension)]
		if !isNumeric(stringToTest) {

			if i == 1 {
				return ""
			}
			//the previous was the correct one
			finalNumberString := noExtension[lastIndexOfPoint-i+1 : len(noExtension)]

			//we convert to int to remove useless 0
			floatNumber, _ := strconv.Atoi(finalNumberString)

			return strconv.Itoa(floatNumber)
		}
	}
	return ""
}
