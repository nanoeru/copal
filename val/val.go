package val

import ()

type (
	//	型インターフェース
	Val interface {
		TypeImple
		OpImpl
		BuiltinImpl
		CompImpl
		NumImpl
		CommonImpl
		FuncImpl
		//Stmts() []ast.Stmt
		ChanImpl
		MapImpl
		CondImpl
		UnaryOpImpl
		BitImpl
	}

	//	共通
	CommonImpl interface {
		Copy() Val
		Type() String
	}

	//	数値型
	NumImpl interface {
		Inc()
		Dec()
	}

	//	マップ型
	MapImpl interface {
		Index(Int) Val
		Map(Val) Val
	}

	//	チャネル型
	ChanImpl interface {
		In(Val)
		Out() Val
		Close()
	}

	//	関数型
	FuncImpl interface {
		Call(...Val) []Val
	}

	//	比較
	CompImpl interface {
		Equal(Val) Val
		NotEqual(Val) Val
		Less(Val) Val
		Greater(Val) Val
		LessOrEqual(Val) Val
		GreaterOrEqual(Val) Val
	}

	//	条件
	CondImpl interface {
		And(Val) Val
		Or(Val) Val
		Xor(Val) Val
		Not() Val
		IsNil() Bool
	}

	//	ビット演算
	BitImpl interface {
		Shr(Val) Val    //	>>
		Shl(Val) Val    //	<<
		BitAnd(Val) Val //	&
		BitOr(Val) Val  //	|
		BitXor(Val) Val //	^
		BitNot() Val    //	~
	}

	//	単項演算
	UnaryOpImpl interface {
		Minus() Val
	}

	//	組み込み
	BuiltinImpl interface {
		Len() Int
		Cap() Int
	}

	//	演算子
	OpImpl interface {
		Add(Val) Val //	+
		Sub(Val) Val //	-
		Mul(Val) Val //	*
		Div(Val) Val //	/
		Rem(Val) Val //	%
	}

	//	型変換
	TypeImple interface {
		Int() Int
		Float() Float
		String() String
		Bool() Bool
		Slice() Slice
		Chan() Chan
		Func() Func
		Nil() Nil
	}

	VFunc func(...Val) []Val

	Int    int
	Float  float64
	String string
	Bool   bool
	Nil    struct{}
	Slice  []Val
	Chan   chan Val
	//FuncCache struct {
	//	Idents []string
	//	//ArgVal []Val
	//	//		RetVal []Val
	//}
	Func struct {
		//		Name string
		//F func(...Val) []Val
		F VFunc
		//F func()
		//	簡易的
		Val
		Stmt  interface{}
		Stmts interface{}
		//		Cache map[FuncCache][]Val
	}
)

var (
	//BinOpIntMethodMap = map[string]func(Val, Int) Int{
	//"+": Val.AddInt,
	//"-": Val.SubInt,
	//"*":  Val.Mul,
	//"/":  Val.Div,
	//"==": Val.Equal,
	//"!=": Val.NotEqual,
	//"<":  Val.Less,
	//">":  Val.Greater,
	//"<=": Val.LessOrEqual,
	//">=": Val.GreaterOrEqual,
	//"&&": Val.And,
	//"||": Val.Or,
	//"^":  Val.Xor,
	//}
	UnaryOpMethodMap = map[string]func(Val) Val{
		"-": Val.Minus,
		"!": Val.Not,
		"~": Val.BitNot,
	}
	BinOpMethodMap = map[string]func(Val, Val) Val{
		"+":  Val.Add,
		"-":  Val.Sub,
		"*":  Val.Mul,
		"/":  Val.Div,
		"%":  Val.Rem,
		"==": Val.Equal,
		"!=": Val.NotEqual,
		"<":  Val.Less,
		">":  Val.Greater,
		"<=": Val.LessOrEqual,
		">=": Val.GreaterOrEqual,
		"&&": Val.And,
		"||": Val.Or,
		"^^": Val.Xor,
		"&":  Val.BitAnd,
		"|":  Val.BitOr,
		"^":  Val.BitXor,
		"<<": Val.Shl,
		">>": Val.Shr,
	}
)
