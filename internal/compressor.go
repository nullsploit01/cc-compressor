package internal

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Compressor struct {
	Filename       string
	FrequencyTable map[rune]uint64
}

func Compress(file *os.File, outputFileName string) error {
	currTime := time.Now()
	compressor := &Compressor{
		Filename:       filepath.Base(file.Name()),
		FrequencyTable: make(map[rune]uint64),
	}

	if outputFileName == "" {
		outputFileName = fmt.Sprintf("compressed_%s", compressor.Filename)
	}

	err := compressor.GenerateFrequencyTable(file)
	if err != nil {
		return err
	}

	root := BuildHuffmanTree(compressor.FrequencyTable)
	huffmanCodes := make(map[rune]string)
	GenerateHuffmanCodes(root, "", huffmanCodes)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	err = compressor.WriteHeader(outputFile, root)
	if err != nil {
		return err
	}

	err = compressor.WriteEncodedData(file, huffmanCodes, outputFile)
	if err != nil {
		return err
	}

	fmt.Printf("finished compressing file %s in %f seconds. saved as %s\n", compressor.Filename, time.Since(currTime).Seconds(), outputFileName)

	return nil
}

func (c *Compressor) GenerateFrequencyTable(file *os.File) error {
	defer file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := []rune(scanner.Text())[0]
		c.FrequencyTable[char] += 1
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (c *Compressor) WriteHeader(outputFile *os.File, root *HuffmanNode) error {
	_, err := outputFile.Write([]byte("1337")) // magic number
	if err != nil {
		return err
	}

	var treeData []byte
	SerializeHuffmanTree(root, &treeData)

	treeLength := uint32(len(treeData))
	err = binary.Write(outputFile, binary.LittleEndian, treeLength)
	if err != nil {
		return err
	}

	_, err = outputFile.Write(treeData)
	if err != nil {
		return err
	}

	return nil
}

func (c *Compressor) WriteEncodedData(inputFile *os.File, huffmanCodes map[rune]string, outputFile *os.File) error {
	defer inputFile.Seek(0, 0)

	var bitBuffer strings.Builder

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := []rune(scanner.Text())[0]
		code := huffmanCodes[char]
		bitBuffer.WriteString(code)
	}

	byteData, err := c.ConvertBitsToBytes(bitBuffer.String())
	if err != nil {
		return err
	}

	_, err = outputFile.Write(byteData)
	if err != nil {
		return err
	}

	return nil
}

func (c *Compressor) ConvertBitsToBytes(bitString string) ([]byte, error) {
	var byteData []byte
	for i := 0; i < len(bitString); i += 8 {
		end := i + 8
		if end > len(bitString) {
			end = len(bitString)
		}

		byteVal := byte(0)
		for j := i; j < end; j++ {
			byteVal <<= 1
			if bitString[j] == '1' {
				byteVal |= 1
			}
		}

		byteData = append(byteData, byteVal)
	}
	return byteData, nil
}
