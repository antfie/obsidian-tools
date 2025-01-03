package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"testing"
)

var filesRelativeToRootPath = []string{
	"_v/u/.obsidian/.gitkeep",
	"_v/u/files/note.md",
	"_v/u/files/data.txt",
	"_v/u/notebook/x/bar.md",
	"_v/u/notebook/foo.md",
	"a/abc.txt",
	"abc.txt",
	"b/.obsidian/.gitkeep",
	"b/docs/doc.txt",
	"b/docs/house/doc.md",
	"b/def.txt",
	"b/z/a/abc.txt",
	"b/z/z.txt",
	"zz/abc.txt",
}

func TestGetAllFilesRelativeToRootPath(t *testing.T) {
	files, err := GetAllFiles(testDataPath, nothingToIgnore, nothingToIgnore)

	assert.NoError(t, err)
	assert.ElementsMatch(t, filesRelativeToRootPath, files)
}

//
//func TestGetAllFilesAbsolutePath(t *testing.T) {
//	c := &Config{}
//	files := c.GetAllFiles(testDataPath, false)
//
//	expectedAbsolutePaths := appendBasePathToFilePaths(absOrPanic(testDataPath), filesRelativeToRootPath)
//
//	assert.ElementsMatch(t, expectedAbsolutePaths, files)
//}

// @antfie
func TestGetAllFilesIgnoresFiles(t *testing.T) {
	files, err := GetAllFiles(testDataPath, nothingToIgnore, []string{"abc.txt"})

	assert.NoError(t, err)

	expectedFilePaths := []string{
		"_v/u/.obsidian/.gitkeep",
		"_v/u/files/note.md",
		"_v/u/files/data.txt",
		"_v/u/notebook/x/bar.md",
		"_v/u/notebook/foo.md",
		"b/.obsidian/.gitkeep",
		"b/docs/doc.txt",
		"b/docs/house/doc.md",
		"b/def.txt",
		"b/z/z.txt",
	}

	assert.ElementsMatch(t, expectedFilePaths, files)
}

func TestGetAllFilesIgnoresFolders(t *testing.T) {
	files, err := GetAllFiles(testDataPath, []string{"a"}, nothingToIgnore)

	assert.NoError(t, err)

	expectedFilePaths := []string{
		"_v/u/.obsidian/.gitkeep",
		"_v/u/files/note.md",
		"_v/u/files/data.txt",
		"_v/u/notebook/x/bar.md",
		"_v/u/notebook/foo.md",
		"abc.txt",
		"b/.obsidian/.gitkeep",
		"b/docs/doc.txt",
		"b/docs/house/doc.md",
		"b/def.txt",
		"b/z/z.txt",
		"zz/abc.txt",
	}

	assert.ElementsMatch(t, expectedFilePaths, files)
}

func TestCopyAllFilesRecursivelyCopiesAllFiles(t *testing.T) {
	tempTestDataPath := createTempTestDataPath(t)
	defer os.RemoveAll(tempTestDataPath)

	ctx := &Context{}
	equal, err := ctx.CompareFolders(testDataPath, tempTestDataPath)

	assert.NoError(t, err)
	assert.True(t, equal)
}

func TestCopyAllFilesWillIgnoreFilesIfTheyExist(t *testing.T) {
	tempTestDataPath := createTempTestDataPath(t)
	defer os.RemoveAll(tempTestDataPath)

	// Alter a file
	err := os.WriteFile(path.Join(tempTestDataPath, "a/abc.txt"), []byte("data"), 0600)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	err = CopyOrMoveFiles(testDataPath, tempTestDataPath, nothingToIgnore, nothingToIgnore, false)
	assert.NoError(t, err)

	resultFiles, err := GetAllFiles(tempTestDataPath, nothingToIgnore, nothingToIgnore)

	assert.NoError(t, err)
	assert.Len(t, resultFiles, 15)
}

func TestMarkdownFileDetection(t *testing.T) {
	assert.True(t, isMarkdownFile("foo/bar.mD"))
	assert.False(t, isMarkdownFile("foo/bar.mDa"))
	assert.False(t, isMarkdownFile("foo/bar"))
}

func TestIsFileShouldReturnTrueForAFile(t *testing.T) {
	result, err := IsFile(path.Join(testDataPath, "_v/u/.obsidian/.gitkeep"))
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsFileShouldReturnFalseForAPath(t *testing.T) {
	result, err := IsFile(testDataPath)
	assert.NoError(t, err)
	assert.False(t, result)
}
