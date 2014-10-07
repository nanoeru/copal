package semanticer

import (
	"fmt"
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/val"
	"github.com/nanoeru/copal/val/vpanic"
	"github.com/nanoeru/copal/vm"
)

//	Return文
func (s *Semanticer) parseReturnStmt(stmt *ast.ReturnStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	for i, v := range stmt.Exprs {
		_ = i
		fs.Combine(s.parseExpr(v))
	}
	return fs
}

//	関数文
func (s *Semanticer) parseFuncStmt(stmt *ast.FuncStmt, varFlag bool) *vm.FuncBlock {
	fs := vm.NewFuncBlock()

	vf := s.createFunc(stmt)
	if varFlag {
		//	ユーザー定義関数
		funcStack := s.GetVarStack(stmt.Name)
		fs.Funcs.Append(func() {
			funcStack.Push(vf)
			stackLog.Info("	pushed func", stmt.Name)
		})
		//	削除
		fs.DeferFuncs.Append(func() {
			funcStack.Pop()
			stackLog.Info("	poped func", stmt.Name)
		})
	} else {
		//	一時関数(無名関数)
		fs.Funcs.Append(func() {
			s.Push(vf)
		})
	}
	return fs
}

//	関数呼び出し
func (s *Semanticer) callFunc(expr *ast.CallExpr, presetFunc val.Val, popFlag, goFlag bool) *vm.FuncBlock {
	fs := vm.NewFuncBlock()

	//	組み込み関数
	builtin, builtinFlag := builtinFuncMap[expr.Name]
	//	実行関数の取得
	f := presetFunc
	if f == nil {
		//	実行時に決定
		if builtinFlag {
			//	組み込み関数
			f = val.NewFunc(builtin.F)
		} else if popFlag {
			//	実行時取得関数
			fs.Combine(s.parseExpr(expr.Expr))
			fs.Funcs.Append(func() {
				f = s.Pop()
			})
		} else {
			//	変数名指定関数
			varStack := s.GetVarStack(expr.Name)
			fs.Funcs.Append(func() {
				f = varStack.Get()
			})
		}
	}

	//	必要最小の引数の数
	var minArgN int
	//	可変長引数フラグ
	var variableFlag bool
	//	実際の引数の数
	actualArgN := len(expr.ArgExprs)

	if builtin.ArgN > 0 {
		minArgN = builtin.ArgN
		variableFlag = false
	} else {
		minArgN = -(builtin.ArgN + 1)
		variableFlag = true
	}

	//	組み込み関数の引数チェック
	if builtinFlag {
		//	引数チェック
		if variableFlag {
			//	可変長引数
			if minArgN > actualArgN {
				vpanic.Do("func", fmt.Sprintf("関数の引数の数が少ないです。 本来は%d以上 実際は%d", minArgN, actualArgN))
			}
		} else {
			//	固定引数
			if minArgN != actualArgN {
				vpanic.Do("func", fmt.Sprintf("関数の引数の数が異なります。 本来は%d 実際は%d", minArgN, actualArgN))
			}
		}
	}

	//	引数実行コード生成
	for _, v := range expr.ArgExprs {
		fs.Combine(s.parseExpr(v))
	}

	if !goFlag {
		//	引数
		vals := make([]val.Val, actualArgN)
		//	同期
		fs.Funcs.Append(func() {
			//	引数格納
			for i := 0; i < actualArgN; i++ {
				vals[actualArgN-1-i] = s.Pop()
			}
			for _, v := range f.Call(vals...) {
				s.Push(v)
			}
		})
	} else {
		//	非同期
		if !builtinFlag {
			//	引数
			vals := make([]val.Val, actualArgN)
			fs.Funcs.Append(func() {
				//	意味解析及び実行
				newS := New()
				newFs := vm.NewFuncBlock()

				//	並列処理終了待ち合わせ用
				wg.Add(1)

				//	スコープ構築
				for k, varStack := range s.VarMap {
					if varStack.Len() != 0 {
						v := varStack.Get()
						if vf, ok := v.(*val.Func); ok {
							//	関数の場合は現在のスコープに合わせた無名関数にする必要性がある
							newFs.Combine(newS.parseFuncStmt(vf.Stmt.(*ast.FuncStmt), true))
						} else {
							//	現在のスコープがアクセス可能な変数コピー
							//	変数の場合は同一だと参照先が同じとなり競合するのでコピー
							stack := newS.GetVarStack(k)
							stack.Push(v.Copy())
							newS.VarMap[k] = stack
						}
					}
				}

				//	引数実行
				newFs.Funcs.Do()

				//	引数準備
				for i := actualArgN - 1; i >= 0; i-- {
					//	変数は並列化する前に受け取るので以前のスコープから取得
					vals[i] = s.Pop()
					//	関数の場合は現在のスコープに合わせた関数を生成
					if vf, ok := vals[i].(*val.Func); ok {
						funcStmt := vf.Stmt.(*ast.FuncStmt)
						newS.parseFuncStmt(funcStmt, true).Do()
						varStack := newS.GetVarStack(funcStmt.Name)
						vals[i] = varStack.Get()
					}
				}

				if !popFlag {
					//	変数より取得
					varStack := newS.GetVarStack(expr.Name)
					f = varStack.Get()
				} else {
					//	返り値として取得
					funcStmt := f.Func().Stmt.(*ast.FuncStmt)
					newS.parseFuncStmt(funcStmt, false).Do()
					f = newS.Pop()
				}

				go func(f val.Val) {
					f.Call(vals...)
					//	新規スコープユーザー定義関数削除
					newFs.DeferFuncs.Do()
					//	並列処理終了待ち合わせ用
					wg.Done()
				}(f)
			})
		} else {
			//	組み込み関数
			fs.Funcs.Append(func() {
				//	新規作成をすることで、引数を格納する同一のスライスを指すことを回避
				vals := make([]val.Val, actualArgN)
				for i := 0; i < actualArgN; i++ {
					vals[actualArgN-1-i] = s.Pop().Copy()
				}
				wg.Add(1)
				//	for文などで同一文を繰り返す場合に、go func()において、クロージャーとして使用している値を上書きしてしまう可能性があるので、引数で渡す
				go func(f val.Val) {
					f.Call(vals...)
					//	終了の待ち合わせ
					wg.Done()
				}(f)
			})
		}
	}
	//}
	return fs
}

