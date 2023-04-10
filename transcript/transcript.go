package transcript

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/jplein/chatbot/storage"
)

// Path relative to the root configuration directory where transcripts are
// recorded
const TranscriptsPath = "transcripts"

type Role string

const (
	System    Role = "system"
	User      Role = "user"
	Assistant Role = "assistant"
)

type Record struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

func getTranscriptsDir(dir *storage.Dir, id int) string {
	return filepath.Join(
		TranscriptsPath,
		fmt.Sprintf("%d", id),
	)
}

func Read(dir *storage.Dir, id int) ([]Record, error) {
	var err error

	records := make([]Record, 0)

	storageDir := dir.Path
	if storageDir == "" {
		return nil, fmt.Errorf("No storage directory set, call Init() on dir before passing it to Send()")
	}

	transcriptsDir := getTranscriptsDir(dir, id)

	var stats os.FileInfo
	if stats, err = os.Stat(transcriptsDir); err != nil {
		// If the directory doesn't exist, return an empty list
		return records, nil
	}
	if !stats.IsDir() {
		return nil, fmt.Errorf("path %s exists but is not a directory", transcriptsDir)
	}

	var files []os.DirEntry
	if files, err = os.ReadDir(transcriptsDir); err != nil {
		return nil, err
	}

	var fileNames []string
	for _, f := range files {
		fileNames = append(
			fileNames,
			f.Name(),
		)
	}

	sort.Strings(fileNames)

	for _, name := range fileNames {
		var buf []byte
		if buf, err = os.ReadFile(filepath.Join(transcriptsDir, name)); err != nil {
			return nil, err
		}

		var record Record
		if err = json.Unmarshal(buf, &record); err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

func Write(dir *storage.Dir, id int, r Record) error {
	var err error

	t := time.Now().UnixMicro()
	transcriptsDir := getTranscriptsDir(dir, id)

	if err = os.MkdirAll(transcriptsDir, 0755); err != nil {
		return err
	}

	f := filepath.Join(transcriptsDir, fmt.Sprintf("%d.json", t))

	var buf []byte
	if buf, err = json.Marshal(r); err != nil {
		return err
	}

	if err = os.WriteFile(f, buf, 0644); err != nil {
		return err
	}

	return nil
}
