package ast

//	文インタフェース
type Stmt interface {
	stmt()
	GetLit() string
	SetLit(string)
}

//	文構造体
type StmtImpl struct {
	Lit string
}

func (x StmtImpl) stmt() {}
func (x StmtImpl) GetLit() string {
	return x.Lit
}
func (x *StmtImpl) SetLit(lit string) {
	x.Lit = lit
}

//	式文
type ExprStmt struct {
	StmtImpl
	Expr Expr
}

//	代入
type AssignStmt struct {
	StmtImpl
	IdentExprs []string
	Exprs      []Expr
}

//	新規代入
type NewAssignStmt struct {
	StmtImpl
	IdentExprs []string
	Exprs      []Expr
}

//	コマンド
type CmdStmt struct {
	StmtImpl
	Cmd     string
	Options []OptionExpr
	Args    []Expr
	In      Expr
	Out     Expr
}

//	関数文
type FuncStmt struct {
	StmtImpl
	Name  string
	Stmts []Stmt
	Args  []string
	//Args    []IdentExpr
	RetVals []string
}

//	Return
type ReturnStmt struct {
	StmtImpl
	Exprs []Expr
}

//	for文パターン
const (
	//	無限ループ
	InfiniteForLoop = iota
	//	条件のみ
	CondForLoop
	//	CondPostForLoop
	//	i := 0; i < n; i++
	InitCondPostForLoop

	//	k, v := range ???
	KeyValueRangeLoop
)

//	for文
type ForStmt struct {
	StmtImpl

	Init  Stmt
	Cond  Expr
	Post  Stmt
	Stmts []Stmt

	//Key   string
	//Value string
	Idents []string
	Expr   Expr

	Pattern int
}

//	if文
type IfStmt struct {
	StmtImpl
	Init      Stmt
	Cond      Expr
	ThenStmts []Stmt
	ElseStmts []Stmt
}

//	インクリメント文
type IncStmt struct {
	StmtImpl
	Expr Expr
}

//	デクリメント文
type DecStmt struct {
	StmtImpl
	Expr Expr
}

//	chan in 文
type ChanInStmt struct {
	StmtImpl
	Chan Expr
	Expr Expr
}

//	関数並列呼び出し
type GoCallStmt struct {
	StmtImpl
	Name     string
	Expr     Expr
	ArgExprs []Expr
}