//	関数並列呼び出し文
func (s *Semanticer) parseGoCallStmt(stmt *ast.GoCallStmt) *vm.FuncBlock {
	fs := vm.NewFuncBlock()
	fs.Combine(s.callFunc(&ast.CallExpr{Name: stmt.Expr.GetLit(), Expr: stmt.Expr, ArgExprs: stmt.ArgExprs}, nil, true, true))
	return fs
}

//	関数作成
func (s *Semanticer) createFunc(stmt *ast.FuncStmt) *val.Func {
	block := s.parseStmts(stmt.Stmts)
	argNames := stmt.Args
	actualArgN := len(argNames)
	//	変数スタック
	varStacks := make([]*val.Stack, actualArgN)
	for i := 0; i < actualArgN; i++ {
		varStacks[i] = s.GetVarStack(argNames[i])
	}
	//	呼び出し・返り値処理
	vf := val.NewFunc(func(vs ...val.Val) []val.Val {
		//	TODO	引数チェック追加
		for i := 0; i < actualArgN; i++ {
			//			s.PushVar(argNames[i], vs[i])
			varStacks[i].Push(vs[i])
		}

		//fmt.Println("args", vs)
		block.Do()

		//	仕様上無駄な処理ではあるが、push->pop->pushすることになるはず
		length := len(stmt.RetVals)
		tmp := make([]val.Val, length)
		for i := length - 1; i >= 0; i-- {
			tmp[i] = s.Pop()
		}

		//for i := 0; i < length; i++ {
		//	s.Push(tmp[i])
		//}

		//	変数削除
		for i := 0; i < actualArgN; i++ {
			varStacks[i].Pop()
			//			s.PopVar(argNames[i])
		}
		return tmp
	})
	//	関数
	vf.SetStmt(stmt)
	//	関数のブロック
	vf.SetStmts(stmt.Stmts)
	return vf
}
