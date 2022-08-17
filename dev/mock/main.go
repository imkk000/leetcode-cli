package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

//go:embed json/inquire1.json
var inquireRawJSON1 []byte

//go:embed json/inquire2.json
var inquireRawJSON2 []byte

//go:embed json/inquire3.json
var inquireRawJSON3 []byte

func main() {
	http.HandleFunc("/problems/maximum-product-of-three-numbers/interpret_solution/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Write([]byte(`{"interpret_id":"1","test_case":"[1,2,3]\n[1,2,3,4]\n[-1,-2,-3]"}`))
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		b, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(b))

		var reqBody struct {
			Query     string         `json:"query"`
			Variables map[string]any `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte{})
			return
		}

		if strings.Contains(reqBody.Query, "questionData") {
			w.Header().Set("Content-Type", "application/json")
			switch reqBody.Variables["titleSlug"] {
			case "count-operations-to-obtain-zero":
				w.Write(inquireRawJSON1)
			case "design-a-text-editor":
				w.Write(inquireRawJSON2)
			case "maximum-product-of-three-numbers":
				w.Write(inquireRawJSON3)
			}
			return
		}
	})

	http.ListenAndServe(":9900", nil)
}
