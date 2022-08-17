package run

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"leetcode-tool/client"
	"leetcode-tool/code"
	"strconv"
	"time"
)

func RunCode(c *code.Code) (*ResponseRunResultBody, error) {
	path := fmt.Sprintf("problems/%s/interpret_solution/", c.Title)
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	resp, err := client.Post(path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var respBody ResponseRunResultBody
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}

func SubmitCode(c *code.Code) (*ResponseSubmissionResultBody, error) {
	path := fmt.Sprintf("/problems/%s/submit/", c.Title)
	reqBody := RequestSubmissionBody{
		QuestionId: c.Id,
		Lang:       c.Language,
		TypedCode:  c.Src,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Post(path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var respBody ResponseSubmissionResultBody
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}

func FetchSubmissionResult(id int64) (*ResponseFetchSubmissionBody, error) {
	submissionId := strconv.FormatInt(id, 10)
	status := new(ResponseFetchSubmissionBody)
	if err := fetchResult(submissionId, status); err != nil {
		return nil, err
	}
	return status, nil
}

func FetchRunResult(id string) (*ResponseFetchRunResultBody, error) {
	status := new(ResponseFetchRunResultBody)
	if err := fetchResult(id, status); err != nil {
		return nil, err
	}
	return status, nil
}

func fetchResult(id string, v any) error {
	t := time.NewTicker(250 * time.Millisecond)
	defer t.Stop()

	path := fmt.Sprintf("submissions/detail/%s/check/", id)
	var count int
	for range t.C {
		count++

		switch v.(type) {
		case *ResponseFetchRunResultBody:
			status := v.(*ResponseFetchRunResultBody)
			if status.State == FetchResultStateSuccess {
				return nil
			}
		case *ResponseFetchSubmissionBody:
			status := v.(*ResponseFetchSubmissionBody)
			if status.State == FetchResultStateSuccess {
				return nil
			}
		}
		if count >= FetchResultRetryLimit {
			return errors.New("reach maximum retry limit")
		}

		resp, err := client.Get(path)
		if err != nil {
			return err
		}
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
		if err := resp.Body.Close(); err != nil {
			return err
		}
	}
	return nil
}
