package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

// ASTに分解して、取得できた文字列をスライスにいれて返すメソッド
func lex(src io.Reader) ([]string, error) {
	var (
		node *ast.File
		err  error
	)
	fset := token.NewFileSet()

	if file, ok := src.(*os.File); ok {
		node, err = parser.ParseFile(fset, file.Name(), nil, parser.Mode(0))
		if err != nil {
			return nil, err
		}
	} else {
		dst := new(bytes.Buffer)
		io.Copy(dst, src)
		node, err = parser.ParseFile(fset, "file.go", dst, parser.Mode(0))
		if err != nil {
			return nil, err
		}
	}

	var strings []string
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.BasicLit:
			strings = append(strings, x.Value)
		}
		return true
	})
	return strings, nil
}
