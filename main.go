package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//goland:noinspection GoUnnecessarilyExportedIdentifiers
var AppVersion = "0.0"

func main() {
	print(fmt.Sprintf("obsidian-tools version %s\n", AppVersion))

	if len(os.Args) < 2 {
		log.Fatal("No command specified. Usage: go run main.go command.\nAvailable commands: move, copy, find_missing_attachments, find_duplicates, find_sync_conflicts, find_empty_files")
	}

	command := os.Args[1]

	switch strings.ToLower(command) {
	case "move":
		if len(os.Args) != 4 {
			log.Fatal("Move requires source and destination.")
		}

		copyNote(os.Args[2], os.Args[3], true)
		return

	case "copy":
		if len(os.Args) != 4 {
			log.Fatal("Copy requires source and destination.")
		}

		copyNote(os.Args[2], os.Args[3], false)
		return

	case "delete":
		if len(os.Args) != 3 {
			log.Fatal("Delete requires a source")
		}

		deleteNote(os.Args[2])
		return

	case "find_missing_attachments":
		if len(os.Args) != 3 {
			log.Fatal("Find missing attachments requires a source.")
		}

		findMissingAttachments(os.Args[2])
		return

	case "find_duplicates":
		if len(os.Args) != 3 {
			log.Fatal("Find duplicates requires a source.")
		}

		findDuplicates(os.Args[2])
		return

	case "find_sync_conflicts":
		if len(os.Args) != 3 {
			log.Fatal("Find sync conflicts requires a source.")
		}

		findSyncConflicts(os.Args[2])
		return

	case "find_empty_files":
		if len(os.Args) != 3 {
			log.Fatal("Find sync conflicts requires a source.")
		}

		findEmptyFiles(os.Args[2])
		return
	}

	log.Fatalf("Command \"%s\" not recognised.", command)
}
