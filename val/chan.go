package val

import (
//"fmt"
)

func NewChan(n int) *Chan {
	ch := Chan(make(chan Val, n))
	return &ch
}

func (_ Chan) Int() (y Int) {
	chanPanic.NotSupport("int")
	return
}

func (_ Chan) Float() (y Float) {
	chanPanic.NotSupport("float")
	return
}

func (v Chan) String() String {
	return String("chan")
}

func (_ Chan) Type() String {
	return String("chan")
}

func (_ Chan) Bool() (y Bool) {
	chanPanic.NotSupport("bool")
	return
}

//func (v *Chan) Chan() Chan {
//	return *v
//}
func (v Chan) Chan() Chan {
	return v
}

func (_ Chan) Nil() (y Nil) {
	chanPanic.NotSupport("nil")
	return
}

func (_ Chan) Func() (y Func) {
	chanPanic.NotSupport("func")
	return
}

func (_ Chan) Slice() (d Slice) {
	chanPanic.NotSupport("slice")
	return
}

//func (v *Chan) Len() Int {
//	return Int(len(*v))
//}

//func (v *Chan) Cap() Int {
//	return Int(cap(*v))
//}
func (v Chan) Len() Int {
	return Int(len(v))
}

func (v Chan) Cap() Int {
	return Int(cap(v))
}

func (_ Chan) Add(_ Val) (y Val) {
	chanPanic.NotSupport("+")
	return
}

func (_ Chan) Sub(_ Val) (y Val) {
	chanPanic.NotSupport("-")
	return
}

func (_ Chan) Mul(_ Val) (y Val) {
	chanPanic.NotSupport("*")
	return
}

func (_ Chan) Div(_ Val) (y Val) {
	chanPanic.NotSupport("/")
	return
}

func (_ Chan) Rem(_ Val) (y Val) {
	chanPanic.NotSupport("%")
	return
}

func (_ Chan) Shl(_ Val) (y Val) {
	chanPanic.NotSupport("<<")
	return
}

func (_ Chan) Shr(_ Val) (y Val) {
	chanPanic.NotSupport(">>")
	return
}

func (v Chan) Equal(x Val) Val {
	y := Bool(v == x.Chan()) && !x.IsNil()
	return Val(&y)
}

func (v Chan) NotEqual(x Val) Val {
	y := Bool(v != x.Chan()) || x.IsNil()
	return Val(&y)
}

func (_ Chan) Less(_ Val) (y Val) {
	chanPanic.NotSupport("<")
	return
}

func (_ Chan) Greater(_ Val) (y Val) {
	chanPanic.NotSupport(">")
	return
}

func (_ Chan) LessOrEqual(_ Val) (y Val) {
	chanPanic.NotSupport("<=")
	return
}

func (_ Chan) GreaterOrEqual(_ Val) (y Val) {
	chanPanic.NotSupport(">=")
	return
}

func (_ Chan) And(_ Val) (y Val) {
	chanPanic.NotSupport("&&")
	return
}
func (_ Chan) Or(_ Val) (y Val) {
	chanPanic.NotSupport("||")
	return
}
func (_ Chan) Xor(_ Val) (y Val) {
	chanPanic.NotSupport("^")
	return
}
func (_ Chan) BitAnd(_ Val) (y Val) {
	chanPanic.NotSupport("&")
	return
}
func (_ Chan) BitOr(_ Val) (y Val) {
	chanPanic.NotSupport("|")
	return
}
func (_ Chan) BitXor(_ Val) (y Val) {
	chanPanic.NotSupport("^")
	return
}
func (_ Chan) Not() (y Val) {
	chanPanic.NotSupport("!")
	return
}

func (_ Chan) BitNot() (y Val) {
	chanPanic.NotSupport("~")
	return
}

func (_ Chan) Minus() (dummy Val) {
	chanPanic.NotSupport("-")
	return
}

func (_ Chan) Inc() {
	chanPanic.NotSupport("++")
	return
}

func (_ Chan) Dec() {
	chanPanic.NotSupport("--")
	return
}

func (_ Chan) Call(_ ...Val) (y []Val) {
	chanPanic.NotSupport("call")
	return
}

//	インターフェースで定義されたメソッドをポインターメソッドとして定義している場合は、
//	構造体のポインタを interface 化しなければならない
//func (v Chan) Copy() Val {
//	x := Chan(v)
//	return Val(&x)
//}
//func (v *Chan) Copy() Val {
//	return Val(v)
//}
func (v Chan) Copy() Val {
	return &v
}

func (v Chan) Copy2() Val {
	return Val(v)
}

////func (_ Chan) Stmts() (y []ast.Stmt) {
////	panic("bool cannot stmts")
////	return
////}

//func (_ Chan) AddInt(_ Int) (y Int) {
//	panic("chan cannot addint")
//	return
//}

//func (_ Chan) SubInt(_ Int) (y Int) {
//	panic("chan cannot addint")
//	return
//}

func (v Chan) IsNil() Bool {
	return false
}

//func (v *Chan) In(x Val) {
//	*v <- x.Copy()
//}

func (v Chan) In(x Val) {
	v <- x.Copy()
}

//	Int の場合は &Int のアドレスのコピーが渡る
//func (v *Chan) In(x Val) {
//	*v <- x
//}

//func (v *Chan) Out() Val {
//	tmp, ok := <-*v
//	if ok {
//		return tmp
//	} else {
//		return NewNil()
//	}
//}
func (v Chan) Out() Val {
	tmp, ok := <-v
	if ok {
		return tmp
	} else {
		return NewNil()
	}
}

func (v Chan) Close() {
	close(v)
}

func (_ Chan) Index(_ Int) (d Val) {
	chanPanic.NotSupport("index")
	return
}

func (_ Chan) Map(_ Val) (d Val) {
	chanPanic.NotSupport("map")
	return
}
