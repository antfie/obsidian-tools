package main

import (
	"log"
	"path"
	"strings"
)

func (ctx *Context) FindSyncConflicts(source string) {
	sourceAbs := assertSourceExists(source)
	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	err = ctx.Repository.PopulateFromVault(sourceObsidianRoot, false)

	if err != nil {
		log.Fatal(err)
	}

	for filePath := range ctx.Repository.GetAllFiles() {
		if strings.Contains(path.Base(filePath), "sync-conflict-") {
			println(filePath)
		}
	}
}
