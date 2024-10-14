package internal

import (
	"bufio"
	"os"
)

type Compressor struct {
	Filename       string
	FrequencyTable map[string]uint64
}

func Compress(file *os.File) error {
	compressor := &Compressor{
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
