package code

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenFile(t *testing.T) {
	expectation := strings.TrimSpace(`
package leetcode_test

/*
metadata:
  id: "108"
  title: convert-sorted-array-to-binary-search-tree
  lang: golang
  type: large
  inputString: |-
    [-10,-3,0,5,9]
    [1,3]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

}

`)

	data, err := GenerateFile(GenerateFileConfig{
		Code: Code{
			Id:          "108",
			Title:       "convert-sorted-array-to-binary-search-tree",
			Language:    "golang",
			Type:        "large",
			Input:       []string{},
			InputString: "[-10,-3,0,5,9]\n[1,3]",
			RawInput:    "",
			Src:         "/**\n * Definition for a binary tree node.\n * type TreeNode struct {\n *     Val int\n *     Left *TreeNode\n *     Right *TreeNode\n * }\n */\nfunc sortedArrayToBST(nums []int) *TreeNode {\n    \n}",
		},
	})

	if assert.NoError(t, err) {
		assert.Equal(t, expectation, strings.TrimSpace(string(data)))
	}
}
