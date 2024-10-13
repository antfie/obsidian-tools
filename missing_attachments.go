package main

import (
	"log"
)

func findMissingAttachments(source string) {
	sourceAbs := assertFolderExists(source)

	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	allFiles := getAllFiles(sourceObsidianRoot)

	var referencedAttachments []string

	for _, note := range getNotes(allFiles) {
		attachments := getAttachmentsFromNote(note)
		referencedAttachments = append(referencedAttachments, attachments...)
	}

	attachmentPath := getAttachmentPath(sourceObsidianRoot)
	allAttachments := getAllFiles(attachmentPath)

	print(allAttachments)
	print(referencedAttachments)
}
