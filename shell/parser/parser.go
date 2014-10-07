//line parser.go.y:1
package parser

import __yyfmt__ "fmt"

//line parser.go.y:3
import (
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/common/token"
)

//	ダミーデータ用の構造体
type Empty struct{}

//line parser.go.y:26
type yySymType struct {
	yys int
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
	tok token.Token
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const EOF = 57349
const NEW_ASSIGN = 57350
const ADD_ASSIGN = 57351
const SUB_ASSIGN = 57352
const MUL_ASSIGN = 57353
const DIV_ASSIGN = 57354
const REM_ASSIGN = 57355
const XOR_ASSIGN = 57356
const INC = 57357
const DEC = 57358
const VAR = 57359
const FUNC = 57360
const RETURN = 57361
const IF = 57362
const NIL = 57363
const FOR = 57364
const IN = 57365
const ELSE = 57366
const BREAK = 57367
const CONTINUE = 57368
const TRUE = 57369
const FALSE = 57370
const GO = 57371
const XOR = 57372
const BLANK = 57373
const SHELL = 57374
const COMMENT = 57375
const RANGE = 57376
const CHAN = 57377
const OR = 57378
const AND = 57379
const LE = 57380
const GE = 57381
const NE = 57382
const EQ = 57383
const SHR = 57384
const SHL = 57385
const UNARY = 57386
const LEFT = 57387

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"STRING",
	"EOF",
	"NEW_ASSIGN",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"REM_ASSIGN",
	"XOR_ASSIGN",
	"INC",
	"DEC",
	"VAR",
	"FUNC",
	"RETURN",
	"IF",
	"NIL",
	"FOR",
	"IN",
	"ELSE",
	"BREAK",
	"CONTINUE",
	"TRUE",
	"FALSE",
	"GO",
	"XOR",
	"BLANK",
	"SHELL",
	"COMMENT",
	"RANGE",
	"CHAN",
	"OR",
	"AND",
	" |",
	" &",
	" ^",
	" <",
	" >",
	"LE",
	"GE",
	"NE",
	"EQ",
	" +",
	" -",
	" *",
	" /",
	" %",
	"SHR",
	"SHL",
	"UNARY",
	" !",
	" ~",
	" [",
	" ]",
	" (",
	" )",
	" .",
	"LEFT",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:468

//line yacctab:1
var yyExca = []int{
	-1, 0,
	1, 1,
	-2, 28,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	64, 1,
	-2, 28,
	-1, 21,
	8, 29,
	65, 29,
	68, 29,
	-2, 39,
	-1, 65,
	64, 1,
	-2, 28,
	-1, 138,
	7, 23,
	33, 23,
	63, 23,
	66, 23,
	67, 23,
	-2, 42,
}

const yyNprod = 71
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 681

var yyAct = []int{

	69, 5, 3, 5, 61, 6, 115, 109, 60, 62,
	59, 63, 72, 150, 64, 110, 106, 149, 139, 57,
	104, 58, 115, 125, 75, 58, 76, 77, 78, 79,
	65, 68, 80, 117, 1, 13, 82, 29, 84, 85,
	86, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 101, 102, 83, 144, 137,
	67, 12, 65, 65, 112, 107, 5, 108, 21, 16,
	24, 13, 56, 113, 66, 58, 56, 103, 140, 58,
	138, 135, 9, 10, 7, 18, 8, 74, 115, 58,
	127, 17, 19, 11, 14, 15, 120, 12, 58, 28,
	111, 35, 122, 36, 115, 34, 114, 121, 49, 48,
	123, 105, 23, 35, 81, 36, 128, 34, 126, 26,
	27, 22, 73, 25, 4, 134, 131, 132, 30, 67,
	14, 15, 136, 20, 5, 2, 141, 129, 130, 142,
	143, 0, 0, 145, 0, 146, 147, 0, 0, 0,
	148, 52, 0, 0, 151, 152, 0, 51, 50, 54,
	53, 55, 42, 43, 44, 45, 47, 46, 37, 38,
	39, 40, 41, 49, 48, 32, 33, 0, 35, 0,
	36, 0, 34, 0, 0, 0, 0, 133, 0, 0,
	52, 0, 0, 0, 0, 31, 51, 50, 54, 53,
	55, 42, 43, 44, 45, 47, 46, 37, 38, 39,
	40, 41, 49, 48, 21, 16, 24, 35, 0, 36,
	0, 34, 0, 65, 0, 0, 0, 0, 9, 10,
	7, 18, 8, 0, 0, 0, 0, 17, 19, 11,
	39, 40, 41, 49, 48, 28, 0, 0, 35, 0,
	36, 0, 34, 0, 0, 0, 0, 0, 23, 0,
	21, 16, 24, 0, 0, 26, 27, 22, 0, 25,
	0, 0, 0, 65, 9, 10, 7, 18, 8, 20,
	0, 0, 0, 17, 19, 11, 0, 0, 0, 0,
	0, 28, 0, 37, 38, 39, 40, 41, 49, 48,
	0, 0, 0, 35, 23, 36, 0, 34, 0, 0,
	0, 26, 27, 22, 0, 25, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 20, 51, 50, 54, 53,
	55, 42, 43, 44, 45, 47, 46, 37, 38, 39,
	40, 41, 49, 48, 32, 33, 0, 35, 0, 36,
	0, 34, 0, 65, 0, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 31, 51, 50, 54, 53, 55,
	42, 43, 44, 45, 47, 46, 37, 38, 39, 40,
	41, 49, 48, 70, 16, 24, 35, 0, 36, 0,
	34, 0, 0, 0, 0, 0, 0, 71, 0, 0,
	18, 0, 0, 0, 0, 0, 17, 19, 0, 0,
	0, 0, 0, 124, 28, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 23, 0, 0,
	0, 0, 0, 0, 26, 27, 22, 0, 25, 0,
	0, 0, 0, 52, 0, 0, 0, 0, 20, 51,
	50, 54, 53, 55, 42, 43, 44, 45, 47, 46,
	37, 38, 39, 40, 41, 49, 48, 0, 0, 52,
	35, 119, 36, 0, 34, 51, 50, 54, 53, 55,
	42, 43, 44, 45, 47, 46, 37, 38, 39, 40,
	41, 49, 48, 0, 0, 52, 35, 0, 36, 118,
	34, 51, 50, 54, 53, 55, 42, 43, 44, 45,
	47, 46, 37, 38, 39, 40, 41, 49, 48, 70,
	16, 24, 35, 0, 36, 0, 34, 0, 0, 0,
	0, 0, 0, 71, 0, 0, 18, 0, 0, 0,
	0, 0, 17, 19, 0, 0, 0, 0, 0, 0,
	28, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 23, 0, 0, 0, 0, 0, 0,
	26, 27, 22, 0, 25, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 20, 51, 50, 54, 53, 55,
	42, 43, 44, 45, 47, 46, 37, 38, 39, 40,
	41, 49, 48, 0, 0, 0, 35, 0, 116, 0,
	34, 50, 54, 53, 55, 42, 43, 44, 45, 47,
	46, 37, 38, 39, 40, 41, 49, 48, 0, 0,
	0, 35, 0, 36, 0, 34, 54, 53, 55, 42,
	43, 44, 45, 47, 46, 37, 38, 39, 40, 41,
	49, 48, 0, 0, 0, 35, 0, 36, 0, 34,
	42, 43, 44, 45, 47, 46, 37, 38, 39, 40,
	41, 49, 48, 0, 0, 0, 35, 0, 36, 0,
	34,
}
var yyPact = []int{

	64, -1000, 64, 28, -1000, 329, 11, 256, 210, 70,
	515, 515, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	118, -1000, 29, 515, -1000, 515, 515, 515, 515, -1000,
	-1000, 515, -1000, -1000, 110, 515, 515, 515, 515, 515,
	515, 515, 515, 515, 515, 515, 515, 515, 515, 515,
	515, 515, 515, 515, 515, 515, 515, 515, 107, -50,
	160, -1000, 160, -59, 7, 64, 5, 102, -62, 465,
	-1000, 1, 549, -1000, -30, 44, 439, 44, 44, 465,
	465, -1000, 413, 36, 191, 191, 56, 56, 56, 246,
	246, 246, 246, 246, 246, 44, 44, 598, 574, 574,
	619, 619, 619, -62, -62, -1000, 515, 78, -1000, 515,
	379, -41, 102, 30, -1000, 515, 515, 515, -1000, -1000,
	-1000, 290, -33, 121, 515, -1000, 21, 0, 465, 20,
	-46, 54, -1000, 256, 290, -1, -1000, 102, -1000, -1000,
	-33, -33, -1000, -1000, 102, -43, -1000, -1000, -47, -33,
	-33, -1000, -1000,
}
var yyPgo = []int{

	0, 4, 34, 135, 2, 124, 0, 20, 5,
}
var yyR1 = []int{

	0, 2, 2, 3, 3, 1, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 5, 5, 5, 5, 8, 8,
	8, 7, 7, 7, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6,
}
var yyR2 = []int{

	0, 0, 2, 2, 1, 3, 1, 3, 3, 3,
	2, 2, 5, 3, 7, 5, 2, 3, 7, 6,
	6, 9, 2, 5, 1, 1, 1, 1, 0, 1,
	3, 0, 1, 3, 1, 1, 1, 1, 2, 1,
	3, 4, 4, 5, 5, 8, 2, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2,
}
var yyChk = []int{

	-1000, -2, -3, -4, -5, -6, -8, 20, 22, 18,
	19, 29, 33, 7, 66, 67, 5, 27, 21, 28,
	69, 4, 57, 48, 6, 59, 55, 56, 35, -2,
	-5, 35, 15, 16, 61, 57, 59, 47, 48, 49,
	50, 51, 41, 42, 43, 44, 46, 45, 53, 52,
	37, 36, 30, 39, 38, 40, 65, 8, 68, -4,
	-6, -1, -6, -4, -8, 63, 4, 59, -7, -6,
	4, 18, -6, 4, 58, -6, -6, -6, -6, -6,
	-6, 4, -6, -7, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -7, -7, 4, 66, -1, -1, 66,
	8, -2, 59, -8, 4, 68, 59, 63, 60, 58,
	60, -6, 24, -6, 34, 64, -8, 60, -6, -7,
	-7, -1, -1, 66, -6, 60, -1, 59, 60, 64,
	24, -4, -1, -1, 59, -8, -1, -1, -8, 60,
	60, -1, -1,
}
var yyDef = []int{

	-2, -2, -2, 0, 4, 6, 0, 28, 28, 0,
	31, 0, 24, 25, 26, 27, 34, 35, 36, 37,
	0, -2, 0, 0, 47, 0, 0, 0, 0, 2,
	3, 0, 10, 11, 0, 0, 31, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 31, 31, 0, 0,
	6, 16, 6, 0, 0, -2, 0, 28, 22, 32,
	39, 0, 0, 38, 0, 46, 0, 68, 69, 70,
	9, 40, 0, 0, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 7, 8, 30, 0, 13, 17, 0,
	31, 0, 28, 0, 29, 0, 31, 31, 48, 41,
	42, 0, 0, 0, 0, 5, 0, 0, 33, 0,
	0, 12, 15, 28, 0, 0, 44, 28, -2, 43,
	0, 0, 19, 20, 28, 0, 14, 18, 0, 0,
	0, 45, 21,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	67, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 69, 51, 39, 3,
	59, 60, 49, 47, 68, 48, 61, 50, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 66,
	41, 65, 42, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 57, 3, 58, 40, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 63, 38, 64, 56,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 43, 44, 45, 46,
	52, 53, 54, 62,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:80
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 2:
		//line parser.go.y:87
		{
			yyVAL.stmts = append([]ast.Stmt{yyS[yypt-1].stmt}, yyS[yypt-0].stmts...)
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 3:
		//line parser.go.y:96
		{
			yyVAL.stmt = yyS[yypt-1].pre_stmt
		}
	case 4:
		//line parser.go.y:100
		{
			yyVAL.stmt = nil
		}
	case 5:
		//line parser.go.y:107
		{
			yyVAL.stmts = yyS[yypt-1].stmts
		}
	case 6:
		//line parser.go.y:114
		{
			lit := yyS[yypt-0].expr.GetLit()
			stmt := &ast.ExprStmt{Expr: yyS[yypt-0].expr}
			stmt.SetLit(lit)
			yyVAL.pre_stmt = stmt
		}
	case 7:
		//line parser.go.y:121
		{
			stmt := &ast.AssignStmt{IdentExprs: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
			yyVAL.pre_stmt = stmt
		}
	case 8:
		//line parser.go.y:126
		{
			stmt := &ast.NewAssignStmt{IdentExprs: yyS[yypt-2].idents, Exprs: yyS[yypt-0].exprs}
			yyVAL.pre_stmt = stmt
		}
	case 9:
		//line parser.go.y:132
		{
			stmt := &ast.ChanInStmt{Chan: yyS[yypt-2].expr, Expr: yyS[yypt-0].expr}
			yyVAL.pre_stmt = stmt
		}
	case 10:
		//line parser.go.y:152
		{
			yyVAL.pre_stmt = &ast.IncStmt{Expr: yyS[yypt-1].expr}
		}
	case 11:
		//line parser.go.y:157
		{
			yyVAL.pre_stmt = &ast.DecStmt{Expr: yyS[yypt-1].expr}
		}
	case 12:
		//line parser.go.y:162
		{
			yyVAL.pre_stmt = &ast.IfStmt{Init: yyS[yypt-3].pre_stmt, Cond: yyS[yypt-1].expr, ThenStmts: yyS[yypt-0].stmts}
		}
	case 13:
		//line parser.go.y:167
		{
			yyVAL.pre_stmt = &ast.IfStmt{Cond: yyS[yypt-1].expr, ThenStmts: yyS[yypt-0].stmts}
		}
	case 14:
		//line parser.go.y:172
		{
			yyVAL.pre_stmt = &ast.IfStmt{Init: yyS[yypt-5].pre_stmt, Cond: yyS[yypt-3].expr, ThenStmts: yyS[yypt-2].stmts, ElseStmts: yyS[yypt-0].stmts}
		}
	case 15:
		//line parser.go.y:177
		{
			yyVAL.pre_stmt = &ast.IfStmt{Cond: yyS[yypt-3].expr, ThenStmts: yyS[yypt-2].stmts, ElseStmts: yyS[yypt-0].stmts}
		}
	case 16:
		//line parser.go.y:182
		{
			yyVAL.pre_stmt = &ast.ForStmt{Stmts: yyS[yypt-0].stmts, Pattern: ast.InfiniteForLoop}
		}
	case 17:
		//line parser.go.y:187
		{
			yyVAL.pre_stmt = &ast.ForStmt{Cond: yyS[yypt-1].expr, Stmts: yyS[yypt-0].stmts, Pattern: ast.CondForLoop}
		}
	case 18:
		//line parser.go.y:192
		{
			yyVAL.pre_stmt = &ast.ForStmt{Init: yyS[yypt-5].pre_stmt, Cond: yyS[yypt-3].expr, Post: yyS[yypt-1].pre_stmt, Stmts: yyS[yypt-0].stmts, Pattern: ast.InitCondPostForLoop}
		}
	case 19:
		//line parser.go.y:196
		{
			yyVAL.pre_stmt = &ast.ForStmt{Idents: yyS[yypt-4].idents, Expr: yyS[yypt-1].expr, Stmts: yyS[yypt-0].stmts, Pattern: ast.KeyValueRangeLoop}
		}
	case 20:
		//line parser.go.y:200
		{
			yyVAL.pre_stmt = &ast.FuncStmt{Name: yyS[yypt-4].tok.Lit, Args: yyS[yypt-2].idents, Stmts: yyS[yypt-0].stmts}
		}
	case 21:
		//line parser.go.y:204
		{
			yyVAL.pre_stmt = &ast.FuncStmt{Name: yyS[yypt-7].tok.Lit, Args: yyS[yypt-5].idents, RetVals: yyS[yypt-2].idents, Stmts: yyS[yypt-0].stmts}
		}
	case 22:
		//line parser.go.y:208
		{
			yyVAL.pre_stmt = &ast.ReturnStmt{Exprs: yyS[yypt-0].exprs}
		}
	case 23:
		//line parser.go.y:213
		{
			yyVAL.pre_stmt = &ast.GoCallStmt{Expr: yyS[yypt-3].expr, ArgExprs: yyS[yypt-1].exprs}
		}
	case 24:
		//line parser.go.y:220
		{
			yyVAL.stmt_delimiter = &Empty{}
		}
	case 25:
		//line parser.go.y:224
		{
			yyVAL.stmt_delimiter = &Empty{}
		}
	case 26:
		//line parser.go.y:228
		{
			yyVAL.stmt_delimiter = &Empty{}
		}
	case 27:
		//line parser.go.y:232
		{
			yyVAL.stmt_delimiter = &Empty{}
		}
	case 28:
		//line parser.go.y:269
		{
			yyVAL.idents = []string{}
		}
	case 29:
		//line parser.go.y:273
		{
			yyVAL.idents = []string{yyS[yypt-0].tok.Lit}
		}
	case 30:
		//line parser.go.y:277
		{
			yyVAL.idents = append(yyS[yypt-2].idents, yyS[yypt-0].tok.Lit)
		}
	case 31:
		//line parser.go.y:283
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 32:
		//line parser.go.y:287
		{
			yyVAL.exprs = []ast.Expr{yyS[yypt-0].expr}
		}
	case 33:
		//line parser.go.y:291
		{
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 34:
		//line parser.go.y:306
		{
			yyVAL.expr = NumberExpr(yyS[yypt-0].tok)
		}
	case 35:
		//line parser.go.y:311
		{
			yyVAL.expr = BoolExpr(true, yyS[yypt-0].tok)
		}
	case 36:
		//line parser.go.y:316
		{
			yyVAL.expr = NilExpr(yyS[yypt-0].tok)
		}
	case 37:
		//line parser.go.y:321
		{
			yyVAL.expr = BoolExpr(false, yyS[yypt-0].tok)
		}
	case 38:
		//line parser.go.y:325
		{
			//	廃止?!
			yyVAL.expr = IdentExpr(yyS[yypt-0].tok)
		}
	case 39:
		//line parser.go.y:330
		{
			yyVAL.expr = IdentExpr(yyS[yypt-0].tok)
		}
	case 40:
		//line parser.go.y:335
		{
			yyVAL.expr = &ast.MapAccessExpr{MapExpr: yyS[yypt-2].expr, Name: yyS[yypt-0].tok.Lit}
		}
	case 41:
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.MapAccessExpr{MapExpr: yyS[yypt-3].expr, Expr: yyS[yypt-1].expr}
		}
	case 42:
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.FuncCallExpr{Expr: yyS[yypt-3].expr, ArgExprs: yyS[yypt-1].exprs}
		}
	case 43:
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.SliceExpr{Exprs: yyS[yypt-1].exprs}
		}
	case 44:
		//line parser.go.y:354
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-2].idents, Stmts: yyS[yypt-0].stmts}
		}
	case 45:
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyS[yypt-5].idents, RetVals: yyS[yypt-2].idents, Stmts: yyS[yypt-0].stmts}
		}
	case 46:
		//line parser.go.y:363
		{
			yyVAL.expr = UnaryOpExpr("-", yyS[yypt-0].expr)
		}
	case 47:
		//line parser.go.y:367
		{
			yyVAL.expr = StringExpr(yyS[yypt-0].tok)
		}
	case 48:
		//line parser.go.y:371
		{
			lit := "( " + yyS[yypt-1].expr.GetLit() + " )"
			yyS[yypt-1].expr.SetLit(lit)
			yyVAL.expr = yyS[yypt-1].expr
			// fmt.Println("( expr )", lit)
		}
	case 49:
		//line parser.go.y:378
		{
			yyVAL.expr = BinOpExpr("+", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 50:
		//line parser.go.y:382
		{
			yyVAL.expr = BinOpExpr("-", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 51:
		//line parser.go.y:386
		{
			yyVAL.expr = BinOpExpr("*", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 52:
		//line parser.go.y:390
		{
			yyVAL.expr = BinOpExpr("/", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 53:
		//line parser.go.y:394
		{
			yyVAL.expr = BinOpExpr("%", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 54:
		//line parser.go.y:398
		{
			yyVAL.expr = BinOpExpr("<", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 55:
		//line parser.go.y:402
		{
			yyVAL.expr = BinOpExpr(">", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 56:
		//line parser.go.y:406
		{
			yyVAL.expr = BinOpExpr("<=", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 57:
		//line parser.go.y:410
		{
			yyVAL.expr = BinOpExpr(">=", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 58:
		//line parser.go.y:414
		{
			yyVAL.expr = BinOpExpr("==", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 59:
		//line parser.go.y:418
		{
			yyVAL.expr = BinOpExpr("!=", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 60:
		//line parser.go.y:422
		{
			yyVAL.expr = BinOpExpr("<<", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 61:
		//line parser.go.y:426
		{
			yyVAL.expr = BinOpExpr(">>", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 62:
		//line parser.go.y:430
		{
			yyVAL.expr = BinOpExpr("&&", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 63:
		//line parser.go.y:434
		{
			yyVAL.expr = BinOpExpr("||", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 64:
		//line parser.go.y:438
		{
			yyVAL.expr = BinOpExpr("^^", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 65:
		//line parser.go.y:442
		{
			yyVAL.expr = BinOpExpr("&", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 66:
		//line parser.go.y:446
		{
			yyVAL.expr = BinOpExpr("|", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 67:
		//line parser.go.y:450
		{
			yyVAL.expr = BinOpExpr("^", yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 68:
		//line parser.go.y:454
		{
			yyVAL.expr = UnaryOpExpr("!", yyS[yypt-0].expr)
		}
	case 69:
		//line parser.go.y:458
		{
			yyVAL.expr = UnaryOpExpr("~", yyS[yypt-0].expr)
		}
	case 70:
		//line parser.go.y:462
		{
			expr := &ast.ChanOutExpr{Chan: yyS[yypt-0].expr}
			yyVAL.expr = expr
			// fmt.Println("out chan")
		}
	}
	goto yystack /* stack new state and value */
}
