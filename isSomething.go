package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// とりあえずいisSQLメソッドを定義して　isSELECT や isEXEX  などにも適応させる
func isSQL(src string) bool {
   if isSelect(src){
	   return true 
   }
}

func isSelect(src string)bool {
	parser := NewParser(src)
	parser.
}

