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
	Short: "run and submit problem",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := code.ParseFilename(args[0])
		if err != nil {
			fmt.Printf("parse %s: %v\n", args[0], err)
			return
		}
		kind, err := cmd.Flags().GetString("kind")
		if err != nil {
			fmt.Println("get kind: ", err)
			return
		}
		var result any
		switch kind {
		case "run":
			id, err := run.RunCode(c)
			if err != nil {
				fmt.Println("run code: ", err)
				return
			}
			fetchResult, err := run.FetchRunResult(id.InterpretId)
			if err != nil {
				fmt.Println("fetch result: ", err)
				return
			}
			result = fetchResult
		case "submit":
			id, err := run.SubmitCode(c)
			if err != nil {
				fmt.Println("submit code: ", err)
				return
			}
			fetchResult, err := run.FetchSubmissionResult(id.SubmissionId)
			if err != nil {
				fmt.Println("fetch result: ", err)
				return
			}
			result = fetchResult
		}

		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Println("json marshal: ", err)
			return
		}
		fmt.Printf("resut: %s\n", data)
	},
}

func prepareRunFlags() {
	RunCmd.Flags().StringP("kind", "k", "run", "set kind run or submit")
}
