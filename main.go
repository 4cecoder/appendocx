package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string
var appendText string

func main() {
	rootCmd := &cobra.Command{
		Use:   "docx-append",
		Short: "A simple CLI to append text to a DOCX file",
		Run: func(cmd *cobra.Command, args []string) {
			if inputFile == "" || outputFile == "" || appendText == "" {
				fmt.Println("Error: Missing required flags.")
				cmd.Help()
				os.Exit(1)
			}

			d, err := docx.ReadDocxFromFS(inputFile)
			if err != nil {
				log.Fatal("Error reading input file:", err)
			}

			d.AppendText(appendText)

			err = d.WriteToFile(outputFile)
			if err != nil {
				log.Fatal("Error writing output file:", err)
			}
			fmt.Println("Text successfully appended and saved to output file.")
		},
	}

	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input DOCX file path")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output DOCX file path")
	rootCmd.Flags().StringVarP(&appendText, "text", "t", "", "Text to append to the DOCX file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
