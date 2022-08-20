package cmd

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/code"
	"leetcode-tool/leetcode/run"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run and submit problem",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := code.ParseFilename(args[0])
		if err != nil {
			fmt.Printf("parse %s: %v\n", args[0], err)
			return
		}
		kind, err := cmd.Flags().GetString("kind")
		if err != nil {
			fmt.Println("get kind:", err)
			return
		}
		var subId string
		switch kind {
		case "run":
			subId, err = run.RunCode(c)
			if err != nil {
				fmt.Println("run code:", err)
				return
			}
		case "submit":
			subId, err = run.SubmitCode(c)
			if err != nil {
				fmt.Println("submit code:", err)
				return
			}
		}

		result, err := run.FetchRunResult(subId)
		if err != nil {
			fmt.Println("fetch result:", err)
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

func prepareRunFlags() {
	RunCmd.Flags().StringP("kind", "k", "run", "set kind run or submit")
}
