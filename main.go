package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ListFilesInDirectory lists all files in the given directory (non-recursively)
func ListFilesInDirectory(rootDir string) ([]string, error) {
	files := []string{}

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(rootDir, entry.Name()))
		}
	}

	return files, nil
}

// SearchFile searches for the given string in a file's content
func SearchFile(filePath string, searchString string, results chan<- string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// fmt.Printf("Starting goroutine for file: %s\n", filePath) // Debug information

	select {
	case <-ctx.Done():
		// Exit if context is canceled
		return
	default:
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filePath, err)
			return
		}

		fileContent := string(content)
		if strings.Contains(strings.ToLower(fileContent), strings.ToLower(searchString)) {
			results <- filePath
		}
	}
}

func main() {
	// Check for correct number of arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: scanfiles <directory> <searchString>")
		return
	}

	rootDir := os.Args[1]
	searchString := os.Args[2]

	// Step 1: List all files in the given directory
	files, err := ListFilesInDirectory(rootDir)
	if err != nil {
		log.Fatalf("Error listing files in directory: %v", err)
	}

	// Step 2: Create context and wait group
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan string)
	var wg sync.WaitGroup

	// Step 3: Search each file concurrently
	for _, file := range files {
		wg.Add(1)
		go SearchFile(file, searchString, results, ctx, &wg)
	}

	// Step 4: Wait for result or goroutines to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	select {
	case filePath := <-results:
		fmt.Printf("String found in: %s\n", filePath)
		cancel() // Cancel all goroutines
	case <-ctx.Done():
		return
	}
}
