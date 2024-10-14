package internal

import (
	"bufio"
	"os"
	"path/filepath"
)

type Compressor struct {
	Filename       string
	FrequencyTable map[string]uint64
}

func Compress(file *os.File) error {
	compressor := &Compressor{
		Filename:       filepath.Base(file.Name()),
		FrequencyTable: make(map[string]uint64),
	}

	err := compressor.GenerateFrequencyTable(file)
	if err != nil {
		return err
	}

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
