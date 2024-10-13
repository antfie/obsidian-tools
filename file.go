package main

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
)

var ignoredFileNames = []string{
	".DS_Store",
}

var ignoredFolderNames = []string{
	".stversions",
}

func getAllFiles(rootPath string) []string {
	var files []string

	err := filepath.Walk(rootPath, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base := path.Base(currentPath)

		if info.IsDir() {
			if isInArray(base, ignoredFolderNames) {
				return filepath.SkipDir
			}
		} else {
			if !isInArray(base, ignoredFileNames) {
				files = append(files, currentPath)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func copyFile(source, destination string) error {
	data, err := os.ReadFile(path.Clean(source))

	if err != nil {
		return err
	}

	err = os.WriteFile(destination, data, 0600)

	if err != nil {
		return err
	}

	return nil
}

func assertFolderExists(source string) string {
	sourceAbs, err := filepath.Abs(source)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(sourceAbs); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Source \"%s\" not found.", sourceAbs)
	}

	return sourceAbs
}

func isDir(path string) bool {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return true
	}

	return false
}
