package internal_test

import (
	"reflect"
	"testing"

	"github.com/nullsploit01/cc-compressor/internal"
)

func TestBuildHuffmanTree(t *testing.T) {
	frequencies := map[string]uint64{
		"a": 5,
		"b": 9,
		"c": 12,
		"d": 13,
		"e": 16,
		"f": 45,
	}

	root := internal.BuildHuffmanTree(frequencies)

	expectedRootFreq := uint64(100) // Sum of all frequencies
	if root.Frequency != expectedRootFreq {
		t.Errorf("Expected root frequency %d, but got %d", expectedRootFreq, root.Frequency)
	}
}

func TestGenerateHuffmanCodes(t *testing.T) {
	frequencies := map[string]uint64{
		"a": 5,
		"b": 9,
		"c": 12,
		"d": 13,
		"e": 16,
		"f": 45,
	}

	root := internal.BuildHuffmanTree(frequencies)

	codes := make(map[string]string)
	internal.GenerateHuffmanCodes(root, "", codes)

	expectedCodes := map[string]string{
		"a": "1100",
		"b": "1101",
		"c": "100",
		"d": "101",
		"e": "111",
		"f": "0",
	}

	if !reflect.DeepEqual(codes, expectedCodes) {
		t.Errorf("Generated Huffman codes don't match expected codes.\nExpected: %v\nGot: %v", expectedCodes, codes)
	}
}

func TestGenerateHuffmanCodesSingleNode(t *testing.T) {
	frequencies := map[string]uint64{
		"a": 1,
	}

	root := internal.BuildHuffmanTree(frequencies)

	codes := make(map[string]string)
	internal.GenerateHuffmanCodes(root, "", codes)

	expectedCodes := map[string]string{
		"a": "",
	}

	if !reflect.DeepEqual(codes, expectedCodes) {
		t.Errorf("Generated Huffman code for single node doesn't match expected code.\nExpected: %v\nGot: %v", expectedCodes, codes)
	}
}
