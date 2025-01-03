package main

import (
	"github.com/antfie/obsidian-tools/utils"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CopyOrMoveFiles(source, destination string, folderNamesToIgnore, fileNamesToIgnore []string, move bool) error {
	files, err := GetAllFiles(source, folderNamesToIgnore, fileNamesToIgnore)

	if err != nil {
		return err
	}

	for _, sourceFilePath := range files {
		relativePath := strings.TrimPrefix(sourceFilePath, source)
		err := CopyOrMoveFile(sourceFilePath, path.Join(destination, relativePath), move)

		if err != nil {
			return err
		}
	}

	return nil
}

func CopyOrMoveFile(sourceFilePath, destinationFilePath string, move bool) error {
	_, err := os.Stat(destinationFilePath)

	// Was the error anything other than an expected file not found?
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// A file with this name exists at this location
	if err == nil {
		// Is it the same though?
		filesAreTheSame, err := CompareFiles(sourceFilePath, destinationFilePath)

		if err != nil {
			return err
		}

		if filesAreTheSame {
			// Remove existing file if moving
			if move {
				log.Printf("Deleting \"%s\" which has same contents as \"%s\"", sourceFilePath, destinationFilePath)
				return os.Remove(sourceFilePath)
			}

			// Otherwise there is nothing to do
			return nil
		}

		// The files are different but with the same name so we need to rename this destination path to avoid a collision
		newDestinationPath := path.Join(strings.TrimSuffix(destinationFilePath, path.Base(destinationFilePath)), utils.NewID(), path.Base(destinationFilePath))
		log.Printf("Existing but different file exists \"%s\", creating new file at \"%s\" ", destinationFilePath, newDestinationPath)
		destinationFilePath = newDestinationPath

		// Sanity check!
		_, err = os.Stat(destinationFilePath)

		// This path should not be found
		if err == nil || !os.IsNotExist(err) {
			return ErrCouldNotAssertDestinationFilePathUnique
		}
	}

	destinationDirectory := filepath.Dir(destinationFilePath)

	_, err = os.Stat(destinationDirectory)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		// Create the directory structure if required
		err = os.MkdirAll(filepath.Dir(destinationFilePath), 0750)

		if err != nil {
			return err
		}
	}

	if move {
		return os.Rename(sourceFilePath, destinationFilePath)
	}

	return copyFile(sourceFilePath, destinationFilePath)
}

func copyFile(sourceFilePath, destinationFilePath string) error {
	sourceFile, err := os.Open(path.Clean(sourceFilePath))

	if err != nil {
		return err
	}

	destinationFile, err := os.Create(path.Clean(destinationFilePath))

	_, err = io.Copy(destinationFile, sourceFile)

	if err != nil {
		return err
	}

	err = destinationFile.Sync()

	if err != nil {
		return err
	}

	err = destinationFile.Close()

	if err != nil {
		return err
	}

	return sourceFile.Close()
}
