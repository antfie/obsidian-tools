package main

import (
	"github.com/antfie/obsidian-tools/utils"
	"github.com/schollz/progressbar/v3"
	"log"
	"path"
	"path/filepath"
	"strings"
)

// CopyNotes will copy or move a note or tree of notes and their attachments.
func (ctx *Context) CopyNotes(source, destination string, move bool) error {
	sourceObsidianRoot, destinationObsidianRoot, err := getSourceAndDestinationObsidianRoots(source, destination)

	if err != nil {
		return err
	}

	absoluteSourcePath, err := filepath.Abs(source)

	if err != nil {
		return err
	}

	absoluteDestinationPath, err := filepath.Abs(destination)

	if err != nil {
		return err
	}

	destinationObsidianConfig, err := ParseObsidianConfig(destinationObsidianRoot)

	if err != nil {
		return err
	}

	isFile, err := IsFile(absoluteSourcePath)

	if err != nil {
		return err
	}

	err = ctx.Repository.PopulateFromVault(sourceObsidianRoot, true)

	if err != nil {
		return err
	}

	copyingOrMovingText := "Copying"

	if move {
		copyingOrMovingText = "Moving"
	}

	// Is source a single file?
	if isFile {
		if !isMarkdownFile(absoluteSourcePath) {
			return ErrNotMarkdownFile
		}

		log.Printf("%s 1 note from \"%s\" to \"%s\"", copyingOrMovingText, absoluteSourcePath, absoluteDestinationPath)

		// Here it is intentional we pass absoluteSourcePath twice. The first is used for filtering attachments when copying folders
		return ctx.copyNote(sourceObsidianRoot, absoluteSourcePath, absoluteSourcePath, destinationObsidianConfig, absoluteDestinationPath, move)
	}

	sourceFilesToCopy := ctx.Repository.GetAllFilePathsInFolder(absoluteSourcePath)

	log.Printf("%s %s from \"%s\" to \"%s\"", copyingOrMovingText, utils.Pluralize("note", len(sourceFilesToCopy)), absoluteSourcePath, absoluteDestinationPath)

	bar := progressbar.Default(int64(len(sourceFilesToCopy)))

	// First process the regular, non-markdown files
	for _, sourceFilePath := range sourceFilesToCopy {
		if isMarkdownFile(sourceFilePath) {
			continue
		}

		relativeSourcePath := strings.TrimPrefix(sourceFilePath, absoluteSourcePath)
		err := ctx.Repository.CopyOrMoveFile(sourceFilePath, path.Join(absoluteDestinationPath, relativeSourcePath), move)

		if err != nil {
			return err
		}

		err = bar.Add(1)

		if err != nil {
			return err
		}
	}

	// Next process the markdown files
	for _, sourceFilePath := range sourceFilesToCopy {
		if !isMarkdownFile(sourceFilePath) {
			continue
		}

		err := ctx.copyNote(sourceObsidianRoot, absoluteSourcePath, sourceFilePath, destinationObsidianConfig, absoluteDestinationPath, move)

		if err != nil {
			return err
		}

		err = bar.Add(1)

		if err != nil {
			return err
		}
	}

	return nil
}

func getSourceAndDestinationObsidianRoots(source, destination string) (string, string, error) {
	sourceObsidianRoot, err := findObsidianRoot(source)

	if err != nil {
		return "", "", ErrSourceVaultNotResolved
	}

	destinationObsidianRoot, err := findObsidianRoot(destination)

	if err != nil {
		return "", "", ErrDestinationVaultNotResolved
	}

	// File operations within the same vault should be done using Obsidian
	if sourceObsidianRoot == destinationObsidianRoot {
		return "", "", ErrDestinationIsSameVaultAsSource
	}

	return sourceObsidianRoot, destinationObsidianRoot, nil
}

func (ctx *Context) copyNote(sourceObsidianRoot, absoluteSourcePath, sourceNotePath string, destinationObsidianConfig *ObsidianConfig, finalDestinationNotePath string, move bool) error {
	err := ctx.CopyAttachments(absoluteSourcePath, sourceNotePath, destinationObsidianConfig, move)

	if err != nil {
		return err
	}

	relativeSourcePath := strings.TrimPrefix(sourceNotePath, sourceObsidianRoot)
	relativeDestinationPath := strings.TrimPrefix(finalDestinationNotePath, destinationObsidianConfig.VaultRoot)

	// Use the default location in the vault for new files if no specific vault folder was specified
	if finalDestinationNotePath == destinationObsidianConfig.VaultRoot {
		finalDestinationNotePath = path.Join(destinationObsidianConfig.VaultRoot, destinationObsidianConfig.NewFileFolderPath, relativeSourcePath)
	} else {
		finalDestinationNotePath = path.Join(destinationObsidianConfig.VaultRoot, relativeDestinationPath, relativeSourcePath)
	}

	return ctx.Repository.CopyOrMoveFile(sourceNotePath, finalDestinationNotePath, move)
}

func (ctx *Context) CopyAttachments(absoluteSourcePath, sourceNotePath string, destinationObsidianConfig *ObsidianConfig, move bool) error {
	attachments := ctx.Repository.GetNoteAttachments(sourceNotePath)

	if len(attachments) < 1 {
		// No attachments so nothing to do
		return nil
	}

	baseAttachmentPath := path.Join(destinationObsidianConfig.VaultRoot, destinationObsidianConfig.AttachmentFolderPath)

	for _, attachmentSourcePath := range attachments {
		// Ignore any attachments at this source level as they will have been copied directly over already
		if strings.HasPrefix(attachmentSourcePath, absoluteSourcePath) {
			continue
		}

		destinationPath := path.Join(baseAttachmentPath, path.Base(attachmentSourcePath))
		err := ctx.CopyAttachment(attachmentSourcePath, destinationPath, move)

		if err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) CopyAttachment(sourcePath, destinationPath string, move bool) error {
	filesReferencingThisAttachment := ctx.Repository.GetFilesReferencingThisAttachment(sourcePath)

	if move && len(filesReferencingThisAttachment) > 1 {
		log.Printf("Copying instead of moving \"%s\" as it is referenced by %s", sourcePath, utils.Pluralize("other note", len(filesReferencingThisAttachment)-1))
		return ctx.Repository.CopyOrMoveFile(sourcePath, destinationPath, false)
	}

	return ctx.Repository.CopyOrMoveFile(sourcePath, destinationPath, move)
}
