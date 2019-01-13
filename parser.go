package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"
)

// Token は意味のある文字を表す
type Token int

const (
	ILLEGALToken = iota // some token that can not recognized
	EOF                 // end of file
	WS                  // 空白
	IDENT               //Fields,TableNameなど
	ASTERISK            //*
	COMMA               //,
	SELECT              // SELECT key word
	FROM                // FROM key word
)

// SelectStmt はセレクト文をトークンに分割して、スペースも排除された状態
type SelectStmt struct {
	tokens []int
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(src io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(src)}
}

func (stmt *SelectStmt) parseSelectStmt(src string) {
	scanner := NewScanner(strings.NewReader(src))
}

func (s *Scanner) scan() (tok Token, lit string) {
	ch := s.read()

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

func (s *Scanner) read() rune {
	eof := rune(0)
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() error {
	if err := p.r.UnreadRune(); err != nil {
		return err
	}
	return nil
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == "\n"
}

func (s *Scanner) scanWhiteSpace() (tok Token, lit string) {
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

func (s *Scanner) ScanIdent() (tok Token, lit string) {
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
