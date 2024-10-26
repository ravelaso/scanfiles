package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ravelaso/scanfiles"
)

func main() {
	// Check for the presence of search string argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: scanfiles [directory] <searchString>")
		return
	}

	// Set rootDir to current directory if not provided
	rootDir := "."
	searchString := ""

	if len(os.Args) == 2 {
		// Only search string is provided, assume current directory
		searchString = os.Args[1]
	} else if len(os.Args) >= 3 {
		// Both directory and search string are provided
		rootDir = os.Args[1]
		searchString = os.Args[2]
	}

	// Step 1: List all files in the given directory
	files, err := scanfiles.ListFilesInDirectory(rootDir)
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
		go scanfiles.SearchFile(file, searchString, results, ctx, &wg) // Access function from scanfiles package
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
