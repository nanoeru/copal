package val

import ()

func NewNil() (x Nil) {
	return
}

func (v Nil) Int() (y Int) {
	//panic("nil cannot call int")
	return
}

func (v Nil) Float() (y Float) {
	//panic("nil cannot call float")
	return
}

func (v Nil) String() (y String) {
	return
}

func (_ Nil) Type() (y String) {
	return String("nil")
}

func (v Nil) Bool() (y Bool) {
	//	panic("nil cannot call bool")
	return
}

func (v Nil) Chan() (y Chan) {
	//panic("nil cannot call chan")
	return
}

func (v Nil) Nil() Nil {
	return v
}

func (v Nil) Func() (y Func) {
	//	panic("nil cannot call func")
	return
}

func (v Nil) Slice() (d Slice) {
	//	panic("nil cannot call slice")
	return
}

func (v Nil) IsNil() Bool {
	return true
}

func (_ Nil) Len() (d Int) {
	nilPanic.NotSupport("len")
	return
}

func (_ Nil) Cap() (dummy Int) {
	nilPanic.NotSupport("cap")
	return
}

func (_ Nil) Add(_ Val) (y Val) {
	nilPanic.NotSupport("+")
	return
}

func (_ Nil) Sub(_ Val) (y Val) {
	nilPanic.NotSupport("-")
	return
}

func (_ Nil) Mul(_ Val) (y Val) {
	nilPanic.NotSupport("*")
	return
}

func (_ Nil) Div(_ Val) (y Val) {
	nilPanic.NotSupport("/")
	return
}

func (_ Nil) Rem(_ Val) (y Val) {
	nilPanic.NotSupport("rem")
	return
}

func (_ Nil) Shl(_ Val) (y Val) {
	nilPanic.NotSupport("<<")
	return
}

func (_ Nil) Shr(_ Val) (y Val) {
	nilPanic.NotSupport(">>")
	return
}

func (v Nil) Equal(x Val) Val {
	return x.IsNil()
}

func (v Nil) NotEqual(x Val) Val {
	return !x.IsNil()
}

func (_ Nil) Less(_ Val) (y Val) {
	nilPanic.NotSupport("<")
	return
}

func (_ Nil) Greater(_ Val) (y Val) {
	nilPanic.NotSupport(">")
	return
}

func (_ Nil) LessOrEqual(_ Val) (y Val) {
	nilPanic.NotSupport("<=")
	return
}

func (v Nil) GreaterOrEqual(_ Val) (y Val) {
	nilPanic.NotSupport(">=")
	return
}

func (_ Nil) And(_ Val) (y Val) {
	nilPanic.NotSupport("&&")
	return
}
func (_ Nil) Or(_ Val) (y Val) {
	nilPanic.NotSupport("||")
	return
}
func (_ Nil) Xor(_ Val) (y Val) {
	nilPanic.NotSupport("^^")
	return
}
func (_ Nil) BitAnd(_ Val) (y Val) {
	nilPanic.NotSupport("&")
	return
}
func (_ Nil) BitOr(_ Val) (y Val) {
	nilPanic.NotSupport("|")
	return
}
func (_ Nil) BitXor(_ Val) (y Val) {
	nilPanic.NotSupport("^")
	return
}
func (_ Nil) Not() (y Val) {
	nilPanic.NotSupport("!")
	return
}
func (_ Nil) BitNot() (y Val) {
	nilPanic.NotSupport("~")
	return
}

func (_ Nil) Minus() (y Val) {
	nilPanic.NotSupport("-")
	return
}

func (_ Nil) Inc() {
	nilPanic.NotSupport("++")
	return
}

func (_ Nil) Dec() {
	nilPanic.NotSupport("--")
	return
}

func (_ Nil) Call(_ ...Val) (y []Val) {
	nilPanic.NotSupport("call")
	return
}

//	インターフェースで定義されたメソッドをポインターメソッドとして定義している場合は、
//	構造体のポインタを interface 化しなければならない
func (v Nil) Copy() Val {
	x := Nil(v)
	return Val(&x)
}

////func (_ Nil) Stmts() (y []ast.Stmt) {
////	panic("bool cannot stmts")
////	return
////}

//func (_ Nil) AddInt(_ Int) (y Int) {
//	panic("nil cannot addint")
//	return
//}

//func (_ Nil) SubInt(_ Int) (y Int) {
//	panic("nil cannot addint")
//	return
//}

func (_ Nil) In(_ Val) {
	nilPanic.NotSupport("in")
	return
}

func (_ Nil) Out() (y Val) {
	nilPanic.NotSupport("out")
	return
}

func (_ Nil) Close() {
	nilPanic.NotSupport("close")
	return
}

func (_ Nil) Index(_ Int) (d Val) {
	nilPanic.NotSupport("index")
	return
}

func (_ Nil) Map(_ Val) (d Val) {
	nilPanic.NotSupport("map")
	return
}
