package cmd

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/leetcode/info"

	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "inquire problem",
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		result, err := info.Get(args[0])
		if err != nil {
			fmt.Println("inquire: ", err)
			return
		}

		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Println("json marshal: ", err)
			return
		}
		fmt.Printf("result: %s\n", data)
	},
}
