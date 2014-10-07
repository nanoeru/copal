package scanner

import (
	"github.com/nanoeru/copal/shell/parser"
)

const (
	EOF = rune(0x19)
)

const (
	OUT_OF_EOF = -1
	ParseError = 0
)

//	走査器
type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

//	オペランドの定義
var opName = map[string]int{
	"==":  parser.EQ,
	"!=":  parser.NE,
	">=":  parser.GE,
	"<=":  parser.LE,
	":=":  parser.NEW_ASSIGN,
	"<-":  parser.CHAN,
	"and": parser.AND,
	"&&":  parser.AND,
	"or":  parser.OR,
	"||":  parser.OR,
	"xor": parser.XOR,
	"^":   parser.XOR,
	"+=":  parser.ADD_ASSIGN,
	"-=":  parser.SUB_ASSIGN,
	"*=":  parser.MUL_ASSIGN,
	"/=":  parser.DIV_ASSIGN,
	"%=":  parser.REM_ASSIGN,
	"^=":  parser.XOR_ASSIGN,
	"++":  parser.INC,
	"--":  parser.DEC,
	//	"...": parser.,
	"<<": parser.SHL,
	">>": parser.SHR,
}

//	ルーンオペランドの定義
var runeOpName = map[rune]int{
	'+':        0,
	'-':        0,
	'*':        0,
	'/':        0,
	'%':        0,
	'(':        0,
	'.':        0,
	'!':        0,
	'&':        0,
	'|':        0,
	'=':        0,
	'^':        0,
	'~':        0,
	':':        0,
	'?':        0,
	'<':        0,
	'>':        0,
	'_':        0,
	')':        0,
	';':        0,
	'{':        0,
	'}':        0,
	',':        0,
	'[':        0,
	']':        0,
	'\n':       0,
	EOF:        parser.EOF,
	OUT_OF_EOF: 0,
}

var runeCheck = map[rune]int{
	'+': 0,
	'-': 0,
	'*': 0,
	'/': 0,
	'%': 0,
	'(': 0,
	'.': 0,
	'!': 0,
	'&': 0,
	'|': 0,
	'=': 0,
	'^': 0,
	'~': 0,
	':': 0,
	'?': 0,
	'<': 0,
	'>': 0,
	'_': 0,
	')': 0,
	';': 0,
	'{': 0,
	'}': 0,
	',': 0,
	'[': 0,
	']': 0,
}

//	キーワードの定義
var keyName = map[string]int{
	"var":      parser.VAR,
	"func":     parser.FUNC,
	"return":   parser.RETURN,
	"if":       parser.IF,
	"for":      parser.FOR,
	"in":       parser.IN,
	"else":     parser.ELSE,
	"break":    parser.BREAK,
	"continue": parser.CONTINUE,
	"true":     parser.TRUE,
	"false":    parser.FALSE,
	"go":       parser.GO,
	"shell":    parser.SHELL,
	"range":    parser.RANGE,
	"nil":      parser.NIL,
}
