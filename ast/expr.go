package ast

import (
	"github.com/nanoeru/copal/val"
)

//	式インタフェース
type Expr interface {
	expr()
	GetLit() string
	SetLit(string)
}

//	式構造体
type ExprImpl struct {
	Lit string
}

func (x *ExprImpl) expr() {}

func (x ExprImpl) GetLit() string {
	return x.Lit
}

func (x *ExprImpl) SetLit(lit string) {
	x.Lit = lit
}

//	数値
type NumberExpr struct {
	ExprImpl
	V val.Val
}

//	真偽値
type BoolExpr struct {
	ExprImpl
	V val.Val
}

//	nil
type NilExpr struct {
	ExprImpl
	V val.Val
}

//	式・演算子・式
type BinOpExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

//	単項演算子の式
type UnaryOpExpr struct {
	ExprImpl
	Operator string
	Expr     Expr
}

//	文字列式
type StringExpr struct {
	ExprImpl
	V val.Val
}

//	キーワード以外
type IdentExpr struct {
	ExprImpl
	V val.Val
}

//	オプション
type OptionExpr struct {
	ExprImpl
	FlagName string
	Expr     Expr
}

//	関数呼び出し
type CallExpr struct {
	ExprImpl
	Name     string
	Expr     Expr
	ArgExprs []Expr
}

//	関数
type FuncExpr struct {
	ExprImpl
	Args    []string
	Stmts   []Stmt
	RetVals []string
}

//	関数呼び出し
type FuncCallExpr struct {
	ExprImpl
	Expr     Expr
	ArgExprs []Expr
}

//	スライス
type SliceExpr struct {
	ExprImpl
	Exprs []Expr
}

//	chan out
type ChanOutExpr struct {
	ExprImpl
	Chan Expr
}

//	mapアクセス
type MapAccessExpr struct {
	ExprImpl
	Name    string
	MapExpr Expr
	Expr    Expr
}
