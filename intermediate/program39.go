package main

import (
	"archive/zip"
	"fmt"
	"os"
)

func main() {
	zipName := "archive.zip"
	newZipFile, err := os.Create(zipName)
	if err != nil {
		panic(err)
	}
	defer newZipFile.Close()
	defer os.Remove(zipName) // Cleanup for demo

	// Create a new zip writer
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add a file to the zip
	fileWriter, err := zipWriter.Create("readme.txt")
	if err != nil {
		panic(err)
	}
	
	// Write content into the zipped file
	_, err = fileWriter.Write([]byte("This file is compressed inside a zip archive."))
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully created", zipName)
}