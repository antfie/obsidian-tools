package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type ObsidianFile struct {
	Hash        string
	Attachments []string
}

func findObsidianRoot(path string) (string, error) {
	if !isDir(path) {
		path = filepath.Dir(path)
	}

	for {
		if info, err := os.Stat(filepath.Join(path, ".obsidian")); err == nil && info.IsDir() {
			return path, nil
		}

		lastPath := path
		path = filepath.Dir(path)

		// If the path has not changed then we have reached the filesystem root
		if lastPath == path {
			return "", os.ErrNotExist
		}
	}
}

func getAttachmentsFromNote(allFiles []string, notePath string) ([]string, []string) {
	var warnings []string
	noteData, err := os.ReadFile(path.Clean(notePath))

	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`!\[\[[^!\]]+]]`)

	// Attachments can be a filename or a (relative) path
	var attachments []string

	for _, attachment := range r.FindAllString(string(noteData), -1) {
		formattedAttachment := strings.TrimSuffix(strings.TrimPrefix(attachment, "![["), "]]")

		// With no file extension it could be a Markdown file
		if len(filepath.Ext(formattedAttachment)) == 0 {
			for _, file := range allFiles {
				if strings.HasSuffix(file, "/"+formattedAttachment+".md") {
					formattedAttachment = file
					break
				}
			}
		}

		attachments = append(attachments, formattedAttachment)
	}

	var resolvedAttachmentAbsolutePaths []string

	// Resolve the absolute paths to the attachments
	for _, file := range allFiles {
		for _, attachment := range attachments {
			if file == attachment || strings.HasSuffix(file, "/"+attachment) {
				resolvedAttachmentAbsolutePaths = append(resolvedAttachmentAbsolutePaths, file)
			}
		}
	}

	// Check to make sure we found all the attachments
	for _, attachment := range attachments {
		found := false

		for _, resolvedAttachment := range resolvedAttachmentAbsolutePaths {
			if resolvedAttachment == attachment || strings.HasSuffix(resolvedAttachment, "/"+attachment) {
				found = true
			}
		}

		if !found {
			warnings = append(warnings, fmt.Sprintf("Could not find attachment \"%s\" referenced by \"%s\".", attachment, notePath))
		}
	}

	return resolvedAttachmentAbsolutePaths, warnings
}
