package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"path/filepath"
	"testing"
)

var testDataPath = path.Join("test", "data")
var nothingToIgnore []string

func createEmptyTempTestDataPath(t *testing.T) string {
	tempTestDataPath, err := os.MkdirTemp("", "obsidian-tools-")
	assert.NoError(t, err)

	tempTestDataAbsolutePath, err := filepath.Abs(tempTestDataPath)
	assert.NoError(t, err)

	return tempTestDataAbsolutePath
}

func createTempTestDataPath(t *testing.T) string {
	tempTestDataPath := createEmptyTempTestDataPath(t)

	testDataAbsolutePath, err := filepath.Abs(testDataPath)
	assert.NoError(t, err)

	// Populate test data
	err = CopyOrMoveFiles(testDataAbsolutePath, tempTestDataPath, nothingToIgnore, nothingToIgnore, false)
	assert.NoError(t, err)

	return tempTestDataPath
}
