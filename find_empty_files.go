package main

import (
	"log"
	"os"
)

func findEmptyFiles(source string) {
	sourceAbs := assertSourceExists(source)

	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	allFiles := getAllFiles(sourceObsidianRoot)

	for _, file := range allFiles {
		stat, err := os.Stat(file)

		if err != nil {
			log.Fatal(err)
		}

		if stat.Size() == 0 {
			println(file)
		}
	}
}
