package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/flytam/filenamify"
)

type BookHighlightsCreator interface {
	CreateBookHighlightsFile() error
}

type bookHighlightsCreator struct {
	book            Book
	outputDirectory string
}

func NewBookHighlightsCreator(book Book, outputDirectory string) BookHighlightsCreator {
	return &bookHighlightsCreator{
		book:            book,
		outputDirectory: outputDirectory,
	}
}

func (hc *bookHighlightsCreator) CreateBookHighlightsFile() error {
	filename, err := hc.getValidFileName(hc.book.ReadableTitle)
	if err != nil {
		return fmt.Errorf("could not create valid file name from title %v: %v", hc.book.Title, err.Error())
	}
	path := filepath.Join(hc.outputDirectory, filename)
	fmt.Printf("%v -> %v\n", hc.book.ReadableTitle, path)
	fileExists := hc.doesFileExist(path)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("error creating file %v: %v", path, err.Error())
	}
	defer file.Close()
	if !fileExists {
		err = hc.writeFileContents(file)
	} else {
		err = hc.appendHighlights(file)
	}
	if err != nil {
		return err
	}
	return nil
}

func (hc *bookHighlightsCreator) doesFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func (hc *bookHighlightsCreator) getValidFileName(bookTitle string) (string, error) {
	filename, err := filenamify.Filenamify(bookTitle, filenamify.Options{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v.md", filename), nil
}

func (hc *bookHighlightsCreator) writeFileContents(file *os.File) error {
	err := hc.writeTitle(file)
	if err != nil {
		return err
	}
	err = hc.writeMetadata(file)
	if err != nil {
		return err
	}
	err = hc.writeHighlights(file)
	if err != nil {
		return err
	}
	return nil
}

func (hc *bookHighlightsCreator) writeTitle(file *os.File) error {
	_, err := file.WriteString(fmt.Sprintf("# %v\n\n", hc.book.ReadableTitle))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf(" ![](%v)\n\n", hc.book.CoverImageURL))
	if err != nil {
		return err
	}
	return nil
}

func (hc *bookHighlightsCreator) writeMetadata(file *os.File) error {
	_, err := file.WriteString("### Metadata\n\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("* Author: %v\n", hc.book.Author))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("* Full Title: %v\n", hc.book.Title))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("* Category: %v\n\n", hc.book.Category))
	if err != nil {
		return err
	}
	return nil
}

func (hc *bookHighlightsCreator) writeHighlights(file *os.File) error {
	_, err := file.WriteString("### Highlights\n\n")
	if err != nil {
		return err
	}
	err = hc.appendHighlights(file)
	if err != nil {
		return err
	}
	return nil
}

func (hc *bookHighlightsCreator) appendHighlights(file *os.File) error {
	for _, highlight := range hc.book.Highlights {
		location := fmt.Sprintf("[Location %v](%v)", highlight.Location, highlight.ReadwiseURL)
		_, err := file.WriteString(fmt.Sprintf("* %v (%v)\n", highlight.Text, location))
		if err != nil {
			return err
		}
	}
	return nil
}
