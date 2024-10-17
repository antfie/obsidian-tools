package main

import (
	"log"
	"os"
	"path"
)

const emptyFileHash = "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a"

func findDuplicates(source string) {
	sourceAbs := assertSourceExists(source)
	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		log.Fatal(err)
	}

	allFiles := getAllFiles(sourceObsidianRoot)

	var items = make(map[string][]string)

	for _, filePath := range allFiles {
		data, err := os.ReadFile(path.Clean(filePath))

		if err != nil {
			log.Fatal(err)
		}

		hash := Sha512ToString(data)

		if _, ok := items[hash]; !ok {
			items[hash] = []string{}
		}

		items[hash] = append(items[hash], filePath)
	}

	for hash, files := range items {
		if len(files) > 1 {
			if hash == emptyFileHash {
				log.Println("Empty file(s)")
			} else {
				log.Println(hash)
			}

			for _, file := range files {
				log.Printf("  %s", file)
			}
		}
	}
}
