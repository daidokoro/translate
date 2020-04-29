# Translate

Translate is a simple compiled cli utility for translating text using the Google Apps API written in Go.

## Requirements

- Google Apps/Cloud Account
- Google Default Credentials Configured - See [Here](https://developers.google.com/identity/protocols/application-default-credentials) for details.


## Install
    $ go get github.com/daidokoro/translate


## Usage

```sh

# make sure your google credentials are set
$ GOOGLE_APPLICATION_CREDENTIALS=/path/to/credential.json

$ translate

Simple app for translating text via Google Translate API

Usage:
  translate -c [Lang Code] [Text] [flags]

Flags:
  -C, --codes                  List of available Language Codes
  -h, --help                   help for translate
  -c, --language-code string   Select language code (default "en")

# translate something
$ translate -c en hola
hello

```

## Install

```
go get -v -u github.com/daidokoro/translate
```
