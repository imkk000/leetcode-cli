package graphql

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequest(t *testing.T) {
	expectation := []byte("{\"query\":\"query questionData($titleSlug: String!) {}\",\"variables\":{\"titleSlug\":\"test\"}}\n")
	query := "query questionData($titleSlug: String!) {}"
	variables := V{
		"titleSlug": "test",
	}

	reqBody, err := CreateRequest(query, variables)
	actual := reqBody.(*bytes.Buffer)

	if assert.NoError(t, err) {
		assert.Equal(t, expectation, actual.Bytes())
	}
}
