package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/antfie/obsidian-tools/utils"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

//goland:noinspection GoUnnecessarilyExportedIdentifiers
var AppVersion = "0.0"

var usageText = "Usage: ./obsidian-tools command.\nAvailable commands:\n  move\n  copy\n  delete\n  find_missing_attachments\n  find_duplicates\n  find_empty_files\n  find_sync_conflicts\n"

//go:embed config.yaml
var defaultConfigData []byte

func main() {

	c := Load(defaultConfigData)

	err := utils.SetupLogger(c.LogFilePath)

	if err != nil {
		log.Fatal(err)
	}

	ctx := &Context{
		Config:     c,
		Repository: NewRepository(c),
	}

	dryRunFormat := ""

	if c.DryRun {
		dryRunFormat = " (dry run)"
	}

	utils.ConsoleAndLogPrintf("Data Tools version %s%s. Using %s for file operations", AppVersion, dryRunFormat, utils.Pluralize("thread", ctx.Config.MaxConcurrentFileOperations))

	if len(os.Args) < 2 {
		log.Fatal("No command specified. " + usageText)
	}

	startTime := time.Now()

	err = ctx.runCommand(strings.ToLower(os.Args[1]))

	if err != nil {
		utils.ConsoleAndLogPrintf("%v", err)
	}

	duration := math.Round(time.Since(startTime).Seconds())
	formattedDuration := fmt.Sprintf("%.0f second", duration)

	if duration != 1 {
		formattedDuration += "s"
	}

	utils.ConsoleAndLogPrintf("Finished in %s", formattedDuration)
}

func (ctx *Context) runCommand(command string) error {
	switch command {
	case "move":
		if len(os.Args) != 4 {
			log.Fatal("Move requires source and destination.")
		}

		return ctx.CopyNotes(os.Args[2], os.Args[3], true)

	case "copy":
		if len(os.Args) != 4 {
			log.Fatal("Copy requires source and destination.")
		}

		return ctx.CopyNotes(os.Args[2], os.Args[3], false)

	case "delete":
		if len(os.Args) != 3 {
			log.Fatal("Delete requires a source")
		}

		return ctx.DeleteNote(os.Args[2])

	case "find_missing_attachments":
		if len(os.Args) != 3 {
			log.Fatal("Find missing attachments requires a source.")
		}

		return ctx.FindMissingAttachments(os.Args[2])

	case "find_duplicates":
		if len(os.Args) != 3 {
			log.Fatal("Find duplicates requires a source.")
		}

		return ctx.FindDuplicates(os.Args[2])

	case "find_empty_files":
		if len(os.Args) != 3 {
			log.Fatal("Find sync conflicts requires a source.")
		}

		return ctx.FindEmptyFiles(os.Args[2])

	case "find_sync_conflicts":
		if len(os.Args) != 3 {
			log.Fatal("Find sync conflicts requires a source.")
		}

		return ctx.FindSyncConflicts(os.Args[2])

	}

	return errors.New(fmt.Sprintf("Command \"%s\" not recognised. %s", command, usageText))
}
