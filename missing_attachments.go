package main

import (
	"log"
	"path"
	"strings"
)

func findMissingAttachments(source string) {
	sourceAbs := assertSourceExists(source)

	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	allFiles := getAllFiles(sourceObsidianRoot)

	for _, note := range getNotes(allFiles) {
		for _, attachment := range getAttachmentsFromNote(note) {
			if !isAttachmentInFiles(attachment, allFiles) {
				log.Printf("Note \"%s\" is missing \"%s\"", note, attachment)
			}
		}
	}
}

func isAttachmentInFiles(attachment string, files []string) bool {
	if strings.Contains(attachment, "|") && len(strings.Split(attachment, "|")) == 2 {
		// Try both variants
		if !isAttachmentInFiles(strings.TrimSpace(strings.Split(attachment, "|")[0]), files) {
			return isAttachmentInFiles(strings.TrimSpace(strings.Split(attachment, "|")[1]), files)
		}
	}

	if len(path.Ext(attachment)) == 0 {
		attachment = attachment + ".md"
	}

	for _, file := range files {
		if strings.HasSuffix(file, attachment) {
			return true
		}
	}

	return false
}
