package testx

import (
	"fmt"
	"math/rand"
	"os"
	"path"
)

// RandMatrix generates a random slice of slices,
// where n is the size of the outer and m is the size of the inner slices.
func RandMatrix[T any](n, m int, fn func() T) [][]T {
	result := make([][]T, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, RandSlice(m, fn))
	}
	return result
}

// RandSlice generates a random slice,
// where n is the size of the slice.
func RandSlice[T any](n int, fn func() T) []T {
	result := make([]T, 0, n)

	for i := 0; i < n; i++ {
		result = append(result, fn())
	}
	return result
}

func RandInt() int {
	return rand.Intn(300)
}

func RandIntPtr() *int {
	return getPtr(RandInt())
}

func RandRune() rune {
	return rune(rand.Intn(256))
}

func getPtr[T any](v T) *T {
	return &v
}

func RandValidStdin() string {
	return fmt.Sprintf("%d %d", rand.Intn(100), rand.Intn(100))
}

func RandInvalidStdin() string {
	return fmt.Sprintf("%d %c", rand.Intn(100), rune(rand.Intn(100)))
}

func RandDirs(workDir string, dirs int) error {
	for i := 0; i < dirs; i++ {
		dirName := fmt.Sprintf("dir%d", i)
		dirPath := path.Join(workDir, dirName)
		if err := os.Mkdir(dirPath, 0o755); err != nil {
			return err
		}

		files := rand.Intn(5) // Max 5 files per directory
		for j := 0; j < files; j++ {
			fileName := fmt.Sprintf("file%d.txt", j)
			filePath := path.Join(dirPath, fileName)
			if _, err := os.Create(filePath); err != nil {
				return err
			}
		}

		subDirs := rand.Intn(2) // Max 2 subdirectories per directory
		if subDirs > 0 {
			if err := RandDirs(dirPath, subDirs); err != nil {
				return err
			}
		}
	}
	return nil
}

func CleanDirs(workDir string) error {
	files, err := os.ReadDir(workDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := path.Join(workDir, file.Name())
		if file.IsDir() {
			if err := os.RemoveAll(filePath); err != nil {
				return err
			}
		} else {
			if err := os.Remove(filePath); err != nil {
				return err
			}
		}
	}

	return nil
}
