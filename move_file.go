package main

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func copyNote(source, destination string, deleteSource bool) {
	sourceAbs := assertFolderExists(source)

	if strings.ToLower(filepath.Ext(sourceAbs)) != ".md" {
		log.Fatalf("Source \"%s\" file is not a markdown file.", sourceAbs)
	}

	destinationAbs, err := filepath.Abs(destination)

	if err != nil {
		log.Fatal(err)
	}

	if sourceAbs == destinationAbs {
		log.Fatalf("Source and destination are the same.")
	}

	destinationStat, err := os.Stat(destinationAbs)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("Creating desitnation path \"%s\",,.", destinationAbs)
			err := os.Mkdir(destinationAbs, 0750)

			if err != nil {
				log.Fatalf("Could not create destination directory \"%s\".", destinationAbs)
			}

			destinationStat, err = os.Stat(destinationAbs)
		} else {
			log.Fatal(err)
			return
		}
	}

	if destinationStat.IsDir() {
		destinationAbs = path.Join(destinationAbs, path.Base(sourceAbs))
	}

	attachmentFiles := getAttachmentsFromNote(sourceAbs)

	var attachments = make(map[string]string)

	for _, attachment := range attachmentFiles {
		attachments[attachment] = ""
	}

	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	allFiles := getAllFiles(sourceObsidianRoot)

	for _, file := range allFiles {
		_, found := attachments[path.Base(file)]

		if found {
			attachments[path.Base(file)] = path.Dir(file)
		}
	}

	// Sanity check to ensure all attachments could be found
	for attachmentFileName, attachmentPath := range attachments {
		if len(attachmentPath) == 0 {
			log.Fatalf("Could not find attachment \"%s\".", attachmentFileName)
		}
	}

	copyAttachments(destinationAbs, attachments)

	err = copyFile(sourceAbs, destinationAbs)

	if err != nil {
		log.Fatal(err)
	}

	if !deleteSource {
		return
	}

	var sourceFileStillInUse = false
	var filesToKeep []string

	notes := getNotes(allFiles)

	// Check if attachments are not referenced and good to delete
	for _, note := range notes {
		data, err := os.ReadFile(path.Clean(note))

		if err != nil {
			log.Fatal(err)
		}

		for attachmentName, _ := range attachments {
			if strings.Contains(string(data), attachmentName) {
				filesToKeep = append(filesToKeep, attachmentName)
			}
		}

		if strings.Contains(string(data), filepath.Base(sourceAbs)) {
			sourceFileStillInUse = true
		}
	}

	// Delete the source file if no longer referenced
	if !sourceFileStillInUse {
		//err = os.Remove(sourceAbs)

		if err != nil {
			log.Fatal(err)
		}
	}

	for attachmentName, attachmentPath := range attachments {
		if !isInArray(attachmentName, filesToKeep) {
			//err = os.Remove(path.Join(attachmentPath,attachmentName ))
			print(attachmentName, attachmentPath)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func copyAttachments(destinationAbs string, attachments map[string]string) {
	destinationObsidianRoot, err := findObsidianRoot(destinationAbs)

	if err != nil {
		log.Fatal(err)
	}

	destinationAttachmentPath := getAttachmentPath(destinationObsidianRoot)

	for fileName, filePath := range attachments {
		sourcePath := path.Join(filePath, fileName)
		destinationPath := path.Join(destinationAttachmentPath, fileName)

		err = copyFile(sourcePath, destinationPath)

		if err != nil {
			log.Fatal(err)
		}
	}
}