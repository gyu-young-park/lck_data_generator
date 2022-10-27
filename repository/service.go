package repository

import (
	"fmt"
	"os"
)

type Repository interface{
	Store(interface{}) error
}

const DEFAULT_RECORDING_FILE_NAME = "./temp.txt"
type FileRepository struct {
	fileName string
}

func NewFileRepository(fileName string) *FileRepository{
	return &FileRepository{
		fileName: fileName,
	}
}

func (f *FileRepository) Store(rawData interface{}) error {
	file, err := os.Open(f.fileName)
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