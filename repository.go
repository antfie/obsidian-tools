package main

import (
	"github.com/antfie/obsidian-tools/crypto"
	"github.com/antfie/obsidian-tools/utils"
	"github.com/schollz/progressbar/v3"
	"log"
	"maps"
	"sort"
	"strings"
	"sync"
)

type Repository struct {
	Config *Config
	files  map[string]ObsidianFile // the key is the hash
}

func NewRepository(config *Config) *Repository {
	return &Repository{
		Config: config,
		files:  make(map[string]ObsidianFile),
	}
}

func (r *Repository) PopulateFromVault(vaultRootPath string, analyze bool) error {
	allFiles, err := GetAllFiles(vaultRootPath, r.Config.FolderNamesToIgnore, r.Config.FileNamesToIgnore)

	if err != nil {
		return nil
	}

	// Sometimes we don't want to find attachments or hash the files
	if !analyze {
		for _, file := range allFiles {
			r.files[file] = ObsidianFile{}
		}

		return nil
	}

	log.Printf("Analysing %s in vault \"%s\" with %s", utils.Pluralize("file", len(allFiles)), vaultRootPath, utils.Pluralize("thread", r.Config.MaxConcurrentFileOperations))

	// Sort by path length descending to facilitate attachment path resolution
	sort.Slice(allFiles, func(i, j int) bool {
		return len(allFiles[i]) > len(allFiles[j])
	})

	bar := progressbar.Default(int64(len(allFiles)))
	var warnings []string

	var mapMutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(allFiles))
	sem := make(chan int, r.Config.MaxConcurrentFileOperations)

	for _, filePath := range allFiles {
		sem <- 1

		go func(filePath string) {
			defer wg.Done()

			hash, err := crypto.HashFile(filePath)

			if err != nil {
				log.Fatalf("hash file %s failed: %v", filePath, err)
			}

			var attachments []string
			var warningsFromThisFile []string

			// Only markdown files can have attachments
			if isMarkdownFile(filePath) {
				attachments, warningsFromThisFile = getAttachmentsFromNote(allFiles, filePath)
			}

			// Maps are not threadsafe
			mapMutex.Lock()

			warnings = append(warnings, warningsFromThisFile...)
			r.files[filePath] = ObsidianFile{
				Hash:        hash,
				Attachments: attachments,
			}

			mapMutex.Unlock()

			err = bar.Add(1)

			if err != nil {
				log.Fatalf("failed to update progress bar: %v", err)
			}

			<-sem
		}(filePath)
	}

	// Wait for the batches to finish
	wg.Wait()

	for _, warning := range warnings {
		log.Printf(warning)
	}

	return nil
}

func (r *Repository) GetAllFilePathsInFolder(absoluteSourcePath string) []string {
	var sourceFilesToCopy []string

	for sourceFilePath := range maps.Keys(r.files) {
		if strings.HasPrefix(sourceFilePath, absoluteSourcePath) {
			sourceFilesToCopy = append(sourceFilesToCopy, sourceFilePath)
		}
	}

	return sourceFilesToCopy
}

func (r *Repository) GetNoteAttachments(notePath string) []string {
	return r.files[notePath].Attachments
}

func (r *Repository) GetAllFiles() map[string]ObsidianFile {
	return r.files
}

func (r *Repository) GetFilesReferencingThisAttachment(sourcePath string) []ObsidianFile {
	var filesReferencingThisAttachment []ObsidianFile

	for filePath, obsidianFile := range r.files {
		if !isMarkdownFile(filePath) {
			continue
		}

		if sourcePath == filePath {
			continue
		}

		for _, attachment := range obsidianFile.Attachments {
			if attachment == sourcePath {
				filesReferencingThisAttachment = append(filesReferencingThisAttachment, obsidianFile)
			}
		}
	}

	return filesReferencingThisAttachment
}

func (r *Repository) CopyOrMoveFile(sourceFilePath, destinationFilePath string, move bool) error {
	copyOrMoveText := "Copying"
	// If we are moving a file out of the vault we need to remove it from the files map
	// And any attachment references?
	if move {
		copyOrMoveText = "Moving"
		// TODO: remove the attachment from the sourceFiles map and from the array of filesReferencingThisAttachment[0]
	}

	log.Printf("%s \"%s\" to \"%s\"", copyOrMoveText, sourceFilePath, destinationFilePath)

	if r.Config.DryRun {
		return nil
	}

	return CopyOrMoveFile(sourceFilePath, destinationFilePath, move)
}
