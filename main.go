package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("コマンドライン引数が足りない")
	}
	for i := 1; i < len(os.Args); i++ {
		f, err := os.Open(os.Args[i])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// 取得してきた文字列をパースメソッドにわたす
		strings, err := lex(f)
		if err != nil {
			log.Fatal(err)
		}

		parseStrings(strings)

	}
}
