package semanticer

import (
	"fmt"
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/val"
	"github.com/nanoeru/copal/val/vpanic"
	"github.com/nanoeru/copal/vm"
	"os"
	"sync"
)

//	TODO	グローバルから変更
//	並列処理終了の同期
var wg sync.WaitGroup

//	メインブロック(主要スレッド)
func (s *Semanticer) parseMainStmts(stmts []ast.Stmt) *vm.FuncBlock {
	fs := s.parseStmts(stmts)
	//	並列処理終了待ち合わせ用
	fs.Funcs.Append(func() {
		wg.Wait()
	})
	return fs
}

//	ブロック
func (s *Semanticer) parseStmts(stmts []ast.Stmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	for _, stmt := range stmts {
		fs.Combine(s.parseStmt(stmt))
	}
	//	ブロック終了時遅延実行
	fs.Funcs.Combine(fs.DeferFuncs)
	fs.CleanDeferFunc()
	return fs
}

//	値残り文
func (s *Semanticer) parseExprStmt(stmt *ast.ExprStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	fs.Combine(s.parseExpr(stmt.Expr))
	return fs
}

//	コマンド文
func (s *Semanticer) parseCmdStmt(stmt *ast.CmdStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	for i, v := range stmt.Args {
		_ = i
		fs.Combine(s.parseExpr(v))
		fs.Funcs.Append(func() {
			fmt.Fprintln(os.Stdout, s.Pop().String())
		})
	}
	return fs
}

//	インクリメント文
func (s *Semanticer) parseIncStmt(stmt *ast.IncStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	////	変数スタック
	//varStack := s.GetVarStack(stmt.Expr..IdentExprs[i])

	fs.Combine(s.parseExpr(stmt.Expr))
	fs.Funcs.Append(func() {
		//varStack.Set(s.Pop())

		v := s.Pop()
		v.Inc()
	})
	return fs
}

//	デクリメント文
func (s *Semanticer) parseDecStmt(stmt *ast.DecStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	fs.Combine(s.parseExpr(stmt.Expr))
	fs.Funcs.Append(func() {
		v := s.Pop()
		v.Dec()
	})
	return fs
}

//	if文
func (s *Semanticer) parseIfStmt(stmt *ast.IfStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	init := s.parseStmt(stmt.Init)
	deferFunc := init.DeferFuncs
	init.CleanDeferFunc()
	cond := s.parseExpr(stmt.Cond)
	thenBlock := s.parseStmts(stmt.ThenStmts)
	elseBlock := s.parseStmts(stmt.ElseStmts)

	//	if文のパターン
	if !init.IsBlank() {
		fs.Funcs.Append(func() {
			if init.Do(); func() val.Bool {
				cond.Do()
				return s.Pop().Bool()
			}() {
				thenBlock.Do()
			} else {
				elseBlock.Do()
			}
			//	遅延実行処理
			deferFunc.Do()
		})
	} else {
		//	初期化文なし
		fs.Funcs.Append(func() {
			cond.Do()
			if s.Pop().Bool() {
				thenBlock.Do()
			} else {
				elseBlock.Do()
			}
		})
	}
	return fs
}

//	for文
func (s *Semanticer) parseForStmt(stmt *ast.ForStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()

	//	各処理のパース
	init := s.parseStmt(stmt.Init)
	deferFunc := init.DeferFuncs
	init.CleanDeferFunc()
	cond := s.parseExpr(stmt.Cond)
	post := s.parseStmt(stmt.Post)
	block := s.parseStmts(stmt.Stmts)

	//	range用
	idents := stmt.Idents
	//expr := s.parseExpr(stmt.Expr)

	//	for文のパターン
	var f func()
	switch stmt.Pattern {
	//	無限ループ
	case ast.InfiniteForLoop:
		f = func() {
			for {
				block.Do()
			}
		}
	//	条件のみ
	case ast.CondForLoop:
		f = func() {
			for func() val.Bool {
				cond.Do()
				return s.Pop().Bool()
			}() {
				block.Do()
			}
		}
	//	全て
	case ast.InitCondPostForLoop:
		f = func() {
			for init.Do(); func() val.Bool {
				cond.Do()
				return s.Pop().Bool()
			}(); post.Do() {
				block.Do()
			}
			//	遅延実行処理
			deferFunc.Do()
		}
	//	イテレート
	case ast.KeyValueRangeLoop:
		switch len(idents) {
		//case 1:
		case 2:
			key := idents[0]
			value := idents[1]

			//	range する値
			fs.Combine(s.parseExpr(stmt.Expr))

			//	変数スタック
			keyVarStack := s.GetVarStack(key)
			valueVarStack := s.GetVarStack(value)

			f = func() {
				expr := s.Pop()
				keyVarStack.Push(nil)
				valueVarStack.Push(nil)
				//	各typeに変換してrange
				switch rv := expr.(type) {
				//	スライス
				case *val.Slice:
					//for k, v := range *rv {
					//	keyVarStack.Set(val.NewInt(k))
					//	valueVarStack.Set(v)
					//	block.Do()
					//}
				//	スライス
				case val.Slice:
					for k, v := range rv {
						keyVarStack.Set(val.NewInt(k))
						valueVarStack.Set(v)
						block.Do()
					}
				//	文字列
				case *val.String:
					for k, v := range *rv {
						keyVarStack.Set(val.NewInt(k))
						valueVarStack.Set(val.NewString(string(v)))
						block.Do()
					}
				////	文字列
				//case val.String:
				//	for k, v := range rv {
				//		keyVarStack.Set(val.NewInt(k))
				//		valueVarStack.Set(val.NewString(string(v)))
				//		block.Do()
				//	}
				//	チャネル
				case *val.Chan:
					k := 0
					for v := range *rv {
						keyVarStack.Set(val.NewInt(k))
						valueVarStack.Set(v)
						block.Do()
						k++
					}
					//case val.Chan:
					//	k := 0
					//	for v := range rv {
					//		keyVarStack.Set(val.NewInt(k))
					//		valueVarStack.Set(v)
					//		block.Do()
					//		k++
					//	}
				default:
					vpanic.Do("for range", fmt.Sprint("require slice or string or channel, not ", expr.Type()))
				}
				keyVarStack.Pop()
				valueVarStack.Pop()
			}
		default:
			vpanic.Do("for range", "require 2 idents")
		}
	default:
		vpanic.Do("for range", fmt.Sprintln("unknown pattern", stmt.Pattern))
	}
	fs.Funcs.Append(f)
	return fs
}

