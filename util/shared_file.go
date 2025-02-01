package util

import (
	"fmt"
	"os"
)

func isDir(path string) bool {

	// Don't allow symbolic link files - treat as not dir
	info, err := os.Lstat(path)
	if err != nil || info.Mode()&os.ModeSymlink != 0 {
		return false
	}

	return info.IsDir()
}

func setupDir(directory string) (bool, error) {

	if !isDir(directory) {
		return false, fmt.Errorf("%s is not a directory", directory)
	}

	entries, err := os.ReadDir(directory)

	if err != nil {
		return false, err
	}

	for _, e := range entries {
		fmt.Printf("test: %v\n", e)
	}

	return false, nil
}
