package scanner

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}

func isAlphabet(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

//	数字判定
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

//	8進数数字判定
func is8Digit(ch rune) bool {
	return '0' <= ch && ch <= '7'
}

//	16進数数字判定
func is16Digit(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

//	文字列開始・終了判定
func isQuote(ch rune) bool {
	return ch == '"'
}

//	読み込み終了
func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

//	空白文字関数
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' // || ch == '\n' // || ch == '\r'
}
