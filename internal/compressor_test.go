package internal_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nullsploit01/cc-compressor/internal"
)

func TestGenerateFrequencyTable(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	content := "hello world"
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatalf("Failed to seek in temp file: %v", err)
	}

	compressor := &internal.Compressor{
		Filename:       filepath.Base(tempFile.Name()),
		FrequencyTable: make(map[string]uint64),
	}

	err = compressor.GenerateFrequencyTable(tempFile)
	if err != nil {
		t.Fatalf("Error generating frequency table: %v", err)
	}

	expectedFreq := map[string]uint64{
		"h": 1,
		"e": 1,
		"l": 3,
		"o": 2,
		" ": 1,
		"w": 1,
		"r": 1,
		"d": 1,
	}

	for char, expectedCount := range expectedFreq {
		if compressor.FrequencyTable[char] != expectedCount {
			t.Errorf("Expected frequency of '%s' to be %d, got %d", char, expectedCount, compressor.FrequencyTable[char])
		}
	}
}

func TestCompress(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up after test

	content := "hello"
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatalf("Failed to seek in temp file: %v", err)
	}

	err = internal.Compress(tempFile, "output")
	if err != nil {
		t.Fatalf("Compress function returned an error: %v", err)
	}
}
