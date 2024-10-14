package cmd

import (
	"os"

	"github.com/nullsploit01/cc-compressor/internal"
	"github.com/spf13/cobra"
)

var compressFile bool

var rootCmd = &cobra.Command{
	Use:   "cccmp [flags] [file]",
	Short: "Another text compressor",
	Long: `Another text file compressor that uses Huffman Encoding for efficient compression. 
It assigns shorter codes to frequent characters and longer ones to less frequent ones. 
This achieves optimal file size reduction. 
The tool also supports decompression to restore the original file. 
It’s a reliable method for reducing text file sizes.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.PrintErr("Error: A file name is required.\n")
			cmd.Usage()
			return
		}

		file, err := os.Open(args[0])
		if err != nil {
			cmd.PrintErrf("Error reading file: %v\n", err)
			return
		}

		defer file.Close()

		if compressFile {
			err := internal.Compress(file)
			if err != nil {
				cmd.PrintErrln("error occured while compressing file", err)
				os.Exit(1)
			}
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&compressFile, "compress", "c", false, "Compress text file")
}
