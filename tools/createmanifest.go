package tools

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateManifest() {
	srcDir := "web/assetsSrc"
	fileMap := make(map[string]string)
	outputFile := "web/assetsManifest.json"

	// Process directory recursively
	err := processDirectory(srcDir, srcDir, fileMap)
	if err != nil {
		log.Fatalf("Error processing directory: %v", err)
	}

	// Convert map to JSON
	jsonData, err := json.MarshalIndent(fileMap, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	// Create and write to output file
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Printf("Successfully created %s with %d files\n", outputFile, len(fileMap))
}

func processDirectory(rootDir, currentDir string, fileMap map[string]string) error {
	// Read all files and directories in the current directory
	entries, err := os.ReadDir(currentDir)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", currentDir, err)
	}

	for _, entry := range entries {
		// Get the full path
		fullPath := filepath.Join(currentDir, entry.Name())
		// Get the relative path from root directory
		relativePath, err := filepath.Rel(rootDir, fullPath)
		if err != nil {
			return fmt.Errorf("error getting relative path for %s: %v", fullPath, err)
		}

		if entry.IsDir() {
			// Recursively process subdirectories
			err = processDirectory(rootDir, fullPath, fileMap)
			if err != nil {
				return err
			}
		} else {
			// Process files
			file, err := os.Open(fullPath)
			if err != nil {
				return fmt.Errorf("error opening file %s: %v", fullPath, err)
			}
			defer file.Close()

			// Create MD5 hash
			hash := md5.New()
			_, err = io.Copy(hash, file)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", fullPath, err)
			}

			// Calculate the MD5 sum
			hashSum := fmt.Sprintf("%x", hash.Sum(nil))
			parts := strings.Split(entry.Name(), ".")
			parts[0] = "/" + parts[0] + "-" + string(hashSum)
			// Include the relative path in the value
			hashedFilename := filepath.Join(filepath.Dir(relativePath), strings.Join(parts, "."))
			if hashedFilename[0] == '.' { // Handle case where there's no parent directory
				hashedFilename = strings.TrimPrefix(hashedFilename, "./")
			}
			fileMap["/"+relativePath] = "/" + hashedFilename
		}
	}
	return nil
}
