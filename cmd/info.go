package cmd

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/leetcode/info"
	"leetcode-tool/leetcode/strparser"

	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		switch getType(args[0]) {
		case GetTypeInfo:
			result, err := info.GetProblemDetail(strparser.ParseUrl(args[1]))
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
		case GetTypeSubmission:
			result, err := info.GetSubmission(args[1])
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
		default:
			fmt.Printf("type %s not found", args[0])
		}
	},
}

type GetType int

const (
	GetTypeInfo GetType = iota + 1
	GetTypeSubmission
)

func getType(t string) GetType {
	switch t {
	case "info":
		return GetTypeInfo
	case "submission":
		return GetTypeSubmission
	}
	return GetType(0)
}
