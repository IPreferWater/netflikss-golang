package organizer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ipreferwater/netflikss-golang/configuration"
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

//TODO we should allow more format than only .jpg
func filterByImg(files []os.FileInfo) []os.FileInfo {
	directories := make([]os.FileInfo, 0)
	for _, file := range files {

		if filepath.Ext(file.Name()) == ".jpg" {
			directories = append(directories, file)
		}
	}
	return directories
}

//ReadAllInside read all info.json files
func ReadAllInfoJson() model.Data {
	path := configuration.GetFileAndStockPath()
	files := getAllDirectories(path)

	//series := make([]model.Serie, 0)

	data := model.Data{}

	for _, directory := range files {
		infoJSONPath := filepath.Join(path, directory.Name(), infoJSONFileName)
		if fileExists(infoJSONPath) {
			content, err := ioutil.ReadFile(infoJSONPath)
			if err != nil {
				log.Fatal(err)
			}

			var result map[string]interface{}
			json.Unmarshal([]byte(content), &result)

			if "serie" == result["type"] {
				serie := model.Serie{}
				err = json.Unmarshal(content, &serie)
				if err != nil {
					fmt.Printf("Failed to unmarshal serie content %s, the error is %v", string(content), err)
				}
				data.Series = append(data.Series, &serie)
			} else if "movie" == result["type"] {
				movie := model.Movie{}
				err = json.Unmarshal(content, &movie)
				if err != nil {
					fmt.Printf("Failed to unmarshal movie content %s, the error is %v", string(content), err)
				}
				data.Movies = append(data.Movies, &movie)
			}

		}
	}
	return data
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsPortNumber(s string) bool {
	number, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return number > 0 && number < 65535
}

//IsURL tests a string to determine if it is a well-structured url or not.
func IsURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
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
