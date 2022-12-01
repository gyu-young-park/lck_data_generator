package config

const CONFIG_JSON_FILE_PATH = "./config.json"

type Config struct {
	Key             string `json:"key"`
	FirebaseKeyPath string `json:"firebase-key-path"`
	PlayList        string `json:"playlist"`
	CrawlerMode 	string `json:"crawler-mode"` // fandom, crawler
	setter          ConfigSetter
}

func NewConfig(setter ConfigSetter) *Config {
	configInstance := &Config{setter: setter}
	configInstance.setUpConfig()
	return configInstance
}

func (c *Config) setUpConfig() {
	err := c.setter.Set(CONFIG_JSON_FILE_PATH, c)
	if err != nil {
		panic(err)
	}
}
