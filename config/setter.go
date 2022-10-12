package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigSetter interface{
	Set(path string, config *Config) error
}

type ConfigSetterJSON struct {

}

func NewConfigSetterJSON() *ConfigSetterJSON {
	return &ConfigSetterJSON{}
}

func (c *ConfigSetterJSON) Set(path string, config *Config) error{
	data, err := os.Open(path)
	if err != nil {
		return err
	}
	byteData, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteData, config)
	if err != nil {
		return err
	}
	return nil
}