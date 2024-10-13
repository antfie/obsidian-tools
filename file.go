package main

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
)

func getAllFiles(path string) []string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
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
