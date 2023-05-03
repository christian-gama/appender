package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <file_extensions> <directory> <excluded_directories>")
		os.Exit(1)
	}

	fileExtensions := strings.Split(os.Args[1], ",")
	directory := os.Args[2]

	var excludedDirs []string
	if len(os.Args) >= 4 {
		excludedDirs = strings.Split(os.Args[3], ",")
	} else {
		excludedDirs = []string{}
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && contains(excludedDirs, info.Name()) {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)
			if contains(fileExtensions, ext) {
				appendToFile(path)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func appendToFile(filePath string) {
	outputFile, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening output.txt:", err)
		return
	}
	defer outputFile.Close()

	header := fmt.Sprintf("# %s\n", filePath)
	outputFile.WriteString(header)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	outputFile.Write(content)
	outputFile.WriteString("\n")
}
