package code

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	"gopkg.in/yaml.v3"
)

type Code struct {
	Id          string   `json:"question_id" yaml:"id"`
	Title       string   `json:"-" yaml:"title"`
	Language    string   `json:"lang" yaml:"lang"`
	Type        string   `json:"judge_type" yaml:"type"`
	Input       []string `json:"-" yaml:"input"`
	InputString string   `json:"-" yaml:"inputString"`
	RawInput    string   `json:"data_input" yaml:"-"`
	Src         string   `json:"typed_code" yaml:"-"`
}

func ParseFilename(filename string) (*Code, error) {
	f, err := parseFile(filename)
	if err != nil {
		return nil, err
	}
	c, err := readMetadata(f.Comments)
	if err != nil {
		return nil, err
	}
	decls, err := readDecls(f.Decls)
	if err != nil {
		return nil, err
	}

	// new file
	c.Src, err = generateSource(&ast.File{
		Name:  ast.NewIdent("code"),
		Decls: decls,
	})

	return c, nil
}

func generateSource(f *ast.File) (string, error) {
	b := new(strings.Builder)
	if err := printer.Fprint(b, token.NewFileSet(), f); err != nil {
		return "", err
	}

	lines := strings.Split(b.String(), "\n")
	b.Reset()

	// strip package header
	for _, line := range lines[1:] {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		line = strings.ReplaceAll(line, "\t", "  ")
		b.WriteString(line)
		b.WriteRune('\n')
	}
	return b.String(), nil
}

func readMetadata(comments []*ast.CommentGroup) (*Code, error) {
	for _, comment := range comments {
		c := strings.TrimSpace(comment.Text())
		if strings.Contains(c, "metadata") {
			var metadata struct {
				Code `yaml:"metadata"`
			}
			if err := yaml.Unmarshal([]byte(c), &metadata); err != nil {
				return nil, err
			}
			if metadata.InputString != "" {
				metadata.Input = append(metadata.Input, metadata.InputString)
			}
			metadata.RawInput = strings.Join(metadata.Input, "\n")
			return &metadata.Code, nil
		}
	}
	return nil, errors.New("metadata not found")
}

func readDecls(decls []ast.Decl) ([]ast.Decl, error) {
	l := len(decls)
	for _, d := range decls {
		switch d.(type) {
		case *ast.GenDecl:
			g := d.(*ast.GenDecl)
			if isExclude(g.Doc.Text()) || isImport(g.Tok.String()) {
				continue
			}
		case *ast.FuncDecl:
			g := d.(*ast.FuncDecl)
			if isExclude(g.Doc.Text()) {
				continue
			}
		}
		decls = append(decls, d)
	}
	return decls[l:], nil
}

func isImport(s string) bool {
	return strings.Contains(s, "import")
}

func isExclude(s string) bool {
	return strings.Contains(strings.TrimSpace(s), "go:exclude")
}

func parseFile(filename ...string) (*ast.File, error) {
	if filename == nil {
		return &ast.File{
			Name: ast.NewIdent("code"),
		}, nil
	}
	return parser.ParseFile(token.NewFileSet(), filename[0], nil, parser.ParseComments)
}
