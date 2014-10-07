package val

import (
	"strconv"
)

func (v Bool) Int() Int {
	if v {
		return 1
	} else {
		return 0
	}
}

func (v Bool) Float() Float {
	if v {
		return OneFloat
	} else {
		return ZeroFloat
	}
}

func (v Bool) String() String {
	return String(strconv.FormatBool(bool(v)))
}

func (_ Bool) Type() String {
	return String("bool")
}

func (v Bool) Bool() Bool {
	return v
}

func (v Bool) Chan() (y Chan) {
	boolPanic.NotSupport("chan")
	return
}

func (v Bool) Nil() (y Nil) {
	boolPanic.NotSupport("nil")
	return
}

func (v Bool) Func() (y Func) {
	boolPanic.NotSupport("func")
	return
}

func (v Bool) Slice() (d Slice) {
	boolPanic.NotSupport("slice")
	return
}

func (v Bool) IsNil() Bool {
	return false
}

func (v Bool) Len() Int {
	if v {
		return 1
	} else {
		return 0
	}
}

func (_ Bool) Cap() (dummy Int) {
	boolPanic.NotSupport("cap")
	return
}

//func (v *Bool) SelfAdd(x Val) {
//	boolPanic.NotSupport("add")
//}
//func (v *Bool) SelfSub(x Val) {
//	boolPanic.NotSupport("sub")
//}
//func (v *Bool) SelfMul(x Val) {
//	boolPanic.NotSupport("mul")
//}
//func (v *Bool) SelfDiv(xx Val) {
//	boolPanic.NotSupport("div")
//}

func (v Bool) Add(x Val) (y Val) {
	boolPanic.NotSupport("+")
	return
}
func (v Bool) Sub(x Val) (y Val) {
	boolPanic.NotSupport("-")
	return
}
func (v Bool) Mul(x Val) (y Val) {
	boolPanic.NotSupport("*")
	return
}
func (v Bool) Div(x Val) (y Val) {
	boolPanic.NotSupport("/")
	return
}
func (v Bool) Rem(x Val) (y Val) {
	boolPanic.NotSupport("%")
	return
}

func (_ Bool) Shl(_ Val) (y Val) {
	boolPanic.NotSupport("<<")
	return
}

func (_ Bool) Shr(_ Val) (y Val) {
	boolPanic.NotSupport(">>")
	return
}

func (_ Bool) BitAnd(_ Val) (y Val) {
	boolPanic.NotSupport("&")
	return
}

func (_ Bool) BitOr(_ Val) (y Val) {
	boolPanic.NotSupport("|")
	return
}

func (_ Bool) BitXor(_ Val) (y Val) {
	boolPanic.NotSupport("^")
	return
}

func (v Bool) Equal(x Val) Val {
	y := Bool(v == x.Bool()) && !x.IsNil()
	return Val(&y)
}

func (v Bool) NotEqual(x Val) Val {
	y := Bool(v != x.Bool()) || x.IsNil()
	return Val(&y)
}

func (_ Bool) Less(_ Val) (y Val) {
	boolPanic.NotSupport("<")
	return
}

func (_ Bool) Greater(_ Val) (y Val) {
	boolPanic.NotSupport(">")
	return
}

func (_ Bool) LessOrEqual(_ Val) (y Val) {
	boolPanic.NotSupport("<=")
	return
}

func (_ Bool) GreaterOrEqual(_ Val) (y Val) {
	boolPanic.NotSupport(">=")
	return
}

func (v Bool) And(x Val) Val {
	y := Bool(v && x.Bool())
	return Val(&y)
}
func (v Bool) Or(x Val) Val {
	y := Bool(v || x.Bool())
	return Val(&y)
}
func (v Bool) Xor(x Val) Val {
	y := Bool(v != x.Bool())
	return Val(&y)
}

func (v Bool) Not() Val {
	y := Bool(!v)
	return Val(&y)
}

func (v Bool) BitNot() Val {
	y := Bool(!v).Int()
	return Val(&y)
}

func (_ Bool) Minus() (dummy Val) {
	boolPanic.NotSupport("-")
	return
}

func (_ Bool) Inc() {
	boolPanic.NotSupport("++")
	return
}

func (_ Bool) Dec() {
	boolPanic.NotSupport("--")
	return
}

func (_ Bool) Call(_ ...Val) (y []Val) {
	boolPanic.NotSupport("call")
	return
}

//func (_ Bool) Stmts() (y []ast.Stmt) {
//	panic("bool cannot stmts")
//	return
//}

//func (_ Bool) AddInt(_ Int) (y Int) {
//	panic("bool cannot addint")
//	return
//}

//func (_ Bool) SubInt(_ Int) (y Int) {
//	panic("bool cannot addint")
//	return
//}

func (_ Bool) In(_ Val) {
	boolPanic.NotSupport("in")
}

func (_ Bool) Out() (_ Val) {
	boolPanic.NotSupport("out")
	return
}

func (_ Bool) Close() {
	boolPanic.NotSupport("close")
	return
}

func (_ Bool) Index(_ Int) (d Val) {
	boolPanic.NotSupport("index")
	return
}

func (_ Bool) Map(_ Val) (d Val) {
	boolPanic.NotSupport("map")
	return
}

//	インターフェースで定義されたメソッドをポインターメソッドとして定義している場合は、
//	構造体のポインタを interface 化しなければならない
func (v Bool) Copy() Val {
	//	x := Bool(v)
	//	return Val(&x)
	return &v
}
