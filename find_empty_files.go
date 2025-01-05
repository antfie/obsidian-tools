package main

import (
	"os"
)

func (ctx *Context) FindEmptyFiles(source string) error {
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
		stat, err := os.Stat(filePath)

		if err != nil {
			return err
		}

		if stat.Size() == 0 {
			println(filePath)
		}
	}

	return nil
}
