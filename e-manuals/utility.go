package main

import (
	"errors"
	"os"
	"strings"
)

func scandir(dirPath string) ([]string, error) {
	f, err := os.Open(dirPath) // this is how we open a directory
	if err != nil {
		return []string{}, errors.New("Error opening directory: " + dirPath) // this is how we return an error
	}

	files, err := f.Readdir(0) // this is how we read the contents of a directory
	if err != nil {
		return []string{}, errors.New("Error reading directory: " + dirPath) // this is how we return an error
	}

	result := make([]string, len(files))
	for i := range files {
		result[i] = strings.Split(files[i].Name(), ".")[0] // this is how we split a string
	}

	return result, nil // this is how we return a value
}
