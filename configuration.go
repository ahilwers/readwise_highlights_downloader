package main

import (
	"flag"
	"os"

	"github.com/peterbourgon/ff"
)

type Configuration struct {
	apiToken      string
	outputDir     string
	lastFetchDate string
}

func GetConfiguration() (Configuration, error) {
	fs := flag.NewFlagSet("readwise_highlights_downloader", flag.ContinueOnError)
	var (
		apiToken      = fs.String("api-token", "", "")
		outputDir     = fs.String("output-directory", "", "")
		lastFetchDate = fs.String("last_fetch_date", "", "")
		_             = fs.String("config", "./readwise_highlights_downloader.conf", "config file (optional)")
	)

	ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("READWISE_HIGHLIGHTS_DOWNLOADER"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	)

	var configuration Configuration
	configuration.apiToken = *apiToken
	configuration.outputDir = *outputDir
	configuration.lastFetchDate = *lastFetchDate
	return configuration, nil
}
