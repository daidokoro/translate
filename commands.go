package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var codes bool

// root command
var rootCmd = &cobra.Command{
	Use:   `translate ["Lang Code:Text"]`,
	Short: "Simple app for translating text via Google Translate API ",
	Run: func(cmd *cobra.Command, args []string) {

		if codes {
			req := request{}
			req.printCodes()
			return
		}

		if len(args) < 1 || !strings.Contains(args[0], ":") {
			cmd.Help()
			return
		}

		t := strings.Split(args[0], ":")

		req := &request{
			target: strings.TrimSpace(t[0]),
			text:   strings.TrimSpace(t[1]),
		}

		resp, err := req.Translate()
		if err != nil {
			fmt.Println("ERROR:", err.Error())
		} else {
			fmt.Println(resp)
		}

	},
}

// Cobra Init function.
func init() {
	rootCmd.Flags().BoolVarP(&codes, "codes", "c", false, "List of available Language Codes")
}
