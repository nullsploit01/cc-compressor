package internal

import (
	"bufio"
	"log"
	"os"
)

type Compressor struct {
	Filename       string
	FrequencyTable map[string]uint64
}

func Compress(file *os.File) {
	defer file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	compressor := &Compressor{
		FrequencyTable: make(map[string]uint64),
	}

	for scanner.Scan() {
		compressor.FrequencyTable[scanner.Text()] += 1
	}

	log.Println(compressor.FrequencyTable)
}
