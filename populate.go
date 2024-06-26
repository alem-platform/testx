package testx

import (
	"os"
	"path"
	"path/filepath"
)

type FileEntry struct {
	Path string
	Type FileType
}

type Cleanup func() error

// PopulateFS создает указанные файлы и директории.
func PopulateFS(workdir string, entries ...FileEntry) (_ Cleanup, err error) {
	paths := make([]string, 0, len(entries))

	defer func() {
		if err != nil {
			removeAll(paths, workdir) //nolint:errcheck //no need to check error
		}
	}()

	for _, entry := range entries {
		fullPath := path.Join(workdir, entry.Path)
		if err := entry.create(fullPath); err != nil {
			return nil, err
		}
		paths = append(paths, entry.Path)
	}

	return func() error {
		return removeAll(paths, workdir)
	}, nil
}

func (e *FileEntry) create(fullPath string) error {
	switch e.Type {
	case TypeDir:
		if err := os.MkdirAll(fullPath, 0o755); err != nil {
			return err
		}
	case TypeFile:
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			return err
		}
		file, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func removeAll(paths []string, workdir string) error {
	for _, p := range paths {
		for p != "." {
			if err := os.RemoveAll(path.Join(workdir, p)); err != nil {
				return err
			}
			p = filepath.Dir(p)
		}
	}
	return nil
}
