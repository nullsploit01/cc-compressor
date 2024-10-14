package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cccmp",
	Short: "Another text compressor",
	Long: `Another text file compressor that uses Huffman Encoding for efficient compression. 
It assigns shorter codes to frequent characters and longer ones to less frequent ones. 
This achieves optimal file size reduction. 
The tool also supports decompression to restore the original file. 
It’s a reliable method for reducing text file sizes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
