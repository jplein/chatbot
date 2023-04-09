package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// Default path, relative to the user's home directory
const DefaultStorageDirectory = ".chatbot"

// Default permissions for storage directory
const DefaultStorageDirectoryPermissions = 0755

type Dir struct {
	Path string
}

func (d *Dir) Init() error {
	var err error

	if d.Path == "" {
		var home string
		if home, err = os.UserHomeDir(); err != nil {
			return err
		}

		d.Path = filepath.Join(
			home,
			DefaultStorageDirectory,
		)
	}

	var stats os.FileInfo
	if stats, err = os.Stat(d.Path); err != nil {
		if err = os.Mkdir(d.Path, DefaultStorageDirectoryPermissions); err != nil {
			return err
		}
	} else if !stats.IsDir() {
		return fmt.Errorf("expected %s to be a directory but found a plain file", d.Path)
	}

	return nil
}
