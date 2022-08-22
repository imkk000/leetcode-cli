package info

import (
	"encoding/json"
	"fmt"
	"leetcode-tool/client"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"rogchap.com/v8go"
)

func GetSubmission(id string) (map[string]json.RawMessage, error) {
	path := fmt.Sprintf("/submissions/detail/%s/", id)
	resp, err := client.Get(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]json.RawMessage
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

		if err := json.Unmarshal(data, &result); err != nil {
			return
		}
	})
	return result, nil
}
