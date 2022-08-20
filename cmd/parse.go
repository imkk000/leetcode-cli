package cmd

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/code"

	"github.com/spf13/cobra"
)

var ParseFileCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse code",
	Run: func(_ *cobra.Command, args []string) {
		c, err := code.ParseFilename(args[0])
		if err != nil {
			fmt.Printf("parse %s: %v\n", args[0], err)
			return
		}

		data, err := json.Marshal(c)
		if err != nil {
			fmt.Println("json marshal:", err)
			return
		}
		fmt.Println(string(data))
	},
}
