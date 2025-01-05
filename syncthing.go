package main

import (
	"path"
	"strings"
)

func (ctx *Context) FindSyncConflicts(source string) error {
	sourceAbs := assertSourceExists(source)
	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		return err
	}

	err = ctx.Repository.PopulateFromVault(sourceObsidianRoot, false)

	if err != nil {
		return err
	}

	for filePath := range ctx.Repository.GetAllFiles() {
		if strings.Contains(path.Base(filePath), "sync-conflict-") {
			println(filePath)
		}
	}

	return nil
}
