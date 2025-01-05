package main

import (
	"fmt"
	"github.com/antfie/obsidian-tools/utils"
	"path"
)

const emptyFileHash = "3Qdi4H7DKe5XH67XhEgeE75C7Tp8JKWWT7PYiTaZHCrtwMdNiw3sUWfHm482HiPDcWsQMpEvJnK5xn68hhSiassT"

func (ctx *Context) FindDuplicates(source string) error {
	sourceAbs := assertSourceExists(source)
	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		return err
	}

	err = ctx.Repository.PopulateFromVault(sourceObsidianRoot, true)

	if err != nil {
		return err
	}

	hashes := make(map[string][]string)

	for filePath, obsidianFile := range ctx.Repository.GetAllFiles() {
		val, found := hashes[obsidianFile.Hash]

		if found {
			hashes[obsidianFile.Hash] = append(val, filePath)
		} else {
			hashes[obsidianFile.Hash] = []string{filePath}
		}
	}

	paths, emptyFileHashFound := hashes[emptyFileHash]

	if emptyFileHashFound {
		utils.PrintFormattedTitle("Empty Files")

		for _, filePath := range paths {
			fmt.Println(filePath)
		}
	}

	duplicateFilesTitleRendered := false

	for hash, paths := range hashes {
		if hash == emptyFileHash {
			continue
		}

		if len(paths) > 1 {
			if !duplicateFilesTitleRendered {
				if emptyFileHashFound {
					fmt.Println()
				}

				utils.PrintFormattedTitle("Duplicate Files")
				duplicateFilesTitleRendered = true
			}

			fmt.Println(hash)
			for index, filePath := range paths {
				if index < len(paths)-1 {
					fmt.Print("├")

				} else {
					fmt.Print("└")
				}

				fmt.Printf(" %s\n", filePath)
			}
		}
	}

	// Detect attachment duplicates
	duplicateAttachmentFileNames := make(map[string][]string)

	for _, obsidianFile := range ctx.Repository.GetAllFiles() {
		for _, attachmentFilePath := range obsidianFile.Attachments {
			attachmentFileName := path.Base(attachmentFilePath)

			val, found := duplicateAttachmentFileNames[attachmentFileName]

			if found {
				if !utils.IsInArray(attachmentFilePath, duplicateAttachmentFileNames[attachmentFileName]) {
					duplicateAttachmentFileNames[attachmentFileName] = append(val, attachmentFilePath)
				}
			} else {
				duplicateAttachmentFileNames[attachmentFileName] = []string{attachmentFilePath}
			}
		}
	}

	duplicateAttachmentsTitleRendered := false

	for attachmentFileName, attachmentPaths := range duplicateAttachmentFileNames {
		if len(attachmentPaths) > 1 {
			if !duplicateAttachmentsTitleRendered {
				if emptyFileHashFound || duplicateFilesTitleRendered {
					fmt.Println()
				}

				utils.PrintFormattedTitle("Duplicate Attachments")
				duplicateAttachmentsTitleRendered = true
			}

			fmt.Println(attachmentFileName)
			for index, filePath := range attachmentPaths {
				if index < len(attachmentPaths)-1 {
					fmt.Print("├")

				} else {
					fmt.Print("└")
				}

				fmt.Printf(" %s\n", filePath)
			}
		}
	}

	return nil
}
