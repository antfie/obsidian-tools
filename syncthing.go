package main

import (
	"log"
	"path"
	"strings"
)

func findSyncConflicts(source string) {
	sourceAbs := assertFolderExists(source)
	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	allFiles := getAllFiles(sourceObsidianRoot)

	for _, file := range allFiles {
		if strings.Contains(path.Base(file), "sync-conflict-") {
			println(file)
		}
	}
}
