package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	lastUpdateTime := readLastUpdateTime()
	highlightsReader := NewHighlightsReader(configuration.apiToken)
	allHighlights, err := highlightsReader.ReadHighlights(lastUpdateTime)
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

	err = writeLastUpdateTime()
	if err != nil {
		fmt.Printf("Could not write last updaste time: %v", err.Error())
		os.Exit(1)
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

func writeLastUpdateTime() error {
	path := getStateFilepath()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(time.Now().UTC().Format(time.RFC3339))
	return nil
}

func readLastUpdateTime() *time.Time {
	path := getStateFilepath()
	_, err := os.Stat(path)
	if err != nil {
		return nil
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	timeString := string(bytes)
	lastUpdateTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return nil
	}
	return &lastUpdateTime
}

func getStateFilepath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		dirname = "./"
	}
	return filepath.Join(dirname, ".readwise_highlights_downloader_lastupdate")
}
