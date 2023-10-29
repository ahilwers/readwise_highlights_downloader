# Readwise Highlights Downloader

A simple command line tool that downloads the highlights from readwise and creates a Markdown file for each book in a specified directory.

On it's first run all highlights will be downloaded from readwise.io and the time is stored in the file `.readwise_highlights_downloader_lastupdate`. On the next run this time is used to only fetch the highlights that were added since the last run.

If you want to get all highlights anyway you can just delete the file `.readwise_highlights_downloader_lastupdate`

## Building

To build the application you can just run `go build` in the main directory.

## Running

To run this application you need to specify the API token and the output directory for the Markdown files. These can either be specified via parameters or you can create a configuration file for them.

### Parameters

You can specify the configuration by passing parameters like so:

```
./readwise_highlights_downloader -api-token [MyApiToken] -output-directory [MyOutputDirectory]
```

### Configuration File

The more convenient way to pass the configuration is by creating a configuration file. By default the application looks for it's configuration in the file `.readwise_highlights_downloader` in the user's home directory. You can specify a different configuration file by using the `-config` parameter.

The file should look like this:

```
api-token [MyApiToken]
output-directory [MyOutputDirectory]
```

### API Token

You can obtain an API token from readwise.io by opening the url https://readwise.io/access_token.








