package token

//	トークン構造体
type Token struct {
	Tok int
	Lit string
	Pos Position
}

//	位置情報構造体
type Position struct {
	Line   int
	Column int
}
