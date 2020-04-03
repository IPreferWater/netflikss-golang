package organizer

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"os"
)

//TODO: we need to let the user choose his media folder
func getAllInStockFolder() []os.FileInfo{
	files, err := ioutil.ReadDir("../stock")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func filterByDirectory(files []os.FileInfo) []os.FileInfo {
    		//filter to get only the directories
			directories := make([]os.FileInfo, 0)
			for _, file := range files {
				if file.IsDir() {
					directories = append(directories, file)
				}
			}
			return directories
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

func removeUselessZero(s string) string {
	intNumber, err := strconv.Atoi(s)
	if err != nil {
		return ""
	}
	return strconv.Itoa(intNumber)
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
			//the last character was not a digit, the filename doesn't end by a number
			if i == 1 {
				return ""
			}
			//the previous was the correct one
			finalNumberString := noExtension[lastIndexOfPoint-i+1 : len(noExtension)]
			return removeUselessZero(finalNumberString)
		}
	}
	return removeUselessZero(noExtension)
}
