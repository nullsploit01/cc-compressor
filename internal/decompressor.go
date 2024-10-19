package internal

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type DeCompressor struct {
	Filename       string
	FrequencyTable map[string]uint64
}

func Decompress(file *os.File, outputFileName string) error {
	decompressor := &DeCompressor{
		Filename: filepath.Base(file.Name()),
	}
	if outputFileName == "" {
		outputFileName = fmt.Sprintf("decompressed_%s", file.Name())
	}

	magic := make([]byte, 4)

	_, err := file.Read(magic)
	if err != nil {
		return err
	}

	if string(magic) != "1337" {
		return fmt.Errorf("invalid magic number")
	}

	reader := bufio.NewReader(file)

	var treeLength uint32
	err = binary.Read(reader, binary.LittleEndian, &treeLength)
	if err != nil {
		return err
	}

	treeData := make([]byte, treeLength)
	_, err = io.ReadFull(reader, treeData)
	if err != nil {
		return err
	}

	treeReader := bufio.NewReader(strings.NewReader(string(treeData)))
	huffmanRoot, err := DeserializeHuffmanTree(treeReader)
	if err != nil {
		return err
	}

	decodedData, err := decompressor.DecodeHuffmanData(reader, huffmanRoot)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	_, err = outputFile.Write(decodedData)
	if err != nil {
		return err
	}

	return nil
}

func (d *DeCompressor) ReadFrequencyTable(file *os.File) (map[string]uint64, error) {
	defer file.Seek(0, 0)
	var numberOfChars uint32

	err := binary.Read(file, binary.LittleEndian, &numberOfChars)
	if err != nil {
		return nil, err
	}

	frequencyTable := make(map[string]uint64)

	for i := 0; i < int(numberOfChars); i++ {
		char := make([]byte, 1)
		_, err = file.Read(char)
		if err != nil {
			return nil, err
		}

		var freq uint64
		err = binary.Read(file, binary.LittleEndian, &freq)
		if err != nil {
			return nil, err
		}

		frequencyTable[string(char)] = freq
	}

	return frequencyTable, nil
}

func (d *DeCompressor) DecodeHuffmanData(inputFile *bufio.Reader, huffmanRoot *HuffmanNode) ([]byte, error) {
	var decodedData []byte
	currentNode := huffmanRoot
	reader := bufio.NewReader(inputFile)

	for {
		byteValue, err := reader.ReadByte()
		if err == io.EOF {
			break // End of file, stop reading
		}
		if err != nil {
			return nil, err
		}

		// Process each bit of the byte
		for i := 7; i >= 0; i-- {
			bit := (byteValue >> i) & 1
			fmt.Printf("Processing bit: %d from byte: %08b\n", bit, byteValue) // Debug log

			if bit == 1 {
				currentNode = currentNode.Right
				fmt.Println("Moved right")
			} else {
				currentNode = currentNode.Left
				fmt.Println("Moved left")
			}

			// If we reach a leaf node, append the character
			if currentNode.Left == nil && currentNode.Right == nil {
				decodedData = append(decodedData, currentNode.Character[0])
				fmt.Printf("Decoded character: %s\n", currentNode.Character) // Debug log
				currentNode = huffmanRoot                                    // Reset to the root for the next character
			}
		}
	}

	return decodedData, nil
}
