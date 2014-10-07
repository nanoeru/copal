%{
package parser

import (
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/common/token"
)

//	ダミーデータ用の構造体
type Empty struct {}
%}

%type<stmts> block
%type<stmts> stmts
%type<stmt> stmt
%type<pre_stmt> pre_stmt
%type<stmt_delimiter> stmt_delimiter
%type<expr> expr
%type<exprs> exprs
// %type<option> option
// %type<options> options
// %type<arg> arg
// %type<args> args
%type<idents> idents

%union{
//	block          []ast.Stmt
	stmts          []ast.Stmt
	stmt           ast.Stmt
	pre_stmt       ast.Stmt
	stmt_delimiter interface{}
	expr           ast.Expr
	exprs          []ast.Expr
	idents         []string
//	idents         []ast.IdentExpr
	// option         *ast.OptionExpr
	// options        []ast.OptionExpr
	// arg            ast.Expr
	// args           []ast.Expr
	tok            token.Token
}

//	トークン定義
%token<tok> IDENT NUMBER STRING EOF NEW_ASSIGN
// %token<tok> OPTION

%token<tok> ADD_ASSIGN SUB_ASSIGN MUL_ASSIGN DIV_ASSIGN REM_ASSIGN XOR_ASSIGN

%token<tok> INC DEC VAR FUNC RETURN IF NIL

%token<tok> FOR IN ELSE BREAK CONTINUE TRUE FALSE GO

%token<tok> XOR BLANK

%token<tok> SHELL

%token<tok> COMMENT

%token<tok> RANGE

//	結合方法
%left CHAN
%left OR XOR
%left AND
%left '|' '&' '^'
%left '<' '>' LE GE NE EQ
%left '+' '-'
%left '*' '/' '%'
%left SHR SHL//	実際には異なる優先順位
%right UNARY
%left '!' '~'
%left '[' ']'
%left '(' ')'
%left '.'
%left LEFT

%%
//	複数の文
stmts :
	{
		$$ = nil
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}
	| stmt stmts
	{
		$$ = append([]ast.Stmt{$1}, $2...)
		if l, ok := yylex.(*Lexer); ok {
			l.stmts = $$
		}
	}

stmt :
	pre_stmt stmt_delimiter
	{
		$$ = $1
	}
	| stmt_delimiter
	{
		$$ = nil
	}

//	ブロック
block :
	'{' stmts '}'
	{
		$$ = $2
	}

