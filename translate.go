package main

import (
	"context"
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/olekukonko/tablewriter"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type request struct {
	target string            // target language
	text   string            // text to translate
	codes  map[string]string // Language code map
}

const usage string = `Usage: translate [OPTIONS]

Simple app for translating text via Google Translate API

Options:
  -l, --language-code=""      Desired ouput Language Code
  -t, --text=""               The text to be translated
  --codes                     Lists some available language codes
`

func (r *request) getCodes() {
	// Sets code property with map of available language codes
	m := map[string]string{
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

	// set codes property in struct
	r.codes = m
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

func (r *request) translateText() (string, error) {
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

func main() {
	// Defining CLI options in main function
	args := make(map[string]*string)
	app := cli.App("translate", "Simple app for translating text via Google Translate API")
	args["language-code"] = app.StringOpt("l language-code", "", `Desired ouput Language Code`)
	args["text"] = app.StringOpt("t text", "", `The text to be translated`)
	codesFlag := app.BoolOpt("codes", false, "Lists some available language codes")

	app.Cmd.Action = func() {

		if *codesFlag {
			req := request{}
			req.printCodes()
			return
		}

		if *args["language-code"] == "" && *args["text"] == "" {
			fmt.Println(usage)
			return
		}

		req := request{
			target: *args["language-code"],
			text:   *args["text"],
		}

		for k, v := range args {
			if *v == "" {
				fmt.Printf(`Please specify "%s"`+"\n", k)
				return
			}
		}

		resp, err := req.translateText()
		if err != nil {
			fmt.Println("ERROR:", err.Error())
		} else {
			fmt.Println(resp)
		}

	}

	app.Run(os.Args)
}
