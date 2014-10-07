package val

import (
	"reflect"
)

func NewFunc(v func(...Val) []Val) *Func {
	//	TODO	きちんと実装
	//x := Int(0)
	//return &Func{
	//	F:   v,
	//	Val: &x,
	//	//		Stmts: []ast.Stmt,
	//}
	return &Func{
		F:   v,
		Val: nil,
	}
}

func (v Func) Func() Func {
	return v
}

func (v *Func) SetStmt(stmt interface{}) {
	v.Stmt = stmt
}
func (v *Func) SetStmts(stmts interface{}) {
	v.Stmts = stmts
}

func (v Func) Call(args ...Val) []Val {
	return v.F(args...)
}

func (v Func) Copy() Val {
	return v
}

func (_ Func) Int() (d Int) {
	funcPanic.NotSupport("int")
	return
}

func (_ Func) Float() (d Float) {
	funcPanic.NotSupport("float")
	return
}

func (_ Func) String() (d String) {
	funcPanic.NotSupport("string")
	return
}

func (_ Func) Type() String {
	return String("func")
}

func (_ Func) Bool() (d Bool) {
	funcPanic.NotSupport("bool")
	return
}

func (_ Func) Chan() (d Chan) {
	funcPanic.NotSupport("chan")
	return
}

func (_ Func) Nil() (d Nil) {
	funcPanic.NotSupport("nil")
	return
}

func (_ Func) Slice() (d Slice) {
	funcPanic.NotSupport("func")
	return
}

func (_ Func) IsNil() Bool {
	return false
}

func (_ Func) Len() (d Int) {
	funcPanic.NotSupport("len")
	return
}

func (_ Func) Cap() (d Int) {
	funcPanic.NotSupport("cap")
	return
}

func (_ Func) Add(_ Val) (d Val) {
	funcPanic.NotSupport("+")
	return
}

func (_ Func) Sub(_ Val) (d Val) {
	funcPanic.NotSupport("-")
	return
}

func (_ Func) Mul(_ Val) (d Val) {
	funcPanic.NotSupport("*")
	return
}

func (_ Func) Div(_ Val) (d Val) {
	funcPanic.NotSupport("/")
	return
}

func (_ Func) Rem(_ Val) (d Val) {
	funcPanic.NotSupport("%")
	return
}

func (_ Func) Shl(_ Val) (d Val) {
	funcPanic.NotSupport("<<")
	return
}

func (_ Func) Shr(_ Val) (d Val) {
	funcPanic.NotSupport(">>")
	return
}

func (v Func) Equal(x Val) Val {
	y := Bool(reflect.DeepEqual(v, x.Func())) && !x.IsNil()
	return Val(&y)
}

func (v Func) NotEqual(x Val) Val {
	y := Bool(!reflect.DeepEqual(v, x.Func())) || x.IsNil()
	return Val(&y)
}

func (_ Func) Less(_ Val) (d Val) {
	funcPanic.NotSupport("<")
	return
}

func (_ Func) Greater(_ Val) (d Val) {
	funcPanic.NotSupport(">")
	return
}

func (_ Func) LessOrEqual(_ Val) (d Val) {
	funcPanic.NotSupport("<=")
	return
}

func (_ Func) GreaterOrEqual(_ Val) (d Val) {
	funcPanic.NotSupport(">=")
	return
}

func (_ Func) And(_ Val) (d Val) {
	funcPanic.NotSupport("&&")
	return
}

func (_ Func) Or(_ Val) (d Val) {
	funcPanic.NotSupport("||")
	return
}

func (_ Func) Xor(_ Val) (d Val) {
	funcPanic.NotSupport("^^")
	return
}
func (_ Func) BitAnd(_ Val) (d Val) {
	funcPanic.NotSupport("&")
	return
}

func (_ Func) BitOr(_ Val) (d Val) {
	funcPanic.NotSupport("|")
	return
}

func (_ Func) BitXor(_ Val) (d Val) {
	funcPanic.NotSupport("^")
	return
}

func (_ Func) Not() (d Val) {
	funcPanic.NotSupport("!")
	return
}
func (_ Func) BitNot() (d Val) {
	funcPanic.NotSupport("~")
	return
}

func (_ Func) Minus() (d Val) {
	funcPanic.NotSupport("-")
	return
}

func (_ Func) Inc() {
	funcPanic.NotSupport("++")
	return
}

func (_ Func) Dec() {
	funcPanic.NotSupport("--")
	return
}

func (_ Func) In(_ Val) {
	funcPanic.NotSupport("in")
}

func (_ Func) Out() (y Val) {
	funcPanic.NotSupport("out")
	return
}

func (_ Func) Close() {
	funcPanic.NotSupport("close")
}

func (_ Func) Index(x Int) (y Val) {
	funcPanic.NotSupport("[]")
	return
}

func (_ Func) Map(key Val) (ret Val) {
	funcPanic.NotSupport("close")
	return
}
