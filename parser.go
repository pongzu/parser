package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"

	s "github.com/derekparker/delve/service/test"
)

// Token は意味のある文字を表す
type Token int

const (
	ILLEGALToken = iota
	EOF
	WS       // 空白
	IDENT    //Fields,TableNameなど
	ASTERISK //*
	COMMA    //,
	SELECT   // SELECT key word
	FROM     // FROM key word
)

type Parser struct {
	r *bufio.Reader
}

func NewParser(src io.Reader) *Parser {
	return &Parser{r: bufio.NewReader(src)}
}

func (p *Parser) read() rune {
	eof := rune(0)
	ch, _, err := p.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (p *Parser) unread() error {
	if err := p.r.UnreadRune(); err != nil {
		return err
	}
	return nil
}

func (p *Parser) scan() (tok Token, lit string) {
	ch := p.read()

	switch {
	case isWhiteSpace(ch):
		p.unread()
		return p.scanWhiteSpace()
	case isLetter(ch):
		return p.ScanIdent()
	case ch == '*':
		return ASTERISK, string(ch)
	case ch == ',':
		return COMMA, string(ch)
	case ch == '0':
		return EOF, string(ch)
	}
	return ILLEGALToken, string(ch)
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == "\n"
}

func (p *Parser) scanWhiteSpace() (tok Token, lit string) {
	var buf bytes.Buffer

	// ここで発生するえらーはそのままログに出しておいたらいいレベル
	_, err := buf.WriteRune(p.read())
	if err != nil {
		log.Println(err)
	}

	for {
		ch := p.read()

		if ch == '0' {
			break
		} else if !isWhiteSpace(ch) {
			p.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func isLetter(ch rune) bool {
	return (ch > 'a' && ch < 'z') || (ch > 'A' && ch < Z) || (ch > 'あ' && ch < 'ん')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func (p *Parser) ScanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	ch := p.read()
	_, err := buf.WriteRune(ch)
	if err != nil {
		log.Println(err)
	}

	for {
		ch := p.read()
		if ch == '0' {
			break
		} else if !isLetter(ch) && !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(r)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	}
	return IDENT, buf.String()
}
