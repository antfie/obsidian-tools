package main

import (
	_ "embed"
	"fmt"
	"github.com/antfie/obsidian-tools/utils"
	"log"
	"os"
	"strings"
)

//goland:noinspection GoUnnecessarilyExportedIdentifiers
var AppVersion = "0.0"

var usageText = "Usage: go run main.go command.\nAvailable commands:\n  move\n  copy\n  delete\n  find_missing_attachments\n  find_duplicates\n  find_empty_files\n  find_sync_conflicts\n"

//go:embed config.yaml
var defaultConfigData []byte

func main() {
	print(fmt.Sprintf("obsidian-tools version %s\n", AppVersion))

	c := Load(defaultConfigData)

	err := utils.SetupLogger(c.LogFilePath)

	if err != nil {
		log.Fatal(err)
	}

	ctx := &Context{
		Config:     c,
		Repository: NewRepository(c),
	}

	if len(os.Args) < 2 {
		log.Fatal("No command specified. " + usageText)
	}

	command := os.Args[1]

	switch strings.ToLower(command) {
	case "move":
		if len(os.Args) != 4 {
			log.Fatal("Move requires source and destination.")
		}

		err = ctx.CopyNotes(os.Args[2], os.Args[3], true)

		if err != nil {
			log.Fatal(err)
		}

		return

	case "copy":
		if len(os.Args) != 4 {
			log.Fatal("Copy requires source and destination.")
		}

		err = ctx.CopyNotes(os.Args[2], os.Args[3], false)

		if err != nil {
			log.Fatal(err)
		}

		return

	case "delete":
		if len(os.Args) != 3 {
			log.Fatal("Delete requires a source")
		}

		ctx.DeleteNote(os.Args[2])
		return

	case "find_missing_attachments":
		if len(os.Args) != 3 {
			log.Fatal("Find missing attachments requires a source.")
		}

		ctx.FindMissingAttachments(os.Args[2])
		return

	case "find_duplicates":
		if len(os.Args) != 3 {
			log.Fatal("Find duplicates requires a source.")
		}

		ctx.FindDuplicates(os.Args[2])
		return

	case "find_empty_files":
		if len(os.Args) != 3 {
			log.Fatal("Find sync conflicts requires a source.")
		}

		ctx.FindEmptyFiles(os.Args[2])
		return

	case "find_sync_conflicts":
		if len(os.Args) != 3 {
			log.Fatal("Find sync conflicts requires a source.")
		}

		ctx.FindSyncConflicts(os.Args[2])
		return
	}

	log.Fatalf("Command \"%s\" not recognised. %s", command, usageText)
}
