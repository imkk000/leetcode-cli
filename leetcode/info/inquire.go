package info

import (
	"encoding/json"
	"leetcode-tool/client"
	"leetcode-tool/client/graphql"
)

func Get(titleSlug string) (*Result, error) {
	variables := graphql.V{
		"titleSlug": titleSlug,
	}
	reqBody, err := graphql.CreateRequest(Query, variables)
	if err != nil {
		return nil, err
	}
	resp, err := client.Post("graphql", reqBody)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
