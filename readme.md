# Leetcode Tools (LCC)

## Command

```text
Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        Get information
  run         Run and submit problem
  testnet     Test send HEAD to server
  version     Get application version

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

## Code Parser

### Source Code

```go
package example_test

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1464"
  title: "reduce-array-size-to-the-half"
  lang: "golang"
  type: "large"
  inputString: "[1,1]\n[2,2,3,4]"
  input:
    - "[3,3,3,3,5,5,5,2,2,7]"
    - "[7,7,7,7,7,7]"
*/

func minSetSize(nums []int) int {
    return 0
}

// go:exclude
const inputConst = "for_test"

// go:exclude
func generateInput() []int {
    return []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
}

func TestCode(t *testing.T) {
    tcs := []Tc{
        {
            expectation: 2,
            input:       []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
        },
        {
            expectation: 1,
            input:       []int{7, 7, 7, 7, 7, 7},
        },
    }

    for _, tc := range tcs {
        actual := minSetSize(tc.input.([]int))

        assert.Equal(t, tc.expectation, actual)
    }
}

func BenchmarkCode(b *testing.B) {
    for i := 0; i < b.N; i++ {
        minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7})
    }
}
```

### Result

```json
{
  "question_id": "1464",
  "lang": "golang",
  "judge_type": "large",
  "data_input": "[3,3,3,3,5,5,5,2,2,7]\n[7,7,7,7,7,7]\n[1,1]\n[2,2,3,4]",
  "typed_code": "func minSetSize(nums []int) int {\n  return 0\n}\n"
}
```

### Output Source Code

```go
func minSetSize(nums []int) int {
  return 0
}

```
