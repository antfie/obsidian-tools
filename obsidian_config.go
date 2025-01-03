package main

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
)

type ObsidianConfig struct {
	VaultRoot            string
	NewFileFolderPath    string `json:"newFileFolderPath"`
	AttachmentFolderPath string `json:"attachmentFolderPath"`
}

func ParseObsidianConfig(vaultRoot string) (*ObsidianConfig, error) {
	obsidianConfigFile := filepath.Join(vaultRoot, ".obsidian", "app.json")

	data, err := os.ReadFile(path.Clean(obsidianConfigFile))

	if err != nil {
		return nil, err
	}

	config := &ObsidianConfig{
		VaultRoot: vaultRoot,
	}

	err = json.Unmarshal(data, config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
