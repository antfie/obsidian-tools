package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestFindObsidianRootShouldReturnCorrectVaultRootPath(t *testing.T) {
	notePath := path.Join(testDataPath, "b/docs/house/doc.md")
	result, err := findObsidianRoot(notePath)

	assert.NoError(t, err)
	assert.Equal(t, "test/data/b", result)
}

func TestFindObsidianRootShouldReturnCorrectVaultRootPathEvenIfDestinationDoesNotExist(t *testing.T) {
	notePath := path.Join(testDataPath, "b/docs/house/fail/foo/doc.md")
	result, err := findObsidianRoot(notePath)

	assert.NoError(t, err)
	assert.Equal(t, "test/data/b", result)
}

func TestFindObsidianRootShouldReturnErrorIfRootNotFound(t *testing.T) {
	result, err := findObsidianRoot(testDataPath)

	assert.ErrorIs(t, err, os.ErrNotExist)
	assert.Empty(t, result)
}
