package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestCopyNotesShouldErrorIfSourceAndDestinationAreInTheSameVault(t *testing.T) {
	tempTestDataPath := createTempTestDataPath(t)
	defer os.RemoveAll(tempTestDataPath)

	sourcePath := path.Join(tempTestDataPath, "_v/u/files/note.md")
	destinationPath := path.Join(tempTestDataPath, "_v/u/docs")

	ctx := &Context{}
	err := ctx.CopyNotes(sourcePath, destinationPath, false)

	assert.ErrorIs(t, err, ErrDestinationIsSameVaultAsSource)
}

func TestCopyNotesShouldErrorIfSourceVaultNotFound(t *testing.T) {
	ctx := &Context{}

	err := ctx.CopyNotes(testDataPath, "", false)

	assert.ErrorIs(t, err, ErrSourceVaultNotResolved)
}

func TestCopyNotesShouldErrorIfDestinationVaultNotFound(t *testing.T) {
	tempTestDataPath := createTempTestDataPath(t)
	defer os.RemoveAll(tempTestDataPath)

	sourcePath := path.Join(tempTestDataPath, "_v/u/files/note.md")

	ctx := &Context{}
	err := ctx.CopyNotes(sourcePath, tempTestDataPath, false)

	assert.ErrorIs(t, err, ErrDestinationVaultNotResolved)
}

func TestCopyNotesShouldCopyASingleNoteIfFilePassed(t *testing.T) {
	tempTestDataPath := createTempTestDataPath(t)
	defer os.RemoveAll(tempTestDataPath)

	sourcePath := path.Join(tempTestDataPath, "_v/u/files/note.md")
	destinationPath := path.Join(tempTestDataPath, "b/docs")

	ctx := &Context{}
	err := ctx.CopyNotes(sourcePath, destinationPath, false)

	assert.NoError(t, err)

	same, err := CompareFiles(sourcePath, path.Join(destinationPath, "note.md"))
	assert.NoError(t, err)
	assert.True(t, same)
}

func TestCopyNotesShouldOnlyWorkOnMarkdownFiles(t *testing.T) {
	tempTestDataPath := createTempTestDataPath(t)
	defer os.RemoveAll(tempTestDataPath)

	sourcePath := path.Join(tempTestDataPath, "_v/u/files/data.txt")
	destinationPath := path.Join(tempTestDataPath, "b/docs")

	ctx := &Context{}
	err := ctx.CopyNotes(sourcePath, destinationPath, false)

	assert.ErrorIs(t, err, ErrNotMarkdownFile)
}

func TestCopyNotesShouldCopyNotesFromAPath(t *testing.T) {
	tempTestDataPath := createTempTestDataPath(t)
	defer os.RemoveAll(tempTestDataPath)

	sourcePath := path.Join(tempTestDataPath, "_v/u/notebook")
	destinationPath := path.Join(tempTestDataPath, "b/docs")

	ctx := &Context{}
	err := ctx.CopyNotes(sourcePath, destinationPath, false)

	assert.NoError(t, err)
}
