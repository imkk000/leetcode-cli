package code

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type (
	InputKind int
)

const (
	KindInputInt InputKind = iota
	KindInputString
	KindInputStringSlice
	KindInputIntSlice
)

func (i InputKind) String() string {
	return [...]string{
		"int",
		"string",
		"stringSlice",
		"intSlice",
	}[i]
}

func MakeInputs(size int, v []any) [][]any {
	g := make([][]any, 0)
	vl, pl := len(v), size
	for i := 0; i < vl; i += pl {
		s := make([]any, pl)
		for j := 0; j < pl; j++ {
			s[j] = v[i+j]
		}
		g = append(g, s)
	}
	return g
}

func ParseInput(s string) ([]any, error) {
	if len(s) == 0 {
		return make([]any, 0), errors.New("empty input")
	}

	lines := strings.Split(strings.TrimSpace(s), "\n")
	if len(lines) == 0 {
		return make([]any, 0), errors.New("empty newline")
	}

	m := make([]any, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)

		switch typeOf(line) {
		case KindInputString:
			m = append(m, line[1:len(line)-1])
		case KindInputInt:
			v, err := strconv.Atoi(line)
			if err != nil {
				return make([]any, 0), fmt.Errorf("convert %s", line)
			}
			m = append(m, v)
		case KindInputIntSlice:
			m = append(m, makeIntSlice(line))
		case KindInputStringSlice:
			m = append(m, makeStringSlice(line))
		}
	}
	return m, nil
}

func makeIntSlice(s string) []int {
	elms := strings.Split(s[1:len(s)-1], ",")
	m := make([]int, 0)
	for _, e := range elms {
		m = append(m, getInt(e))
	}
	return m
}

func makeStringSlice(s string) []string {
	elms := strings.Split(s[1:len(s)-1], ",")
	m := make([]string, 0)
	for _, e := range elms {
		m = append(m, e[1:len(e)-1])
	}
	return m
}

func typeOf(s string) InputKind {
	if isInt(s) {
		return KindInputInt
	}
	if isString(s) {
		return KindInputString
	}
	if isStringSlice(s) {
		return KindInputStringSlice
	}
	if isIntSlice(s) {
		return KindInputIntSlice
	}
	return -1
}

func isIntSlice(s string) bool {
	return matchString(`^\[\d(.+)\d\]$`, s)
}

func isString(s string) bool {
	return matchString(`^"(.+)"$`, s)
}

func isStringSlice(s string) bool {
	return matchString(`^\[\"(.+)\"\]$`, s)
}

func isInt(s string) bool {
	return matchString(`^\d+$`, s)
}

func getInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func matchString(pattern, s string) bool {
	ok, err := regexp.MatchString(pattern, s)
	return ok && err == nil
}
