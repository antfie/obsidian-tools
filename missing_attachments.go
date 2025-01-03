package main

import (
	"log"
)

func (ctx *Context) FindMissingAttachments(source string) {
	sourceAbs := assertSourceExists(source)

	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	// This will output any warnings about missing attachments
	err = ctx.Repository.PopulateFromVault(sourceObsidianRoot, true)

	if err != nil {
		log.Fatal(err)
	}
}
