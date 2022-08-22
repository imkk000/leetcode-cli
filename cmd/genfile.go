package cmd

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/code"
	"leetcode-tool/leetcode/info"
	"leetcode-tool/leetcode/strparser"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var GenerateFileCmd = &cobra.Command{
	Use:   "genfile",
	Short: "Generate Test File",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println("get path:", err)
			return
		}
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			fmt.Println("get filename:", err)
			return
		}
		write, err := cmd.Flags().GetBool("write")
		if err != nil {
			fmt.Println("get write:", err)
			return
		}

		result, err := info.GetProblemDetail(strparser.ParseUrl(args[0]), true)
		if err != nil {
			fmt.Println("inquire:", err)
			return
		}
		if _, exists := result["data"]; !exists {
			data, err := json.Marshal(result)
			if err != nil {
				fmt.Println("json marshal:", err)
				return
			}
			fmt.Println(string(data))
			return
		}

		// fix golang only
		lang := "golang"

		var data struct {
			Question struct {
				QuestionId       string `json:"questionId"`
				TitleSlug        string `json:"titleSlug"`
				JudgeType        string `json:"judgeType"`
				ExampleTestcases string `json:"exampleTestcases"`
				CodeSnippets     []struct {
					Lang string `json:"langSlug"`
					Src  string `json:"code"`
				} `json:"codeSnippets"`
			} `json:"question"`
		}
		if err := json.Unmarshal(result["data"], &data); err != nil {
			fmt.Println("json unmarshal:", err)
			return
		}
		var src string
		for _, cs := range data.Question.CodeSnippets {
			if cs.Lang == lang {
				src = cs.Src
				break
			}
		}
		if src == "" {
			fmt.Printf("code lang %s not found\n", lang)
			return
		}
		if filename == "" {
			filename = fmt.Sprintf("%s_test.go", data.Question.QuestionId)
		}

		cfg := code.GenerateFileConfig{
			Code: code.Code{
				Id:          data.Question.QuestionId,
				Title:       data.Question.TitleSlug,
				Language:    lang,
				Type:        data.Question.JudgeType,
				InputString: data.Question.ExampleTestcases,
				Src:         src,
			},
		}
		content, err := code.GenerateFile(cfg)
		if err != nil {
			fmt.Println("generate file:", err)
			return
		}
		fmt.Printf("output:\n\n")
		fmt.Println(string(content))

		if write {
			filename = filepath.Join(path, filename)
			if err := os.WriteFile(filename, content, 0644); err != nil {
				fmt.Println("write to file:", err)
				return
			}
		}
	},
}

func prepareGenerateFileFlags() {
	GenerateFileCmd.Flags().StringP("filename", "f", "", "set output filename")
	GenerateFileCmd.Flags().StringP("path", "p", "", "set output path")
	GenerateFileCmd.Flags().BoolP("write", "w", true, "set write output directly to file")
}
