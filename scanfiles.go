package scanfiles

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ListFilesInDirectory lists all files in the given directory (non-recursively)
func ListFilesInDirectory(rootDir string) ([]string, error) {
	var files []string

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
		if strings.Contains(fileContent, searchString) {
			results <- filePath
		}
	}
}
