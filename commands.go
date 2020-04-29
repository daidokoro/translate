package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	codes    bool   // used as language code listing flag
	langCode string // used as language code selection flag
)

// Cobra Init function.
func init() {
	rootCmd.Flags().BoolVarP(&codes, "codes", "C", false, "List of available Language Codes")
	rootCmd.Flags().StringVarP(&langCode, "language-code", "c", "en", "Select language code")
}

// root command
var rootCmd = &cobra.Command{
	Use:   `translate -c [Lang Code] [Text]`,
	Short: "Simple app for translating text via Google Translate API ",
	Run: func(cmd *cobra.Command, args []string) {

		if codes {
			req := request{}
			req.printCodes()
			return
		}

		if len(args) < 1 {
			cmd.Help()
			return
		}

		t := args[0]

		req := &request{
			target: langCode,
			text:   t,
		}

		resp, err := req.Translate()
		if err != nil {
			fmt.Println("ERROR:", err.Error())
		} else {
			fmt.Println(resp)
		}

	},
}
