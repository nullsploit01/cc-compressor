# CCCMP - Another Compression Tool

This project is a custom implementation of a compression and decompression tool based on Huffman coding. It was created as part of a coding challenge [here](https://codingchallenges.fyi/challenges/challenge-huffman/). The utility is designed to compress and decompress files efficiently using Huffman coding. It includes support for both compression and decompression through a command-line interface (CLI).

## Features

- Compress files using Huffman coding.
- Decompress files back to their original state.
- Validate compressed files for structural correctness.
- Command-line interface (CLI) powered by Cobra for easy usage.
- Supports file sizes up to 100MB for compression and decompression.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- You need to have Go installed on your machine (Go 1.15 or later is recommended).
- You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/nullsploit01/cc-compressor
cd cccmp
```

### Building

Compile the project using:

```bash
go build -o cccmp
```

### Testing

Run all tests:

```bash
go test ./...
```

### Usage

To run the utility, you can either compress or decompress files using the following commands:

#### Compress a file:

```bash
./cccmp -c path/to/inputfile.txt -o path/to/outputfile.txt
```

#### Decompress a file:

```bash
./cccmp -d path/to/inputfile.txt -o path/to/outputfile.txt
```

### Example Usage

```bash
# Compressing a file
./cccmp -c test_data/les_miserables.txt -o compressed.txt

# Decompressing a file
./cccmp -d compressed.txt -o decompressed.txt
```
