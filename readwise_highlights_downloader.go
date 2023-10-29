package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Readwise Highlights Downloader")

	configuration, err := GetConfiguration()
	if err != nil {
		fmt.Println("configuration could not be read")
		os.Exit(1)
	}

	if len(strings.TrimSpace(configuration.apiToken)) == 0 {
		fmt.Println("please specify an api token")
		os.Exit(1)
	}

	err = checkOutputDir(configuration.outputDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	highlightsReader := NewHighlightsReader(configuration.apiToken)
	allHighlights, err := highlightsReader.ReadHighlights(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, highlights := range allHighlights {
		err := createHighlightFiles(highlights, configuration.outputDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func checkOutputDir(outputDir string) error {
	if len(strings.TrimSpace(outputDir)) == 0 {
		return fmt.Errorf("please specify a valid output directory")
	}
	_, err := os.Stat(outputDir)
	if err != nil {
		return fmt.Errorf("the specified output directory \"%v\" does not exist", outputDir)
	}
	return nil
}

func createHighlightFiles(highlights Highlights, outputDir string) error {
	fmt.Printf("Generating highlight files in %v...\n", outputDir)
	for _, book := range highlights.Results {
		highlightsCreator := NewBookHighlightsCreator(book, outputDir)
		err := highlightsCreator.CreateBookHighlightsFile()
		if err != nil {
			return err
		}
	}
	return nil
}
