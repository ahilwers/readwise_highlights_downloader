package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/peterbourgon/ff"
)

type Configuration struct {
	apiToken  string
	outputDir string
}

func GetConfiguration() (Configuration, error) {
	configDir, err := os.UserHomeDir()
	if err != nil {
		configDir = "./"
	}
	configPath := filepath.Join(configDir, ".readwise_highlights_downloader")

	fs := flag.NewFlagSet("readwise_highlights_downloader", flag.ContinueOnError)
	var (
		apiToken  = fs.String("api-token", "", "The API token can be obtained from https://readwise.io/access_token")
		outputDir = fs.String("output-directory", "", "The directory where the markdown files will be written")
		_         = fs.String("config", configPath, "config file (optional)")
	)

	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("READWISE_HIGHLIGHTS_DOWNLOADER"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)

	var configuration Configuration
	configuration.apiToken = *apiToken
	configuration.outputDir = *outputDir
	return configuration, nil
}
