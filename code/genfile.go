package code

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v3"
)

type GenerateFileConfig struct {
	Filename string
	Code
}

func GenerateFile(cfg GenerateFileConfig) ([]byte, error) {
	metadata := struct {
		Code `yaml:"metadata"`
	}{
		Code: cfg.Code,
	}
	b := new(bytes.Buffer)
	enc := yaml.NewEncoder(b)
	enc.SetIndent(2)
	if err := enc.Encode(metadata); err != nil {
		return nil, err
	}
	if err := enc.Close(); err != nil {
		return nil, err
	}

	var sb strings.Builder
	sb.WriteString("package leetcode_test\n")
	sb.WriteString("/*\n")
	sb.WriteString(b.String())
	sb.WriteString("*/\n\n")
	sb.WriteString(cfg.Src)

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", sb.String(), parser.ParseComments)
	if err != nil {
		return nil, err
	}
	b.Reset()
	if err := printer.Fprint(b, fs, f); err != nil {
		return nil, err
	}
	// import
	return imports.Process("", b.Bytes(), nil)
}