//	文
pre_stmt :
	expr
	{
		lit := $1.GetLit()
		stmt := &ast.ExprStmt{Expr: $1}
		stmt.SetLit(lit)
		$$ = stmt
	}
	| idents '=' exprs
	{
		stmt := &ast.AssignStmt{IdentExprs: $1, Exprs: $3}
		$$ = stmt
	}
	| idents NEW_ASSIGN exprs
	{
		stmt := &ast.NewAssignStmt{IdentExprs: $1, Exprs: $3}
		$$ = stmt
	}
	//	A <- B
	| expr CHAN expr
	{
		stmt := &ast.ChanInStmt{Chan: $1, Expr: $3}
		$$ = stmt
	}
	//	複数のチャネルの渡し方
	// | exprs CHAN exprs
	// {
	// 	stmt := &ast.ChanInStmt{Chans: $1, Exprs: $3}
	// 	$$ = stmt
	// }
	// | SHELL exprs
	// {
	// 	lit := "$2.Lit"// + " " + $2.GetLit()
	// 	stmt := &ast.CmdStmt{Cmd: "$2.Lit", Args: $2}
	// 	stmt.SetLit(lit)
	// 	$$ = stmt
	// 	// fmt.Println("コマンド実行", lit)
	// }
	//	インクリメント
	| expr INC
	{
		$$ = &ast.IncStmt{Expr: $1}
	}
	//	デクリメント
	| expr DEC
	{
		$$ = &ast.DecStmt{Expr: $1}
	}
	//	if [init]; [cond] { [block] }
	| IF pre_stmt ';' expr block
	{
		$$ = &ast.IfStmt{Init: $2, Cond: $4, ThenStmts: $5}
	}
	//	if [cond] { [block] }
	| IF expr block
	{
		$$ = &ast.IfStmt{Cond: $2, ThenStmts: $3}
	}
	//	if [init]; [cond] { [block] } else { [block] }
	| IF pre_stmt ';' expr block ELSE block
	{
		$$ = &ast.IfStmt{Init: $2, Cond: $4, ThenStmts: $5, ElseStmts: $7}
	}
	//	if [init]; [cond] { [block] } else { [block] }
	| IF expr block ELSE block
	{
		$$ = &ast.IfStmt{Cond: $2, ThenStmts: $3, ElseStmts: $5}
	}
	//	for { [block] }
	| FOR block
	{
		$$ = &ast.ForStmt{Stmts: $2, Pattern: ast.InfiniteForLoop}
	}
	//	for [init]; [cond]; [post] { [block] }
	| FOR expr block
	{
		$$ = &ast.ForStmt{Cond: $2, Stmts: $3, Pattern: ast.CondForLoop}
	}
	//	for [init]; [cond]; [post] { [block] }
	| FOR pre_stmt ';' expr ';' pre_stmt block
	{
		$$ = &ast.ForStmt{Init: $2, Cond: $4, Post: $6, Stmts: $7, Pattern: ast.InitCondPostForLoop}
	}
	| FOR idents NEW_ASSIGN RANGE expr block
	{
		$$ = &ast.ForStmt{Idents: $2, Expr: $5, Stmts: $6, Pattern: ast.KeyValueRangeLoop}
	}
	| FUNC IDENT '(' idents ')' block
	{
		$$ = &ast.FuncStmt{Name: $2.Lit, Args: $4, Stmts: $6}
	}
	| FUNC IDENT '(' idents ')' '(' idents ')' block
	{
		$$ = &ast.FuncStmt{Name: $2.Lit, Args: $4, RetVals: $7, Stmts: $9}
	}
	| RETURN exprs
	{
		$$ = &ast.ReturnStmt{Exprs: $2}		
	}
	//	関数呼び出し
	| GO expr '(' exprs ')'
	{
		$$ = &ast.GoCallStmt{Expr: $2, ArgExprs: $4}
	}

//	文の区切り
stmt_delimiter :
	COMMENT
	{
		$$ = &Empty{}
	}
	| EOF
	{
		$$ = &Empty{}
	}
	| ';'
	{
		$$ = &Empty{}
	}
	| '\n'
	{
		$$ = &Empty{}
	}

// args :
// 	// {
// 	// 	$$ = []ast.Expr{}
// 	// }
// 	// |
// 	arg
// 	{
// 		$$ = []ast.Expr{$1}
// 	}
// 	| args arg
// 	{
// 		$$ = append($1, $2)
// 	}

// option: OPTION '=' expr
// 	{
// 		lit := "-" + $1.Lit + "=" + $3.GetLit()
// 		expr := &ast.OptionExpr{FlagName: $1.Lit, Expr: $3}
// 		expr.SetLit(lit)
// 		$$ = expr
// 		fmt.Println("OPTION=expr", lit)
// 	}
// 	| OPTION
// 	{
// 		lit := "-" + $1.Lit
// 		expr := &ast.OptionExpr{FlagName: $1.Lit, Expr: nil}
// 		expr.SetLit(lit)
// 		$$ = expr
// 		fmt.Println("OPTION", lit)
// 	}

//	連続IDENT
idents :
	{
		$$ = []string{}		
	}
	| IDENT
	{
		$$ = []string{$1.Lit}
	}
	| idents ',' IDENT
	{
		$$ = append($1, $3.Lit)
	}

//	複数の式	arg0, arg1, arg2, argn...
exprs :
	{
		$$ = []ast.Expr{}
	}
	| expr
	{
		$$ = []ast.Expr{$1}
	}
	| exprs ','  expr
	{
		$$ = append($1, $3)
	}

// arg : expr
// 	{
// 		$$ = $1
// 	}
// 	| option
// 	{
// 		$$ = $1
// 	}

