package testx

import (
	"fmt"
	"os"
	"path"
)

type FileEntry struct {
	Path string
	Type FileType
}

type cleanup func() error

// PopulateFS создает указанные файлы и директории.
func PopulateFS(workdir string, entries ...FileEntry) (cleanup, error) {
	cleanups := make([]cleanup, 0)

	for _, entry := range entries {
		fullPath := path.Join(workdir, entry.Path)
		if err := entry.create(fullPath, cleanups); err != nil {
			return nil, err
		}
		cleanups = append(cleanups, func() error {
			return os.RemoveAll(fullPath)
		})
	}

	return func() error {
		return cleanupAll(cleanups)
	}, nil
}

func (e *FileEntry) create(fullPath string, cleanups []cleanup) error {
	if e.Type == TypeDir {
		if err := os.MkdirAll(fullPath, 0o755); err != nil {
			if cleanupErr := cleanupAll(cleanups); cleanupErr != nil {
				return fmt.Errorf("error cleaning up: %v, original error: %v", cleanupErr, err)
			}
			return err
		}
	} else {
		file, err := os.Create(fullPath)
		if err != nil {
			if cleanupErr := cleanupAll(cleanups); cleanupErr != nil {
				return fmt.Errorf("error cleaning up: %v, original error: %v", cleanupErr, err)
			}
			return err
		}
		file.Close()
	}
	return nil
}

func cleanupAll(cleanups []cleanup) error {
	for _, cleanupFunc := range cleanups {
		if cleanupErr := cleanupFunc(); cleanupErr != nil {
			return cleanupErr
		}
	}
	return nil
}
