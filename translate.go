package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/translate"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/language"
)

type request struct {
	target string            // target language
	text   string            // text to translate
	codes  map[string]string // Language code map
}

func (r *request) getCodes() {
	// Sets code property with map of available language codes

	// set codes property in struct
	r.codes = map[string]string{
		"Spanish":  "es",
		"Greek":    "el",
		"Japanese": "ja",
		"Catalan":  "ca",
		"German":   "de",
		"Dutch":    "nl",
		"French":   "fr",
		"Italian":  "it",
		"Swedish":  "sv",
	}
}

func (r *request) printCodes() {
	// Print codes in a table
	r.getCodes()

	// Define table
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Language", "Code"})
	t.SetBorder(false)

	// add items
	for k, v := range r.codes {
		t.Append([]string{k, v})
	}
	fmt.Printf("\n")
	t.Render()
	fmt.Printf("\n" + `See https://ctrlq.org/code/19899-google-translate-languages for details on other available codes` + "\n\n")
}

func (r *request) Translate() (string, error) {
	// Returns translate text string via Google Translate API
	ctx := context.Background()

	lang, err := language.Parse(r.target)
	if err != nil {
		return "", err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}

	defer client.Close()

	resp, err := client.Translate(ctx, []string{r.text}, lang, nil)
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}
