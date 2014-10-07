package parser

import (
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/common/token"
	"github.com/nanoeru/copal/val"
	"strconv"
)

//	数値
func NumberExpr(tok token.Token) *ast.NumberExpr {
	lit := tok.Lit
	var tmp val.Val
	//	base -> 0 -> 基数自動認識
	if i, err := strconv.ParseInt(lit, 0, 64); err == nil {
		tmp = val.NewInt(int(i))
	} else if f, err := strconv.ParseFloat(lit, 64); err == nil {
		tmp = val.NewFloat(f)
	} else {
		panic("NOT SUPPORT NUMBER FORMAT")
	}
	v := ast.NumberExpr{V: tmp}
	v.SetLit(lit)
	exprLog.Info("NUMBER", lit)
	return &v
}

//	nil
func NilExpr(tok token.Token) *ast.NilExpr {
	lit := tok.Lit
	tmp := val.NewNil()
	v := ast.NilExpr{V: &tmp}
	v.SetLit(lit)
	exprLog.Info("Nil", lit)
	return &v
}

//	真偽値
func BoolExpr(flag bool, tok token.Token) *ast.BoolExpr {
	lit := tok.Lit
	tmp := val.Bool(flag)
	v := ast.BoolExpr{V: &tmp}
	v.SetLit(lit)
	exprLog.Info("BOOL", lit)
	return &v
}

//	識別子
func IdentExpr(tok token.Token) *ast.IdentExpr {
	lit := tok.Lit
	tmp := val.String(lit)
	v := ast.IdentExpr{V: &tmp}
	v.SetLit(lit)
	exprLog.Info("IDENT", lit)
	return &v
}

//	文字列
func StringExpr(tok token.Token) *ast.StringExpr {
	lit := tok.Lit
	tmp := val.String(lit)
	v := ast.StringExpr{V: &tmp}
	v.SetLit(lit)
	exprLog.Info("String", lit)
	return &v
}

//	単項演算
func UnaryOpExpr(op string, expr ast.Expr) *ast.UnaryOpExpr {
	lit := op + expr.GetLit()
	v := ast.UnaryOpExpr{Expr: expr, Operator: op}
	v.SetLit(lit)
	exprLog.Info("expr", op, "expr", lit)
	return &v
}

//	二項演算
func BinOpExpr(op string, lhs, rhs ast.Expr) *ast.BinOpExpr {
	lit := lhs.GetLit() + op + rhs.GetLit()
	v := ast.BinOpExpr{Lhs: lhs, Operator: op, Rhs: rhs}
	v.SetLit(lit)
	exprLog.Info("expr", op, "expr", lit)
	return &v
}
