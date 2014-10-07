package semanticer

import (
	"fmt"
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/val"
	"github.com/nanoeru/copal/vm"
)

//	値
func (s *Semanticer) parseExpr(expr ast.Expr) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	var v val.Val
	switch expr := expr.(type) {
	//	変数取得
	case *ast.IdentExpr:
		varStack := s.GetVarStack(expr.Lit)
		fs.Funcs.Append(func() {
			//fmt.Println("get-var", expr.Lit)
			//fmt.Println("get-var", varStack.Get().String())
			//			s.Push(s.GetVar(expr.Lit))
			s.Push((varStack.Get()))
			//fmt.Println(s.Pop())
		})
	case *ast.NumberExpr:
		v = expr.V
		goto copy
	case *ast.BoolExpr:
		v = expr.V
		goto copy
	case *ast.NilExpr:
		v = expr.V
		goto copy
	case *ast.StringExpr:
		v = expr.V
		goto copy
	//	スライス初期化
	case *ast.SliceExpr:
		for _, v := range expr.Exprs {
			fs.Combine(s.parseExpr(v))
		}
		length := len(expr.Exprs)
		fs.Funcs.Append(func() {
			slice := val.Slice(make([]val.Val, length))
			for i := length - 1; i >= 0; i-- {
				slice[i] = s.Pop()
			}
			s.Push(slice)
		})
	//chan out
	case *ast.ChanOutExpr:
		fs.Combine(s.parseExpr(expr.Chan))
		fs.Funcs.Append(func() {
			ch := s.Pop()
			s.Push(ch.Out())
		})
	//	無名関数
	case *ast.FuncExpr:
		v = s.createFunc(&ast.FuncStmt{Name: "", Args: expr.Args, RetVals: expr.RetVals, Stmts: expr.Stmts})
		goto copy
	//	単項演算
	case *ast.UnaryOpExpr:
		fs.Combine(s.parseExpr(expr.Expr))
		f, ok := val.UnaryOpMethodMap[expr.Operator]
		if !ok {
			panic(fmt.Sprintln("unknown operatpr", expr.Operator))
		} else {
			fs.Funcs.Append(func() {
				expr := s.Pop()
				//fmt.Println("unaryopexpr", lhs.String(), rhs.String())
				s.Push(f(expr))
			})
		}
	//	2項演算
	case *ast.BinOpExpr:
		fs.Combine(s.parseExpr(expr.Lhs))
		fs.Combine(s.parseExpr(expr.Rhs))
		f, ok := val.BinOpMethodMap[expr.Operator]
		if !ok {
			panic(fmt.Sprintln("unknown operatpr", expr.Operator))
		} else {
			fs.Funcs.Append(func() {
				rhs := s.Pop()
				lhs := s.Pop()
				//fmt.Println("binopexpr", lhs.String(), rhs.String())
				s.Push(f(lhs, rhs))
			})
		}
	//	関数呼び出し
	case *ast.CallExpr:
		fs.Combine(s.callFunc(expr, nil, false, false))
	//	関数呼び出し
	case *ast.FuncCallExpr:
		fs.Combine(s.callFunc(&ast.CallExpr{Name: expr.Expr.GetLit(), Expr: expr.Expr, ArgExprs: expr.ArgExprs}, nil, expr.Expr.GetLit() == "", false))
	//	マップアクセス
	case *ast.MapAccessExpr:
		fs.Combine(s.parseExpr(expr.MapExpr))
		if expr.Name != "" {
			//	文字列でマップアクセス
			fs.Funcs.Append(func() {
				s.Push(s.Pop().Map(val.String(expr.Name)))
			})
		} else {
			//	返り値でマップアクセス
			fs.Combine(s.parseExpr(expr.Expr))
			fs.Funcs.Append(func() {
				v := s.Pop()
				s.Push(s.Pop().Map(v))
			})
		}
	default:
		//fmt.Printf("parseExpr: %v %T\n", expr, expr)
	}
	return fs
copy:
	fs.Funcs.Append(func() {
		//fmt.Println("pushed", v.String())
		s.Push(v.Copy())
	})
	return fs
}
