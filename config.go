package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

type yamlConfig struct {
	LogFilePath                 string   `yaml:"log_file_path"`
	DryRun                      bool     `yaml:"dry_run"`
	MaxConcurrentFileOperations int      `yaml:"max_concurrent_file_operations"`
	FileNamesToIgnore           []string `yaml:"file_names_to_ignore"`
	FolderNamesToIgnore         []string `yaml:"folder_names_to_ignore"`
}

type Config struct {
	LogFilePath                 string
	DryRun                      bool
	MaxConcurrentFileOperations int
	FileNamesToIgnore           []string
	FolderNamesToIgnore         []string
}

func Load(defaultConfigData []byte) *Config {
	configFile := "config.yaml"

	if len(os.Args) == 2 {
		configFile = os.Args[1]

		_, err := os.Stat(configFile)

		if err != nil {
			log.Fatalf("Could not open config file \"%s\".", configFile)
		}
	}

	_, err := os.Stat(configFile)

	if err != nil {
		log.Print("No config file found. Creating a new config file...")
		err := os.WriteFile(configFile, defaultConfigData, 0600)

		if err != nil {
			log.Fatal(err)
		}
	}

	return parseConfigFile(configFile)
}

func parseConfigFile(configFilePath string) *Config {
	yamlFile, err := os.ReadFile(path.Clean(configFilePath))

	if err != nil {
		log.Panic(err)
	}

	config := &yamlConfig{}

	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		log.Panic(err)
	}

	return &Config{
		LogFilePath:                 configFilePath,
		DryRun:                      config.DryRun,
		MaxConcurrentFileOperations: config.MaxConcurrentFileOperations,
		FileNamesToIgnore:           config.FileNamesToIgnore,
		FolderNamesToIgnore:         config.FolderNamesToIgnore,
	}
}
