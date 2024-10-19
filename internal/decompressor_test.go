package internal_test

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/nullsploit01/cc-compressor/internal"
)

func createMockHuffmanTree() *internal.HuffmanNode {
	return &internal.HuffmanNode{
		Left: &internal.HuffmanNode{
			Character: 'a', // 'a' -> 0
		},
		Right: &internal.HuffmanNode{
			Left:  &internal.HuffmanNode{Character: 'b'}, // 'b' -> 10
			Right: &internal.HuffmanNode{Character: 'c'}, // 'c' -> 11
		},
	}
}

func TestDecodeHuffmanData(t *testing.T) {
	root := createMockHuffmanTree()

	encodedData := []byte{0b01011000} // First 6 bits represent "010110"

	inputReader := bytes.NewReader(encodedData)
	bufReader := bufio.NewReader(inputReader)

	decompressor := &internal.DeCompressor{}

	decodedData, err := decompressor.DecodeHuffmanData(bufReader, root)
	if err != nil {
		t.Fatalf("Error while decoding Huffman data: %v", err)
	}

	expectedOutput := []byte("abcaaa")
	if !bytes.Equal(decodedData, expectedOutput) {
		t.Errorf("Expected decoded output %s, but got %s", expectedOutput, decodedData)
	}
}

func TestReadFrequencyTable(t *testing.T) {
	serializedFrequencyTable := []byte{
		0x02, 0x00, 0x00, 0x00,
		'a', 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		'b', 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	tmpFile, err := os.CreateTemp("", "freq_table_test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temp file after the test

	_, err = tmpFile.Write(serializedFrequencyTable)
	if err != nil {
		t.Fatalf("Failed to write mock data to temp file: %v", err)
	}

	decompressor := &internal.DeCompressor{}
	frequencyTable, err := decompressor.ReadFrequencyTable(tmpFile)
	if err != nil {
		t.Fatalf("Error reading frequency table: %v", err)
	}

	expectedTable := map[string]uint64{
		"a": 5,
		"b": 3,
	}

	if !compareMaps2(expectedTable, frequencyTable) {
		t.Errorf("Expected frequency table %v, but got %v", expectedTable, frequencyTable)
	}
}

func compareMaps2(expected, got map[string]uint64) bool {
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
