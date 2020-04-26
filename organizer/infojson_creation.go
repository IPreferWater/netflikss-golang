package organizer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ipreferwater/netflikss-golang/configuration"
	"github.com/ipreferwater/netflikss-golang/di"
	"github.com/ipreferwater/netflikss-golang/graph/model"
)

//BuildInfoJSONFile build info.json files
func BuildInfoJSONFile() {
	path := configuration.GetFileAndStockPath()
	directories := getAllDirectories(path)

	for _, directory := range directories {
		directoryPath := filepath.Join(path, directory.Name())
		infoJSONPath := filepath.Join(directoryPath, infoJSONFileName)

		if !fileExists(infoJSONPath) {

			allFiles := getAllFiles(directoryPath)

			guessType := guessType(allFiles)

			if guessType == "serie" {
				createSerie(allFiles, directory, infoJSONPath, directoryPath)
			} else if guessType == "movie" {
				fmt.Println("need to create a movie")
			}

		}
	}
}

func guessType(files []os.FileInfo) string {
	//TODO: make an enum with iota

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == ".mp4" || ext == ".MP4" {
			return "movie"
		}
	}

	return "serie"
}

func findImage(files []os.FileInfo) string {
	allImages := filterByImg(files)
	if len(allImages) < 0 {
		//no image found
		return ""
	}
	return allImages[0].Name()
}

func createSerie(files []os.FileInfo, directory os.FileInfo, infoJSONPath string, directoryPath string) {

	seasonsDirs := filterByDirectory(files)

	img := findImage(files)
	serieToCreate := model.Serie{
		DirectoryName: directory.Name(),
		Label:         directory.Name(),
		StockPath:     di.Configuration.StockPath,
		Img:           img,
	}

	seasonsToCreate := make([]*model.Season, 0)

	for _, seasonDir := range seasonsDirs {
		fileName := seasonDir.Name()

		episodesPath := filepath.Join(directoryPath, fileName)
		episodes := getAllFiles(episodesPath)
		episodeCreated := createAllEpisode(episodes)

		guessNumber := guessNumber(fileName)
		number, _ := strconv.Atoi(guessNumber)

		newSeason := model.Season{
			DirectoryName: seasonDir.Name(),
			Number:        number,
			Label:         guessNumber,
			Episodes:      episodeCreated,
		}
		seasonsToCreate = append(seasonsToCreate, &newSeason)

	}

	serieToCreate.Seasons = seasonsToCreate

	writeInfoJSONFile(serieToCreate, infoJSONPath)
	/*file, _ := json.MarshalIndent(serieToCreate, "", " ")
	_ = ioutil.WriteFile(infoJSONPath, file, 0644)*/

}

func writeInfoJSONFile(object interface{}, infoJSONPath string) {
	file, _ := json.MarshalIndent(object, "", " ")
	_ = ioutil.WriteFile(infoJSONPath, file, 0644)
}

func createAllEpisode(episodes []os.FileInfo) []*model.Episode {
	episodesToCreate := make([]*model.Episode, 0)
	for _, episode := range episodes {
		fileName := episode.Name()

		guessNumber := guessNumber(fileName)
		number, _ := strconv.Atoi(guessNumber)

		//TODO: we should initialize a label without extenstion
		newEpisode := model.Episode{
			FileName: fileName,
			Number:   number,
			Label:    fileName,
		}
		episodesToCreate = append(episodesToCreate, &newEpisode)
	}
	return episodesToCreate
}
