package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("コマンドライン引数が間違っている")
	} 

}

func run() error {

	//コマンラインで受け取った名前のファイルを開く
	f, err := getFile()
	if err != nil {
		return err
	}

	//ファイルをパーサーからパースする
	fs := token.NewFileSet()

	file, err := parser.ParseFile(fs, f.Name(), f, parser.Mode(0))
	if err != nil {
		return err
	}

	// 今から全てのノードの値を取得してきて、そこから文字列か判断して、文字列であれば、SQL文かどうか判断して、その場合は綺麗にする

	for _, v := range file.Decls {
		ast.Print(fs, v)
	}
	return nil
}


func getFile() (*os.File, error) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return nil, err
	}
	return file, nil
}
