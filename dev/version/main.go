package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// bump version
	filename := "config/version.go"
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("parse file:", err)
		return
	}

	for _, d := range f.Decls {
		switch d.(type) {
		case *ast.GenDecl:
			g := d.(*ast.GenDecl)
			for _, s := range g.Specs {
				switch s.(type) {
				case *ast.ValueSpec:
					v := s.(*ast.ValueSpec)
					for _, name := range v.Names {
						if name.Name == "Version" {
							b := v.Values[0].(*ast.BasicLit)
							b.Value = strings.Trim(b.Value, "\"")
							major, minor, patch := getVersion(b.Value)
							vs := []int{major, minor, patch}
							bumpPatch(vs)
							b.Value = getString(vs)
						}
					}
				}
			}
		}
	}

	b := new(bytes.Buffer)
	if err := printer.Fprint(b, fs, f); err != nil {
		fmt.Println("print file:", err)
		return
	}
	if err := os.WriteFile(filename, b.Bytes(), 0644); err != nil {
		fmt.Println("write file:", err)
		return
	}
}

func bumpMajor(v []int) {
	v[0], v[1], v[2] = v[0]+1, 0, 0
}

func bumpMinor(v []int) {
	v[1], v[2] = v[1]+1, 0
}

func bumpPatch(v []int) {
	v[2]++
}

func getVersion(v string) (major, minor, patch int) {
	// major.minor.patch
	re, err := regexp.Compile(`(\d+)\.(\d+)\.(\d+)`)
	if err != nil {
		return
	}
	s := re.FindStringSubmatch(v)
	if len(s) != 4 {
		return
	}
	s = s[1:]
	major, err = strconv.Atoi(s[0])
	if err != nil {
		return
	}
	minor, err = strconv.Atoi(s[1])
	if err != nil {
		return
	}
	patch, err = strconv.Atoi(s[2])
	if err != nil {
		return
	}
	return
}

func getString(v []int) string {
	if len(v) != 3 {
		return ""
	}
	return fmt.Sprintf(`"%d.%d.%d"`, v[0], v[1], v[2])
}
