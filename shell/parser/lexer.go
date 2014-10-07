package parser

//	字句解析

import (
	"errors"
	"fmt"
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/common/token"
)

type Scanner interface {
	Scan() (tok int, lit string, pos token.Position)
}

//	以下、解析器の汎用処理
type Lexer struct {
	s     Scanner
	lit   string
	pos   token.Position
	stmts []ast.Stmt
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	lexLog.Info("	lexer:", "[", lit, "]", pos)
	if tok == -1 {
		return 0 //	0 を返すと字句解析終了
	}
	if tok == 0 {
		lexLog.Fatal(fmt.Sprintf("Line %d, Column %d: %s\n", pos.Line, pos.Column, lit))
		return tok
	}
	lval.tok = token.Token{Tok: tok, Lit: lit, Pos: pos}
	l.lit = lit
	l.pos = pos
	return tok
}

//	エラー
func (l *Lexer) Error(e string) {
	lexLog.Fatal(fmt.Sprintf("Line %d, Column %d: %q %s\n", l.pos.Line, l.pos.Column, l.lit, e))
}

//	パース
func ParseScanner(s Scanner) ([]ast.Stmt, error) {
	//^[[Aなどの入力で
	//panic: runtime error: invalid memory address or nil pointer dereference
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Fprintln(*s.ios.Out, "Characters that are unexpected:")
	//	}
	//}()

	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, errors.New("Parse error")
	}
	return l.stmts, nil
}
