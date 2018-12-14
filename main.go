package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	if len(os.Args) > 2 {
		log.Fatal("コマンドライン引数が多すぎる")
	}

	// コマンドライン引数に指定したファイル名のスキャナーを作成
	scanner, err := getScanner(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	parse(scanner)
}

func getScanner(name string) (*bufio.Scanner, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return bufio.NewScanner(file), nil
}
