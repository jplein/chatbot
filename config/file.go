package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jplein/chatbot/storage"
)

// Default configuration file location, relative to the storage directory
const DefaultConfigPath = "config.json"

// Permissions used when initializing config file
const DefaultConfigFilePermissions = 0644

type Config struct {
	APIKey        string `json:"api_key"`
	LogTokenUsage bool   `json:"log_token_usage"`
}

type File struct {
	Path       string
	StorageDir storage.Dir
}

func (c *File) Init() error {
	var err error

	if c.Path == "" {
		if c.StorageDir.Path == "" {
			if err = c.StorageDir.Init(); err != nil {
				return err
			}
		}

		c.Path = filepath.Join(c.StorageDir.Path, DefaultConfigPath)
	}

	var stat os.FileInfo
	if stat, err = os.Stat(c.Path); err != nil {
		var buf []byte
		if buf, err = json.Marshal(Config{}); err != nil {
			return err
		}
		if err = os.WriteFile(c.Path, buf, DefaultConfigFilePermissions); err != nil {
			return err
		}
	} else {
		if stat.IsDir() {
			return fmt.Errorf("expected file %s to be a regular file but found a directory", c.Path)
		}
	}

	return nil
}

func (c *File) Read() (Config, error) {
	var err error
	var config Config

	var buf []byte
	if buf, err = os.ReadFile(c.Path); err != nil {
		return config, err
	}

	if err = json.Unmarshal(buf, &config); err != nil {
		return config, err
	}

	return config, nil
}

func (c *File) Write(config Config) error {
	var err error

	var buf []byte
	if buf, err = json.Marshal(config); err != nil {
		return err
	}
	if err = os.WriteFile(c.Path, buf, DefaultConfigFilePermissions); err != nil {
		return err
	}

	return nil
}
