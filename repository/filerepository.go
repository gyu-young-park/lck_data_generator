package repository

import (
	"fmt"
	"io/ioutil"
	"os"
)

type RECORDING_KEY_TYPE string

const (
	ALL_MATCH = RECORDING_KEY_TYPE("all-match")
	ALL_TEAM_WITH_SEASON = RECORDING_KEY_TYPE("team-season")
	ALL_SEASON_WITH_TEAM = RECORDING_KEY_TYPE("season-team")
)

const DEFAULT_RECORDING_JSON_MATCH_FILE_NAME = "./all-match.json"
const DEFAULT_RECORDING_JSON_TEAM_FILE_NAME = "./all-team.json"
const DEFAULT_RECORDING_JSON_SEASON_FILE_NAME = "./all-season.json"

type FileRepository struct {
}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (f *FileRepository) Store(key string, rawData interface{}) error {
	switch RECORDING_KEY_TYPE(key) {
	case ALL_MATCH:
		f.storeJSON(rawData, DEFAULT_RECORDING_JSON_MATCH_FILE_NAME)
	case ALL_TEAM_WITH_SEASON:
		f.storeJSON(rawData, DEFAULT_RECORDING_JSON_TEAM_FILE_NAME)
	case ALL_SEASON_WITH_TEAM:
		f.storeJSON(rawData, DEFAULT_RECORDING_JSON_SEASON_FILE_NAME)
	default:
		return fmt.Errorf("Not support type %s", string(key))
	}
	return nil
}

func (f *FileRepository) Get(key string) (string, error) {
	switch RECORDING_KEY_TYPE(key) {
	case ALL_MATCH:
		return f.getJSONFile(DEFAULT_RECORDING_JSON_MATCH_FILE_NAME)
	case ALL_TEAM_WITH_SEASON:
		return f.getJSONFile(DEFAULT_RECORDING_JSON_TEAM_FILE_NAME)
	case ALL_SEASON_WITH_TEAM:
		return f.getJSONFile(DEFAULT_RECORDING_JSON_SEASON_FILE_NAME)
	default:
		return "", fmt.Errorf("Not support type %s", string(key))
	}
}

func (f *FileRepository) storeJSON(rawData interface{}, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	data, ok := rawData.(string)
	if !ok {
		return fmt.Errorf("It can't be convert rawData to string!")
	}
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) getJSONFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}