package main

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestCompareFoldersReturnsTrueWhenSameContents(t *testing.T) {
	ctx := &Context{}
	same, err := ctx.CompareFolders(testDataPath, testDataPath)

	assert.NoError(t, err)
	assert.True(t, same)
}

func TestCompareFoldersReturnsFalseWhenContentsAreDifferent(t *testing.T) {
	ctx := &Context{}
	same, err := ctx.CompareFolders(path.Join(testDataPath, "_v"), path.Join(testDataPath, "b"))

	assert.NoError(t, err)
	assert.False(t, same)
}
