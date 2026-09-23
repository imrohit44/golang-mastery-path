package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	root := "." // Start at the current directory

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		// Skip the hidden .git directory to save time
		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			fmt.Println("File found:", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking directory:", err)
	}
}