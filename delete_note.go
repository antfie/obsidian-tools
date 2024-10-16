package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func assertSourceIsMarkdownFile(source string) {
	if strings.ToLower(filepath.Ext(source)) != ".md" {
		log.Fatalf("Source \"%s\" file is not a markdown file.", source)
	}
}

func deleteNote(source string) {
	sourceAbs := assertSourceExists(source)

	assertSourceIsMarkdownFile(sourceAbs)

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

	// Resolve absolute path to the attachments
	for _, file := range allFiles {
		_, found := attachments[path.Base(file)]

		if found {
			attachments[path.Base(file)] = path.Dir(file)
		}
	}

	var filesToKeep []string

	// Add this note to the pile to delete
	attachments[path.Base(sourceAbs)] = path.Dir(sourceAbs)

	notes := getNotes(allFiles)

	// Check if attachments are not referenced and good to delete
	for _, note := range notes {
		// Ignore this note which we are trying to delete, we care about the other notes referencing it
		if note == sourceAbs {
			continue
		}

		data, err := os.ReadFile(path.Clean(note))

		if err != nil {
			log.Fatal(err)
		}

		for attachmentName, attachmentPath := range attachments {
			if strings.Contains(string(data), attachmentName) {
				filesToKeep = append(filesToKeep, path.Join(attachmentPath, attachmentName))
			}
		}
	}

	for attachmentName, attachmentPath := range attachments {
		// We couldn't find this file
		if len(attachmentPath) == 0 {
			continue
		}

		filePath := path.Join(attachmentPath, attachmentName)
		if !isInArray(filePath, filesToKeep) {
			err = os.Remove(filePath)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
