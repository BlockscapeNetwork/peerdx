package config

import (
	"github.com/BlockscapeLab/peerdx/files"
)

var config Config

func init() {
	config = createDefaultConfig()
}

// GetConfig returns copy of config
func GetConfig() Config {
	return config
}

// LoadConfig from dir. Filename must be included in dir.
func LoadConfig(dir string) error {
	c := &Config{}
	err := files.LoadFile(dir, c)
	if err != nil {
		return err
	}
	config = *c
	return nil
}

//GenDefaultConfigFile creates example config file at provided dir. Dir must include filename.
func GenDefaultConfigFile(dir string) error {
	c := createDefaultConfig()
	c.KnownIds = append(c.KnownIds, KnownId{ID: "6faae02a965b040d471be1cbcf2bebec53921029", Name: "Example ID"})
	return files.WriteJSONFile(dir, c)
}

func createDefaultConfig() Config {
	return Config{
		KnownIds: make([]KnownId, 0),
	}
}
