package main

import (
	"log"
	"path/filepath"
	"strings"
)

func assertFileIsMarkdownFile(file string) {
	if strings.ToLower(filepath.Ext(file)) != ".md" {
		log.Fatalf("\"%s\" is not a markdown file.", file)
	}
}

func (ctx *Context) DeleteNote(source string) {
	sourceAbs := assertSourceExists(source)

	assertFileIsMarkdownFile(sourceAbs)
	//
	//attachmentFiles := oldGetAttachmentsFromNote(sourceAbs)
	//
	//var attachments = make(map[string]string)
	//
	//for _, attachment := range attachmentFiles {
	//	attachments[attachment] = ""
	//}

	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	err = ctx.Repository.PopulateFromVault(sourceObsidianRoot, true)

	if err != nil {
		log.Fatal(err)
	}
	//
	//// Resolve absolute path to the attachments
	//for _, file := range allFiles {
	//	_, found := attachments[path.Base(file)]
	//
	//	if found {
	//		attachments[path.Base(file)] = path.Dir(file)
	//	}
	//}
	//
	//var filesToKeep []string
	//
	//// Add this note to the pile to delete
	//attachments[path.Base(sourceAbs)] = path.Dir(sourceAbs)
	//
	//notes := getNotes(allFiles)
	//
	//// Check if attachments are not referenced and good to delete
	//for _, note := range notes {
	//	// Ignore this note which we are trying to delete, we care about the other notes referencing it
	//	if note == sourceAbs {
	//		continue
	//	}
	//
	//	data, err := os.ReadFile(path.Clean(note))
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	for attachmentName, attachmentPath := range attachments {
	//		if strings.Contains(string(data), attachmentName) {
	//			filesToKeep = append(filesToKeep, path.Join(attachmentPath, attachmentName))
	//		}
	//	}
	//}
	//
	//for attachmentName, attachmentPath := range attachments {
	//	// We couldn't find this file
	//	if len(attachmentPath) == 0 {
	//		continue
	//	}
	//
	//	filePath := path.Join(attachmentPath, attachmentName)
	//	if !utils.IsInArray(filePath, filesToKeep) {
	//		err = os.Remove(filePath)
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//}
}
