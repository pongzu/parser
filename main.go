package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var s = "select * from sometable"
	stmt := &SelectStmt{}
	stmt.FormatSelectStmt(s)
}

type Scanner struct {
	r *bufio.Reader
}

// SelectStmt はセレクト文をトークンに分割して、スペースも排除された状態
type SelectStmt struct {
	tokens []Token
	fields []string
}

type Token int

var eof = rune(0)

const (
	ILLEGAL  = iota // some token that can not recognized
	EOF             // end of file
	WS              // 空白
	IDENT           //Fields,TableNameなど
	ASTERISK        //*
	COMMA           //,
	SELECT          // SELECT key word
	FROM            // FROM key word
)

func (stmt *SelectStmt) FormatSelectStmt(src string) {
	scanner := NewScanner(strings.NewReader(src))
	for {
		token, field := scanner.ScanIgnoreWhiteSpace()
		if token == EOF {
			break
		}
		stmt.tokens = append(stmt.tokens, token)
		stmt.fields = append(stmt.fields, field)
	}
	fmt.Println(stmt.format())
}

func (stmt *SelectStmt) format() string {
	buf := &bytes.Buffer{}

	for i, token := range stmt.tokens {
		switch token {
		case SELECT:
			buf.WriteString("SELECT\n")
		case IDENT:
			buf.WriteString("\t" + stmt.fields[i])
		case FROM:
			buf.WriteString("FROM\n")
		case ASTERISK:
			buf.WriteString("\t*\n")
		}
	}
	return buf.String()
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDisit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func NewScanner(src io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(src)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) Scan() (Token, string) {
	ch := s.read()
	if isWhiteSpace(ch) {
		s.unread()
		return s.ProcessWhiteSpace()
	} else if isLetter(ch) {
		s.unread()
		return s.ProcessIdent()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	}
	return ILLEGAL, string(ch)
}

func (s *Scanner) ScanIgnoreWhiteSpace() (Token, string) {
	tok, field := s.Scan()
	if tok == WS {
		tok, field = s.Scan()
	}
	return tok, field
}

func (s *Scanner) ProcessWhiteSpace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhiteSpace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func (s *Scanner) ProcessIdent() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDisit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
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
