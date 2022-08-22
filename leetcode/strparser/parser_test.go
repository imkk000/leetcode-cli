package strparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUrl(t *testing.T) {
	tcs := []struct {
		expectation string
		url         string
	}{
		{
			expectation: "convert-sorted-array-to-binary-search-tree",
			url:         "https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/",
		},
		{
			expectation: "convert-sorted-array-to-binary-search-tree",
			url:         "https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree",
		},
		{
			expectation: "convert-sorted-array-to-binary-search-tree",
			url:         "convert-sorted-array-to-binary-search-tree/",
		},
		{
			expectation: "convert-sorted-array-to-binary-search-tree",
			url:         "convert-sorted-array-to-binary-search-tree",
		},
	}

	for _, tc := range tcs {
		actual := ParseUrl(tc.url)

		assert.Equal(t, tc.expectation, actual)
	}
}
