package scanner

import (
	"errors"
	//"fmt"
	"github.com/nanoeru/copal/common/token"
	"github.com/nanoeru/copal/shell/parser"
)

func New() *Scanner {
	return new(Scanner)
}

//	初期化(ソース代入)
func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

//	走査開始
func (s *Scanner) Scan() (tok int, lit string, pos token.Position) {
	var err error
	//	再走査ラベル
	//retry:
	pos = s.pos()
	//	空白文字無視
	if s.skipBlank() {
		//tok = parser.BLANK
		//lit = "BLANK"
		//return
	}
	pos = s.pos()

	//fmt.Println("scan-start", s.peek(), string(s.peek()))

	//	rune別にscanメソッドを実行
	switch ch := s.peek(); {
	//	関数によるrune判定
	//	キーワード・オペランド、識別子
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			tok = ParseError
		}
		//	キーワード・オペランドと識別子の判定
		if name, ok := keyName[lit]; ok {
			//fmt.Println("keyword", lit)
			tok = name
		} else {
			//fmt.Println("letter", lit)
			tok = parser.IDENT
		}
	//	数字
	case isDigit(ch):
		//fmt.Println("digit")
		tok = parser.NUMBER
		if s.peek() == '0' {
			lit, err = s.scanInt()
		} else {
			lit, err = s.scanFloat()
		}
		if err != nil {
			tok = ParseError
		}
	//	文字列
	case isQuote(ch):
		tok = parser.STRING
		lit, err = s.scanString()
		if err != nil {
			tok = ParseError
			lit = "\"" + lit + "\""
		}
	default:
		//	複雑な処理
		switch ch {
		//	1行コメント
		case '#':
			//fmt.Println("#	1行コメント")
			for !isEOL(s.peek()) {
				s.next()
			}
			s.next()
			//	コメント後に走査を続けられるように再捜査ラベルへ
			//goto retry
			tok = parser.COMMENT
			lit = "#"
			return
		//	1行コメント・複数コメント
		case '/':
			s.next()
			if s.peek() == '/' {
				for !isEOL(s.peek()) {
					s.next()
				}
				s.next()
				tok = parser.COMMENT
				lit = "//"
				return
				//	コメント後に走査を続けられるように再捜査ラベルへ
				//goto retry
			} else if s.peek() == '*' {
				s.next()
				err = s.scanComment()
				if err != nil {
					tok = ParseError
				} else {
					//	コメント後に走査を続けられるように再捜査ラベルへ
					//goto retry
					tok = parser.COMMENT
					lit = "/**/"
					return
				}
			} else {
				s.back()
			}
		default:
			tok = ParseError
			//			lit = ConvToString(ch)
			err = errors.New("Characters that are unexpected")
		}

		//	オペランド一致検証
		lit, err = s.scanOp()
		if err != nil {
			tok = ParseError
		}

		if name, ok := opName[lit]; ok {
			tok = name
			break
		} else {
			//	戻る
			for _, _ = range lit {
				s.back()
			}
			tok = ParseError
		}

		//	オペランド一致検証
		if v, ok := runeOpName[ch]; ok {
			tok = int(ch)
			lit = string(ch)
			if v != 0 {
				tok = v
			}
			s.next()
			break
		} else {
			tok = ParseError
		}

		s.next()
	}
	//	パースエラーの場合はエラーを文字列として返す
	if tok == ParseError {
		if err != nil {
			lit = "[" + lit + "] " + err.Error()
		} else {
			lit = "[" + lit + "] parse error"
		}
	}
	return
}

//	現在の位置の文字を読み取る
func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		if !s.reachNextEOF() {
			return EOF
		} else {
			return OUT_OF_EOF //	字句解析終了
		}
	}
}

//	次の文字へ進む
func (s *Scanner) next() {
	if !s.reachNextEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

//	戻る
func (s *Scanner) back() {
	s.offset--
}

//	EOF 判定
func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

//	Next EOF 判定
func (s *Scanner) reachNextEOF() bool {
	return len(s.src) < s.offset
}

//	位置情報構造体(Line, Column)取得
func (s *Scanner) pos() token.Position {
	return token.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

//	空白登録文字を無視する
func (s *Scanner) skipBlank() (flag bool) {
	if isBlank(s.peek()) {
		flag = true
	}
	for isBlank(s.peek()) {
		s.next()
	}
	return
}