//	単一式
expr : NUMBER
	{
		$$ = NumberExpr($1)
	}
	//	真偽値
	| TRUE
	{
		$$ = BoolExpr(true, $1)
	}
	//	nil
	| NIL
	{
		$$ = NilExpr($1)
	}
	//	真偽値
	| FALSE
	{
		$$ = BoolExpr(false, $1)
	}
	| '$' IDENT
	{
		//	廃止?!
		$$ = IdentExpr($2)
	}
	| IDENT
	{
		$$ = IdentExpr($1)
	}
	//	mapアクセス
	| expr '.' IDENT
	{
		$$ = &ast.MapAccessExpr{MapExpr: $1, Name: $3.Lit}
	}
	| expr '[' expr ']'
	{
		$$ = &ast.MapAccessExpr{MapExpr: $1, Expr: $3}
	}
	//	関数呼び出し
	| expr '(' exprs ')'
	{
		$$ = &ast.FuncCallExpr{Expr: $1, ArgExprs: $3}
	}
	//	スライス
	| '[' ']' '{' exprs '}'
	{
		$$ = &ast.SliceExpr{Exprs: $4}
	}
	//	関数の変数的宣言	引数名の羅列
	| FUNC '(' idents ')' block
	{
		$$ = &ast.FuncExpr{Args: $3, Stmts: $5}
	}
	//	関数の変数的宣言	引数名の羅列
	| FUNC '(' idents ')' '(' idents ')' block
	{
		$$ = &ast.FuncExpr{Args: $3, RetVals: $6, Stmts: $8}
	}
	| '-' expr %prec UNARY
	{
		$$ = UnaryOpExpr("-", $2)
	}
	| STRING
	{
		$$ = StringExpr($1)
	}
	| '(' expr ')'
	{
		lit := "( " +  $2.GetLit() + " )"
		$2.SetLit(lit)
		$$ = $2
		// fmt.Println("( expr )", lit)
	}
	| expr '+' expr
	{
		$$ = BinOpExpr("+", $1, $3)
	}
	| expr '-' expr
	{
		$$ = BinOpExpr("-", $1, $3)
	}
	| expr '*' expr
	{
		$$ = BinOpExpr("*", $1, $3)
	}
	| expr '/' expr
	{
		$$ = BinOpExpr("/", $1, $3)
	}
	| expr '%' expr
	{
		$$ = BinOpExpr("%", $1, $3)
	}
	| expr '<' expr
	{
		$$ = BinOpExpr("<", $1, $3)
	}
	| expr '>' expr
	{
		$$ = BinOpExpr(">", $1, $3)
	}
	| expr LE expr
	{
		$$ = BinOpExpr("<=", $1, $3)
	}
	| expr GE expr
	{
		$$ = BinOpExpr(">=", $1, $3)
	}
	| expr EQ expr
	{
		$$ = BinOpExpr("==", $1, $3)
	}
	| expr NE expr
	{
		$$ = BinOpExpr("!=", $1, $3)
	}
	| expr SHL expr
	{
		$$ = BinOpExpr("<<", $1, $3)
	}
	| expr SHR expr
	{
		$$ = BinOpExpr(">>", $1, $3)
	}
	| expr AND expr
	{
		$$ = BinOpExpr("&&", $1, $3)
	}
	| expr OR expr
	{
		$$ = BinOpExpr("||", $1, $3)
	}
	| expr XOR expr
	{
		$$ = BinOpExpr("^^", $1, $3)
	}
	| expr '&' expr
	{
		$$ = BinOpExpr("&", $1, $3)
	}
	| expr '|' expr
	{
		$$ = BinOpExpr("|", $1, $3)
	}
	| expr '^' expr
	{
		$$ = BinOpExpr("^", $1, $3)
	}
	| '!' expr
	{
		$$ = UnaryOpExpr("!", $2)
	}
	| '~' expr
	{
		$$ = UnaryOpExpr("~", $2)
	}
	| CHAN expr
	{
		expr := &ast.ChanOutExpr{Chan: $2}
		$$ = expr
		// fmt.Println("out chan")
	}

%%
