package main

import (
	"errors"
	"github.com/antfie/obsidian-tools/utils"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetAllFiles(rootPath string, folderNamesToIgnore, fileNamesToIgnore []string) ([]string, error) {
	var files []string

	err := filepath.Walk(rootPath, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileName := path.Base(currentPath)

		if info.IsDir() {
			if utils.IsInArray(fileName, folderNamesToIgnore) {
				return filepath.SkipDir
			}
		} else {
			if !utils.IsInArray(fileName, fileNamesToIgnore) {
				absoluteFileName, err := filepath.Abs(currentPath)

				if err != nil {
					return err
				}

				files = append(files, absoluteFileName)
			}
		}

		return nil
	})

	return files, err
}

func assertSourceExists(source string) string {
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

func isMarkdownFile(filePath string) bool {
	return strings.ToLower(filepath.Ext(filePath)) == ".md"
}

func IsFile(path string) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return !info.IsDir(), nil
}
