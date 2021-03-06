package organizer

import (
	"encoding/json"
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
			var objectToWrite interface{}
			allFiles := getAllFiles(directoryPath)
			guessType := guessType(allFiles)
			img := findImage(allFiles)

			info := model.Info{
				Directory: directory.Name(),
				Label:     directory.Name(),
				StockPath: di.Configuration.StockPath,
				Img:       img,
				Type:      guessType,
			}

			if guessType == "serie" {
				serie := createSerie(allFiles, directoryPath)
				serie.Info = &info
				objectToWrite = serie

			} else if guessType == "movie" {
				movie := createMovie(allFiles)
				movie.Info = &info
				objectToWrite = movie
			}
			writeInfoJSONFile(objectToWrite, infoJSONPath)
		}
	}
}

func createMovie(files []os.FileInfo) model.Movie {
	movieToCreate := model.Movie{}
//TODO: for the serie we create the episode name depending on fileName but for movie it depend on directoryName
//we shouldn't use 2 solutions
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if isExtensionVideo(ext) {
			movieToCreate.FileName = file.Name()
		}
	}
	return movieToCreate
}

func guessType(files []os.FileInfo) string {
	//TODO: make an enum with iota

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if isExtensionVideo(ext) {
			return "movie"
		}
	}

	return "serie"
}

func writeInfoJSONFile(object interface{}, infoJSONPath string) {
	file, _ := json.MarshalIndent(object, "", " ")
	_ = ioutil.WriteFile(infoJSONPath, file, 0644)
}

func findImage(files []os.FileInfo) string {
	allImages := filterByImg(files)
	if len(allImages) <= 0 {
		//no image found
		return ""
	}
	return allImages[0].Name()
}

func createSerie(files []os.FileInfo, directoryPath string) model.Serie {

	seasonsDirs := filterByDirectory(files)

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

	serieToCreate := model.Serie{
		Seasons: seasonsToCreate,
	}
	return serieToCreate
}

func createAllEpisode(episodes []os.FileInfo) []*model.Episode {
	episodesToCreate := make([]*model.Episode, 0)
	for _, episode := range episodes {
		fileName := episode.Name()
		ext := filepath.Ext(fileName)
		if !isExtensionVideo(ext) {
			continue
		}

		guessNumber := guessNumber(fileName)
		number, _ := strconv.Atoi(guessNumber)

		newEpisode := model.Episode{
			FileName: fileName,
			Number:   number,
			Label:    removeExtFromFilename(fileName),
		}
		episodesToCreate = append(episodesToCreate, &newEpisode)
	}
	return episodesToCreate
}
