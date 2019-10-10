package osutil

import (
	"os"
	"path/filepath"
)

// Abs returns absolute file path inclusing resolving Symbolic links.
func Abs(path string) (string, error) {
	var abs string
	var err error

	if abs, err = filepath.Abs(filepath.Clean(path)); err != nil {
		return "", err
	}

	if abs, err = filepath.EvalSymlinks(abs); err != nil {
		return "", err
	}

	return abs, err
}

// Dir returns the full directory for path
// including symbolic link resolving.
func Dir(path string) (string, error) {
	var abs string
	var err error

	if abs, err = Abs(path); err != nil {
		return "", err
	}

	return filepath.Dir(abs), err
}

// IsDirectory will check if the provided path points to a
// directory or a file
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

// Exists returns true if the provided path exists.
func Exists(path string) (bool, error) {
	fullpath, err := Abs(path)
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

// List returns a list of all files within the provided path
func List(path string) ([]string, error) {
	files := make([]string, 0)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// EOF
