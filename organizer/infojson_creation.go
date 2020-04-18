package organizer

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ipreferwater/netflikss-golang/graph/model"
)

//BuildInfoJSONFile build info.json files
func BuildInfoJSONFile() {
	path := filepath.Join(FileServerPath, StockPath)
	directories := getAllDirectories(path)

	for _, directory := range directories {
		infoJSONPath := filepath.Join(path, directory.Name(), infoJSONFileName)
		serieToCreate := model.Serie{
			DirectoryName: directory.Name(),
			Label:         directory.Name(),
			StockPath:     StockPath,
		}

		if !fileExists(infoJSONPath) {
			seasonDirPath := filepath.Join(path, directory.Name())

			seasonsDirs := getAllDirectories(seasonDirPath)
			seasonsToCreate := make([]*model.Season, 0)

			for _, seasonDir := range seasonsDirs {
				fileName := seasonDir.Name()

				//fmt.Printf("? %s \n",seasonDir.Name())
				episodesPath := filepath.Join(path, directory.Name(), fileName)
				episodes := getAllFiles(episodesPath)
				episodeCreated := createAllEpisode(episodes)

				guessNumber := guessNumber(fileName)
				number, _ := strconv.Atoi(guessNumber)

				//TODO: do we really need a label on the season ?
				newSeason := model.Season{
					DirectoryName: seasonDir.Name(),
					Number:        number,
					Label:         guessNumber,
					Episodes:      episodeCreated,
				}
				seasonsToCreate = append(seasonsToCreate, &newSeason)

			}

			serieToCreate.Seasons = seasonsToCreate

			file, _ := json.MarshalIndent(serieToCreate, "", " ")
			_ = ioutil.WriteFile(infoJSONPath, file, 0644)
		}
	}
}

func createAllEpisode(episodes []os.FileInfo) []*model.Episode {
	episodesToCreate := make([]*model.Episode, 0)
	for _, episode := range episodes {
		//fmt.Printf("? %s \n",episode.Name())
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
