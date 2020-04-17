package organizer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ipreferwater/netflikss-golang/graph/model"
)

const (
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

func getAllFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func getAllDirectories(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return filterByDirectory(files)
}

//TODO: we loop 2 time, to improve
func GetAllDirectoriesName(path string) []string {

	listDirectoryNames := make([]string, 0)

	directories := getAllDirectories(path)
	for _, directory := range directories {
		listDirectoryNames = append(listDirectoryNames, directory.Name())
	}
	return listDirectoryNames
}

//filterByDirectory filter to get only the directories
func filterByDirectory(files []os.FileInfo) []os.FileInfo {
	directories := make([]os.FileInfo, 0)
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}
	return directories
}

//ReadAllInside read all info.json files
func ReadAllInside() []model.Serie {
	path := filepath.Join(FileServerPath, StockPath)
	files := getAllDirectories(path)

	series := make([]model.Serie, 0)

	for _, directory := range files {
		infoJSONPath := filepath.Join(path, directory.Name(), infoJSONFileName)
		if fileExists(infoJSONPath) {
			content, err := ioutil.ReadFile(infoJSONPath)
			if err != nil {
				log.Fatal(err)
			}
			serie := model.Serie{}
			err = json.Unmarshal(content, &serie)
			if err != nil {
				fmt.Printf("Failed to unmarshal content %s, the error is %v", string(content), err)
			}
			series = append(series, serie)
		}
	}
	return series
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
