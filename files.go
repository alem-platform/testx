package testx

import (
	"bytes"
	"io/fs"
	"os"
	"sort"
)

type File struct {
	// Path is a relative path to a file.
	Path string

	// Type is for FILE or LINK.
	Type FileType

	Permission fs.FileMode

	// Link is a symlink.
	Link string

	Content []byte
}

type FileType string

const (
	TypeFile FileType = "FILE"
	TypeLink FileType = "LINK"
	TypeDir  FileType = "DIR"
)

type Flag struct {
	IgnorePermission bool
	IgnoreContent    bool
}

// CompareFiles compares whether exp files equal to files containing in path.
func CompareFiles(exp []File, path string, flag Flag) (equal bool, actual []File, err error) {
	act, err := GetFiles(path)
	if err != nil {
		return false, nil, err
	}

	return EqualFiles(exp, act, flag), act, nil
}

// GetFiles returns a list of files in given path.
// Each file entry will contain .Path field equal to relative path.
func GetFiles(path string) ([]File, error) {
	if err := os.Chdir(path); err != nil {
		return nil, err
	}

	files := []File{}
	err := fs.WalkDir(os.DirFS(path), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		var (
			isFile = info.Mode().IsRegular()
			isLink = info.Mode().Type()&fs.ModeSymlink > 0
		)

		// Determine file type.
		var fileType FileType
		switch {
		case isFile:
			fileType = TypeFile
		case isLink:
			fileType = TypeLink
		}

		// Set link if it is symlink.
		var link string
		if isLink {
			tmpLink, err := os.Readlink(path)
			if err != nil {
				return err
			}
			link = tmpLink
		}

		// Set content if it is a non-empty file.
		var content []byte
		if isFile && info.Size() > 0 {
			bts, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			content = bts
		}

		files = append(files, File{
			Path:       path,
			Type:       fileType,
			Permission: info.Mode().Perm(),
			Link:       link,
			Content:    content,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func EqualFiles(files1, files2 []File, flag Flag) bool {
	if len(files1) != len(files2) {
		return false
	}

	sort.SliceStable(files1, func(i, j int) bool {
		return files1[i].Path < files1[j].Path
	})

	sort.SliceStable(files2, func(i, j int) bool {
		return files2[i].Path < files2[j].Path
	})

	for i := 0; i < len(files1); i++ {
		if !EqualFile(files1[i], files2[i], flag) {
			return false
		}
	}

	return true
}

func EqualFile(a, b File, flag Flag) bool {
	if a.Path != b.Path || a.Type != b.Type || a.Link != b.Link {
		return false
	}

	if !flag.IgnorePermission && a.Permission != b.Permission {
		return false
	}

	if !flag.IgnoreContent && !bytes.Equal(a.Content, b.Content) {
		return false
	}

	return true
}
