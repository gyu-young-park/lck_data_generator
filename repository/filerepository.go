package repository

import (
	"fmt"
	"io/ioutil"
	"os"
)

type RECORDING_KEY_TYPE string

const (
	ALL_MATCH = RECORDING_KEY_TYPE("all-match")
)

const DEFAULT_RECORDING_JSON_FILE_NAME = "./all-match.json"

type FileRepository struct {
	fileName string
}

func NewFileRepository(fileName string) *FileRepository {
	return &FileRepository{
		fileName: fileName,
	}
}

func (f *FileRepository) Store(key string, rawData interface{}) error {
	switch RECORDING_KEY_TYPE(key) {
	case ALL_MATCH:
		f.storeAllMatch(rawData)
	default:
		return fmt.Errorf("Not support type %s", string(key))
	}
	return nil
}

func (f *FileRepository) storeAllMatch(rawData interface{}) error {
	file, err := os.OpenFile(f.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

func (f *FileRepository) Get(key string) (string, error) {
	switch RECORDING_KEY_TYPE(key) {
	case ALL_MATCH:
		return f.GetAllMatch()
	default:
		return "", fmt.Errorf("Not support type %s", string(key))
	}
}

func (f *FileRepository) GetAllMatch() (string, error) {
	data, err := ioutil.ReadFile(DEFAULT_RECORDING_JSON_FILE_NAME)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
