package main

import (
	"log"
	"os"
)

func (ctx *Context) FindEmptyFiles(source string) {
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
		stat, err := os.Stat(filePath)

		if err != nil {
			log.Fatal(err)
		}

		if stat.Size() == 0 {
			println(filePath)
		}
	}
}
