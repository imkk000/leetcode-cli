package cmd

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/leetcode/info"
	"leetcode-tool/leetcode/strparser"

	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		result, err := info.GetProblemDetail(strparser.ParseUrl(args[0]))
		if err != nil {
			fmt.Println("inquire:", err)
			return
		}
		data, err := json.Marshal(result)
		if err != nil {
			fmt.Println("json marshal:", err)
			return
		}
		fmt.Println(string(data))
	},
}
