package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

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

		// If the path has not changed we have reached the root
		if lastPath == path {
			return "", os.ErrNotExist
		}
	}
}

type ObsidianConfig struct {
	AttachmentFolderPath string `json:"attachmentFolderPath"`
}

func getAttachmentPath(obsidianRoot string) string {
	obsidianConfigFile := filepath.Join(obsidianRoot, ".obsidian", "app.json")

	data, err := os.ReadFile(path.Clean(obsidianConfigFile))

	if err != nil {
		log.Fatal(err)
	}

	config := ObsidianConfig{}
	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatal(err)
	}

	return path.Join(obsidianRoot, config.AttachmentFolderPath)
}

func getNotes(files []string) []string {
	var notes []string

	for _, file := range files {
		if strings.ToLower(path.Ext(file)) == ".md" {
			notes = append(notes, file)
		}

	}

	return notes
}

func getAttachmentsFromNote(notePath string) []string {
	sourceFileData, err := os.ReadFile(path.Clean(notePath))

	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`!\[\[.+]]`)
	var attachments []string

	for _, attachment := range r.FindAllString(string(sourceFileData), -1) {
		attachment = strings.TrimSuffix(strings.TrimPrefix(attachment, "![["), "]]")
		attachments = append(attachments, attachment)
	}

	return attachments
}
