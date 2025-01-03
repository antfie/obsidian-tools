package main

import (
	"github.com/antfie/obsidian-tools/crypto"
	"github.com/antfie/obsidian-tools/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CompareFiles(leftFilePath, rightFilePath string) (bool, error) {
	leftHash, err := crypto.HashFile(leftFilePath)

	if err != nil {
		return false, err
	}

	rightHash, err := crypto.HashFile(rightFilePath)

	if err != nil {
		return false, err
	}

	return leftHash == rightHash, nil
}

func (ctx *Context) CompareFolders(leftPath, rightPath string) (bool, error) {
	same := true

	err := filepath.Walk(leftPath, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		baseName := path.Base(currentPath)

		if info.IsDir() {
			if utils.IsInArray(baseName, ctx.Config.FolderNamesToIgnore) {
				return filepath.SkipDir
			}
		} else {
			if !utils.IsInArray(baseName, ctx.Config.FileNamesToIgnore) {
				fileName := strings.TrimPrefix(strings.TrimPrefix(currentPath, leftPath), "/")
				same, err = CompareFiles(path.Join(leftPath, fileName), path.Join(rightPath, fileName))

				if !same {
					return filepath.SkipDir
				}
			}
		}

		return nil
	})

	return same, err
}
