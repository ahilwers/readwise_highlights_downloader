package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	requestUrl := "https://readwise.io/api/v2/export/"
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err.Error())
		os.Exit(1)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Token %v", configuration.apiToken))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Error fetching highlights: %v\n", err.Error())
		os.Exit(1)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error resding reponse body: %v", err.Error())
		os.Exit(1)
	}

	var highlights Highlights
	json.Unmarshal(responseBody, &highlights)
	if err != nil {
		fmt.Printf("Error unmarshalling results: %v\n", err.Error())
		os.Exit(1)
	}

	createHighlightFiles(highlights, configuration.outputDir)
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

func createHighlightFiles(highlights Highlights, outputDir string) {
	fmt.Printf("Generating highlight files in %v...\n", outputDir)
	for _, book := range highlights.Results {
		highlightsCreator := NewBookHighlightsCreator(book, outputDir)
		highlightsCreator.CreateBookHighlightsFile()
	}
}
