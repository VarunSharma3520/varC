package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/VarunSharma3520/varC/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	// Create a temporary file for testing
	tmpFile := filepath.Join(os.TempDir(), "testfile.txt")
	content := "line1\nline2\nline3\n"

	// Write content to temp file
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	assert.NoError(t, err, "should write temp file without error")

	// Read using utils.ReadFile
	lines, err := utils.ReadFile(tmpFile)
	assert.NoError(t, err, "ReadFile should not return error")
	assert.Equal(t, []string{"line1", "line2", "line3"}, lines, "lines should match file content")

	// Clean up temp file
	_ = os.Remove(tmpFile)
}

func TestReadFile_FileNotExist(t *testing.T) {
	// Try reading a non-existent file
	_, err := utils.ReadFile("non_existent_file.txt")
	assert.Error(t, err, "should return error for non-existent file")
}
