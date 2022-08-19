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

const (
	FetchResultStateSuccess = "\"SUCCESS\""
	FetchResultRetryLimit   = 150
)

func RunCode(c *code.Code) (string, error) {
	path := fmt.Sprintf("problems/%s/interpret_solution/", c.Title)
	body, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	resp, err := client.Post(path, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	var respBody struct {
		InterpretId string `json:"interpret_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}
	return respBody.InterpretId, nil
}

func SubmitCode(c *code.Code) (string, error) {
	path := fmt.Sprintf("/problems/%s/submit/", c.Title)
	reqBody := struct {
		QuestionId string `json:"question_id"`
		Lang       string `json:"lang"`
		TypedCode  string `json:"typed_code"`
	}{
		QuestionId: c.Id,
		Lang:       c.Language,
		TypedCode:  c.Src,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	resp, err := client.Post(path, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	var respBody struct {
		SubmissionId int64 `json:"submission_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}
	return strconv.FormatInt(respBody.SubmissionId, 10), nil
}

func FetchRunResult(id string) (map[string]json.RawMessage, error) {
	t := time.NewTicker(250 * time.Millisecond)
	defer t.Stop()

	path := fmt.Sprintf("submissions/detail/%s/check/", id)
	var count int
	var v map[string]json.RawMessage
	for range t.C {
		count++

		if count >= FetchResultRetryLimit {
			return nil, errors.New("reach maximum retry limit")
		}

		resp, err := client.Get(path)
		if err != nil {
			return nil, err
		}
		if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return nil, err
		}
		if err := resp.Body.Close(); err != nil {
			return nil, err
		}

		state, ok := v["state"]
		if !ok {
			return nil, errors.New("state not found in response")
		}
		if string(state) == FetchResultStateSuccess {
			break
		}
	}
	return v, nil
}
