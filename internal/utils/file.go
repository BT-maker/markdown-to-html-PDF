package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadFile reads the content of a file and returns it as a string
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return string(content), nil
}

// WriteFile writes content to a file
func WriteFile(filename string, content string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}
	return nil
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// IsDirectory checks if a path is a directory
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// GetMarkdownFiles returns all .md files in a directory
func GetMarkdownFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file has .md extension
		if strings.HasSuffix(strings.ToLower(path), ".md") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", dir, err)
	}

	return files, nil
}

// EnsureDirectoryExists creates a directory if it doesn't exist
func EnsureDirectoryExists(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	return nil
}

// GetRelativePath returns the relative path from base to target
func GetRelativePath(base, target string) (string, error) {
	rel, err := filepath.Rel(base, target)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path from %s to %s: %w", base, target, err)
	}
	return rel, nil
}

// ReadFileFromReader reads content from an io.Reader and returns it as a string
func ReadFileFromReader(reader interface{ Read([]byte) (int, error) }) (string, error) {
	var content []byte
	buffer := make([]byte, 1024)
	
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			content = append(content, buffer[:n]...)
		}
		if err != nil {
			break
		}
	}
	
	return string(content), nil
}
