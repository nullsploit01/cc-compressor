package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Compressor struct {
	Filename       string
	FrequencyTable map[string]uint64
}

func Compress(file *os.File) error {
	currTime := time.Now()
	compressor := &Compressor{
		Filename:       filepath.Base(file.Name()),
		FrequencyTable: make(map[string]uint64),
	}

	err := compressor.GenerateFrequencyTable(file)
	if err != nil {
		return err
	}

	root := BuildHuffmanTree(compressor.FrequencyTable)
	huffmanCodes := make(map[string]string)
	GenerateHuffmanCodes(root, "", huffmanCodes)

	fmt.Printf("finished compressing file %s in %f seconds\n", compressor.Filename, time.Since(currTime).Seconds())

	return nil
}

func (c *Compressor) GenerateFrequencyTable(file *os.File) error {
	defer file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		c.FrequencyTable[scanner.Text()] += 1
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (c *Compressor) GenerateHuffmanCodingTree() {

}
