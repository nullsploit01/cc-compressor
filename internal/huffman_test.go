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

	runeFrequencies := convertStringToRuneMap(frequencies)

	root := internal.BuildHuffmanTree(runeFrequencies)

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

	runeFrequencies := convertStringToRuneMap(frequencies)

	root := internal.BuildHuffmanTree(runeFrequencies)

	codes := make(map[rune]string)
	internal.GenerateHuffmanCodes(root, "", codes)

	expectedCodes := map[rune]string{
		100: "101",
		101: "111",
		102: "0",
		97:  "1100",
		98:  "1101",
		99:  "100",
	}

	if !compareMaps(expectedCodes, codes) {
		t.Errorf("Expected: %v, Got: %v", expectedCodes, codes)
	}
}

func TestGenerateHuffmanCodesSingleNode(t *testing.T) {
	frequencies := map[string]uint64{
		"a": 1,
	}

	runeFrequencies := convertStringToRuneMap(frequencies)

	root := internal.BuildHuffmanTree(runeFrequencies)

	codes := make(map[rune]string)
	internal.GenerateHuffmanCodes(root, "", codes)

	expectedCodes := map[rune]string{
		97: "",
	}

	if !reflect.DeepEqual(codes, expectedCodes) {
		t.Errorf("Generated Huffman code for single node doesn't match expected code.\nExpected: %v\nGot: %v", expectedCodes, codes)
	}
}

func convertStringToRuneMap(stringMap map[string]uint64) map[rune]uint64 {
	runeMap := make(map[rune]uint64)
	for key, value := range stringMap {
		if len(key) == 1 { // Ensure the string key is a single character
			runeMap[rune(key[0])] = value
		}
	}
	return runeMap
}

func compareMaps(expected, got map[rune]string) bool {
	if len(expected) != len(got) {
		return false
	}

	for key, expectedValue := range expected {
		gotValue, exists := got[key]
		if !exists || expectedValue != gotValue {
			return false
		}
	}

	return true
}
