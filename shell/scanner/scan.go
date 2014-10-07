package scanner

import (
	"errors"
	"fmt"
	"strconv"
)

//	以下、scan判定メソッド
//	キーワードスキャン
func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for isLetter(s.peek()) || isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

//	オペランドスキャン
func (s *Scanner) scanOp() (string, error) {
	var ret []rune
	//for !isAlphabet(s.peek()) && !isDigit(s.peek()) {
	for {
		ch := s.peek()
		//fmt.Println(string(ch))
		if _, ok := runeCheck[ch]; ok {
			ret = append(ret, ch)
			s.next()
		} else {
			break
		}
	}
	return string(ret), nil
}

//	整数スキャン
func (s *Scanner) scanInt() (string, error) {
	var ret []rune
	if s.peek() == '0' {
		ret = append(ret, s.peek())
		s.next()
		if s.peek() == 'x' || s.peek() == 'X' {
			//	16進数表記
			ret = append(ret, s.peek())
			s.next()
			//	数字必須
			if is16Digit(s.peek()) {
				ret = append(ret, s.peek())
				s.next()
				for is16Digit(s.peek()) {
					ret = append(ret, s.peek())
					s.next()
				}
			} else {
				return string(ret), errors.New("int based 16 parse error: Require number after x or X")
			}
		} else {
			//	8進数表記
			for is8Digit(s.peek()) {
				ret = append(ret, s.peek())
				s.next()
			}
			if s.peek() == '.' {
				//	小数
				fret, err := s.scanFloat()
				ret = append(ret, []rune(fret)...)
				return string(ret), err
			}
			if isDigit(s.peek()) {
				for isDigit(s.peek()) {
					ret = append(ret, s.peek())
					s.next()
				}
				if s.peek() == '.' {
					//	小数
					fret, err := s.scanFloat()
					ret = append(ret, []rune(fret)...)
					return string(ret), err
				} else {
					return string(ret), errors.New("int parse error: Require 0-7 number after 0")
				}
			}
		}
	} else {
		//	10進数表記
		for isDigit(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
		}
	}
	return string(ret), nil
}

//	小数スキャン
func (s *Scanner) scanFloat() (string, error) {
	var ret []rune
	//	[num][.][num][e][+,-][num]
	for isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	if s.peek() == '.' {
		ret = append(ret, s.peek())
		s.next()
		for isDigit(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
		}
	}
	if s.peek() == 'e' || s.peek() == 'E' {
		ret = append(ret, s.peek())
		s.next()
		if s.peek() == '+' || s.peek() == '-' {
			ret = append(ret, s.peek())
			s.next()
		}
		//	数字必須
		if isDigit(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
			for isDigit(s.peek()) {
				ret = append(ret, s.peek())
				s.next()
			}
		} else {
			return string(ret), errors.New("float parser error: Require number after e or E")
		}
	}
	return string(ret), nil
}

//	コメント解析
func (s *Scanner) scanComment() error {
	for {
		//fmt.Println(string(s.peek()))
		//	入力終了->エラー
		if s.peek() == OUT_OF_EOF {
			return errors.New("Comment Parser Error Require */")
		}
		//	特殊文字処理
		if s.peek() == '*' {
			s.next()
			if s.peek() == '/' {
				s.next()
				return nil
			}
			s.back()
		}
		s.next()
	}
}

//	文字列解析
func (s *Scanner) scanString() (string, error) {
	var ret []rune
	for {
		s.next()
		//	入力終了->エラー
		if s.peek() == EOF {
			return string(ret), errors.New("String Parser Error Require [\"]")
		}
		//	特殊文字処理
		if s.peek() == '\\' {
			s.next()
			switch s.peek() {
			case '\\':
				ret = append(ret, '\\')
				continue
			case '"':
				ret = append(ret, '"')
				continue
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			//	16進数表記で文字コード入力
			case 'x':
				s.next()
				var num []rune
				//	数字必須
				for is16Digit(s.peek()) {
					num = append(num, s.peek())
					s.next()
				}
				s.back()
				//fmt.Println(string(num))
				if i, err := strconv.ParseInt("0x"+string(num), 0, 64); err == nil {
					ret = append(ret, rune(i))
				} else {
					return string(ret), errors.New(fmt.Sprint("String Parser Error After [\\x]", err))
				}
				continue
			}
			return string(ret), errors.New("String Parser Error After [\\]")
		}
		//	文字列終了
		if isQuote(s.peek()) {
			s.next()
			break
		}
		//	現在の解析文字を文字列に追加
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}
