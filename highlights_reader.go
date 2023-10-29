package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HighlightsReader interface {
	ReadHighlights(updatedAfter *time.Time) ([]Highlights, error)
}

type highlightsReader struct {
	apiToken string
}

func NewHighlightsReader(apiToken string) HighlightsReader {
	return &highlightsReader{
		apiToken: apiToken,
	}
}

func (hr *highlightsReader) ReadHighlights(updatedAfter *time.Time) ([]Highlights, error) {
	var allHighlights []Highlights

	allHighlights, nextPage, err := hr.appendHighlightsForPage(allHighlights, 0)
	if err != nil {
		return nil, err
	}

	for nextPage > 0 {
		allHighlights, nextPage, err = hr.appendHighlightsForPage(allHighlights, nextPage)
		if err != nil {
			return nil, err
		}
	}
	return allHighlights, nil
}

func (hr *highlightsReader) appendHighlightsForPage(highlights []Highlights, pageNumber int) ([]Highlights, int, error) {
	highlightsForPage, err := hr.readHighlightsForPage(pageNumber)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading highlights: %v", err.Error())
	}
	allHighlights := append(highlights, *highlightsForPage)
	nextPage := highlightsForPage.NextPageCursor
	return allHighlights, nextPage, nil
}

func (hr *highlightsReader) readHighlightsForPage(pageNumber int) (*Highlights, error) {
	fmt.Printf("Requesting highlights for page %v...\n", pageNumber)
	requestUrl := "https://readwise.io/api/v2/export/"
	if pageNumber != 0 {
		requestUrl = fmt.Sprintf("%v?pageCursor=%v", requestUrl, pageNumber)
	}
	fmt.Println(requestUrl)
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err.Error())
	}

	request.Header.Add("Authorization", fmt.Sprintf("Token %v", hr.apiToken))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Error fetching highlights: %v", err.Error())
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching highlights - response status code is %v", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error resding reponse body: %v", err.Error())
	}

	var highlights Highlights
	json.Unmarshal(responseBody, &highlights)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling results: %v\n", err.Error())
	}
	return &highlights, nil
}