//	代入文
func (s *Semanticer) parseAssignStmt(stmt *ast.AssignStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()

	//	右辺値
	exprLen := len(stmt.Exprs)
	for i := 0; i < exprLen; i++ {
		fs.Combine(s.parseExpr(stmt.Exprs[i]))
	}

	//	左辺値
	leftLen := len(stmt.IdentExprs)
	for i := leftLen - 1; i >= 0; i-- {
		v := stmt.IdentExprs[i]
		varStack := s.GetVarStack(v)
		//	捨て変数
		if v == "_" {
			//	変数スタック
			fs.Funcs.Append(func() {
				s.Pop()
			})
		} else {
			//	変数スタック
			fs.Funcs.Append(func() {
				varStack.Set(s.Pop())
			})
		}
	}
	return fs
}

//	新規宣言代入文
func (s *Semanticer) parseNewAssignStmt(stmt *ast.NewAssignStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()

	//	右辺値
	exprLen := len(stmt.Exprs)
	for i := 0; i < exprLen; i++ {
		fs.Combine(s.parseExpr(stmt.Exprs[i]))
	}

	//	左辺値
	leftLen := len(stmt.IdentExprs)
	for i := leftLen - 1; i >= 0; i-- {
		v := stmt.IdentExprs[i]
		varStack := s.GetVarStack(v)
		//	捨て変数
		if v == "_" {
			//	変数スタック
			fs.Funcs.Append(func() {
				s.Pop()
			})
		} else {
			//	変数スタック
			fs.Funcs.Append(func() {
				varStack.Push(s.Pop())
			})
			//	削除
			fs.DeferFuncs.Append(func() {
				varStack.Pop()
			})
		}
	}
	return fs
}

//	chan in 文
func (s *Semanticer) parseChanInStmt(stmt *ast.ChanInStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	fs.Combine(s.parseExpr(stmt.Chan))
	fs.Combine(s.parseExpr(stmt.Expr))
	fs.Funcs.Append(func() {
		v := s.Pop()
		ch := s.Pop()
		ch.In(v)
	})
	return fs
}

//	文
func (s *Semanticer) parseStmt(stmt ast.Stmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	switch stmt := stmt.(type) {
	////	宣言文
	//case *ast.DeclStmt:
	//	fs.Combine(parseDecl(stmt.Decl))
	//	関数文
	case *ast.FuncStmt:
		fs.Combine(s.parseFuncStmt(stmt, true))
	//	関数並列呼び出し文
	case *ast.GoCallStmt:
		fs.Combine(s.parseGoCallStmt(stmt))
	//	代入文
	case *ast.AssignStmt:
		fs.Combine(s.parseAssignStmt(stmt))
	//	新規代入文
	case *ast.NewAssignStmt:
		fs.Combine(s.parseNewAssignStmt(stmt))
	//	chan in 文
	case *ast.ChanInStmt:
		fs.Combine(s.parseChanInStmt(stmt))
	//	インクリメント文
	case *ast.IncStmt:
		fs.Combine(s.parseIncStmt(stmt))
	//	デクリメント文
	case *ast.DecStmt:
		fs.Combine(s.parseDecStmt(stmt))
	//	for文
	case *ast.ForStmt:
		fs.Combine(s.parseForStmt(stmt))
	//	if文
	case *ast.IfStmt:
		fs.Combine(s.parseIfStmt(stmt))
	//	単一式文
	case *ast.CmdStmt:
		fs.Combine(s.parseCmdStmt(stmt))
	//	Return文
	case *ast.ReturnStmt:
		fs.Combine(s.parseReturnStmt(stmt))
	//	単一式文
	case *ast.ExprStmt:
		fs.Combine(s.parseExprStmt(stmt))
	default:
		//fmt.Printf("parseStmt: %v %T\n", stmt, stmt)
	}
	return fs
}
