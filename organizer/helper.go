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
	stockPath        string = "../stock"
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

//TODO: we need to delete this method
func getAllInStockFolder() []os.FileInfo {
	files, err := ioutil.ReadDir(stockPath)
	if err != nil {
		log.Fatal(err)
	}
	return files
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

//ReadAllInside read all info.json files
func ReadAllInside() []model.Serie {
	files := getAllInStockFolder()
	filtered := filterByDirectory(files)
	series := make([]model.Serie, 0)

	for _, directory := range filtered {
		infoJSONPath := filepath.Join(stockPath, directory.Name(), infoJSONFileName)
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

//BuildInfoJSONFile build info.json files
func BuildInfoJSONFile(){
	directories := getAllDirectories(stockPath)

	for _, directory := range directories {
		infoJSONPath := filepath.Join(stockPath, directory.Name(), infoJSONFileName)
		serieToCreate := model.Serie{
			DirectoryName: directory.Name(),
			Label: directory.Name(),
		}

		if !fileExists(infoJSONPath) {
			seasonDirPath := filepath.Join(stockPath, directory.Name())

			seasonsDirs := getAllDirectories(seasonDirPath)
			seasonsToCreate := make([]*model.Season, len(seasonsDirs))

			for _, seasonDir := range seasonsDirs {
				fileName := seasonDir.Name()

				//fmt.Printf("? %s \n",seasonDir.Name())
				episodesPath := filepath.Join(stockPath, directory.Name(), fileName)
				episodes := getAllFiles(episodesPath)
				episodeCreated := createAllEpisode(episodes)

				guessNumber := guessNumber(fileName)
				number, _ := strconv.Atoi(guessNumber)

				//TODO: do we really need a label on the season ?
				newSeason := model.Season{
					DirectoryName: seasonDir.Name(),
					Number: number,
					Label: guessNumber,
					Episodes: episodeCreated,
				}
				seasonsToCreate = append(seasonsToCreate, &newSeason)

			}

			serieToCreate.Seasons=seasonsToCreate

			file, _ := json.MarshalIndent(serieToCreate, "", " ")
			_ = ioutil.WriteFile("info.json", file, 0644)
		}
	}
}

func createAllEpisode(episodes []os.FileInfo) []*model.Episode{
	episodesToCreate := make([]*model.Episode, len(episodes))
	for _, episode := range episodes {
		//fmt.Printf("? %s \n",episode.Name())
		fileName := episode.Name()
		
		guessNumber := guessNumber(fileName)
		number, _ := strconv.Atoi(guessNumber)

		//TODO: we should initialize a label without extenstion
		newEpisode := model.Episode{
			FileName: fileName,
			Number: number,
			Label: fileName,
		}
		episodesToCreate = append(episodesToCreate, &newEpisode)
	}
	return episodesToCreate
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
