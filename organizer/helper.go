package organizer

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"os"
	"path/filepath"
)

const (
	 stockPath string = "../stock"
	 infoJSONFileName string = "info.json"
)
	


// checks if a file exists and is not a directory
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

//TODO: we need to let the user choose his media folder
func getAllInStockFolder() []os.FileInfo{
	files, err := ioutil.ReadDir(stockPath)
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

//work method
func readAllInside() {
	files := getAllInStockFolder()
	filtered := filterByDirectory(files)
	for _, directory := range filtered {
		infoJSONPath := filepath.Join(stockPath, directory.Name(),infoJSONFileName)
		fmt.Printf("in the directory %s the info.json exit ? %t\n", directory.Name(), fileExists(infoJSONPath))
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
