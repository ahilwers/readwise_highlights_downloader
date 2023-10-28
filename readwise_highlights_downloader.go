package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Readwise Highlights Downloader")

	configuration, err := GetConfiguration()
	if err != nil {
		fmt.Errorf("api token not specified")
		os.Exit(1)
	}
	fmt.Println(configuration.apiToken)

	requestUrl := "https://readwise.io/api/v2/export/"
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Token %v", configuration.apiToken))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	var highlights Highlights
	json.Unmarshal(responseBody, &highlights)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	fmt.Println(highlights.Count)

	for _, result := range highlights.Results {
		fmt.Println(result.Title)
	}
}
