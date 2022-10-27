package repository

import (
	"fmt"
	"os"
)

const DEFAULT_RECORDING_FILE_NAME = "./temp.txt"
const DEFAULT_RECORDING_JSON_FILE_NAME = "./temp.json"
type FileRepository struct {
	fileName string
}

func NewFileRepository(fileName string) *FileRepository{
	return &FileRepository{
		fileName: fileName,
	}
}

func (f *FileRepository) Store(rawData interface{}) error {
	file, err := os.OpenFile(f.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	data ,ok := rawData.(string)
	if !ok {
		return fmt.Errorf("It can't be convert rawData to string!")
	}
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}