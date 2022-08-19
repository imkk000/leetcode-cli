package info

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/client"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"rogchap.com/v8go"
)

func GetSubmission(id string) error {
	path := fmt.Sprintf("/submissions/detail/%s/", id)
	resp, err := client.Get(path)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	doc.Find("script").Each(func(_ int, s *goquery.Selection) {
		if !strings.Contains(s.Text(), "submissionData") {
			return
		}
		ctx := v8go.NewContext()
		if _, err := ctx.RunScript(s.Text(), "main.js"); err != nil {
			return
		}
		v, err := ctx.Global().Get("pageData")
		if err != nil {
			return
		}
		if !v.IsObject() {
			return
		}
		data, err := v.MarshalJSON()
		if err != nil {
			return
		}

		var sub map[string]json.RawMessage
		if err := json.Unmarshal(data, &sub); err != nil {
			return
		}
		data, err = json.MarshalIndent(sub, "", "  ")
		if err != nil {
			return
		}
		fmt.Println(string(data))
	})
	return nil
}
