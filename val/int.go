package val

import (
	"strconv"
)

func NewInt(v int) Val {
	tmp := Int(v)
	return &tmp
}

//	型変換
func (v Int) Int() Int {
	return v
}

func (v Int) Float() Float {
	return Float(v)
}

func (v Int) String() String {
	return String(strconv.Itoa(int(v)))
}

func (_ Int) Type() String {
	return String("int")
}

func (v Int) Bool() Bool {
	if v != ZeroInt {
		return OneBool
	} else {
		return ZeroBool
	}
}

func (_ Int) Chan() (y Chan) {
	intPanic.NotSupport("chan")
	return
}

func (_ Int) Nil() (y Nil) {
	intPanic.NotSupport("nil")
	return
}

func (_ Int) Func() (y Func) {
	intPanic.NotSupport("func")
	return
}

//func (v Int) Slice() (y Slice) {
//	//intPanic.NotSupport("slice")
//	return NewSlice(&v)
//}

func (_ Int) Slice() (y Slice) {
	intPanic.NotSupport("slice")
	return
}

func (_ Int) IsNil() Bool {
	return false
}

//	組み込み関数
func (_ Int) Len() (x Int) {
	intPanic.NotSupport("len")
	return
}

func (_ Int) Cap() (dummy Int) {
	intPanic.NotSupport("cap")
	return
}

//func (v *Int) SetInt(x Int) {
//	*v = x
//}

//func (v *Int) SetString(x String) {
//	*v = x
//}

//func (v *Int) SelfAdd(x Val) {
//	*v += x.Int()
//}
//func (v *Int) SelfSub(x Val) {
//	*v -= x.Int()
//}
//func (v *Int) SelfMul(x Val) {
//	*v *= x.Int()
//}
//func (v *Int) SelfDiv(xx Val) {
//	x := xx.Int()
//	if x == 0 {
//		panic("devide by 0")
//	}
//	*v /= x
//}

//func (v Int) Add(x Val) Val {
//	v += x.Int()
//	//	v.SelfAdd(x)
//	return &v
//	//return v + x.Int()
//}
//func (v Int) Sub(x Val) Val {
//	v -= x.Int()
//	//v.SelfSub(x)
//	return &v
//	//return v - x.Int()
//}
//func (v Int) Mul(x Val) Val {
//	v *= x.Int()
//	//v.SelfMul(x)
//	return &v
//	//return v * x.Int()
//}
//func (v Int) Div(x Val) Val {
//	y := x.Int()
//	if y == 0 {
//		panic("devide by 0")
//	}
//	v /= y
//	//v.SelfDiv(x)
//	return &v
//	//return v / x.Int()
//}

//	四則演算等
func (v Int) Add(x Val) Val {
	v += x.Int()
	//	v.SelfAdd(x)
	return &v
	//return v + x.Int()
}
func (v Int) Sub(x Val) Val {
	v -= x.Int()
	//v.SelfSub(x)
	return &v
	//return v - x.Int()
}
func (v Int) Mul(x Val) Val {
	v *= x.Int()
	//v.SelfMul(x)
	return &v
	//return v * x.Int()
}

func (v Int) Div(x Val) Val {
	y := x.Int()
	if y == 0 {
		intPanic.DevidedByZero()
	}
	v /= y
	//v.SelfDiv(x)
	return &v
	//return v / x.Int()
}

func (v Int) Rem(x Val) Val {
	y := x.Int()
	if y == 0 {
		intPanic.DevidedByZero()
	}
	v %= y
	//v.SelfDiv(x)
	return &v
	//return v / x.Int()
}

func (v Int) Shl(x Val) Val {
	v <<= uint(x.Int())
	return &v
}

func (v Int) Shr(x Val) Val {
	v >>= uint(x.Int())
	return &v
}

func (v Int) BitAnd(x Val) Val {
	v &= x.Int()
	return &v
}

func (v Int) BitOr(x Val) Val {
	v |= x.Int()
	return &v
}

func (v Int) BitXor(x Val) Val {
	v ^= x.Int()
	return &v
}

//func (v Int) AddInt(x Int) Int {
//	return v + x
//}

//func (v Int) SubInt(x Int) Int {
//	return v - x
//}

//func (v *Int) AddInt(x Int) {
//	*v += x
//}

//func (v *Int) SubInt(x Int) {
//	*v -= x
//}

//func (v *Int) MulInt(x Int) {
//	*v *= x
//}

//func (v *Int) DivInt(x Int) {
//	if x == 0 {
//		panic("devide by 0")
//	}
//	*v /= x
//}

//	比較
func (v Int) Equal(x Val) Val {
	y := Bool(v == x.Int()) && !x.IsNil()
	return Val(&y)
}

func (v Int) NotEqual(x Val) Val {
	y := Bool(v != x.Int()) || x.IsNil()
	return Val(&y)
}

func (v Int) Less(x Val) Val {
	y := Bool(v < x.Int())
	return Val(&y)
}

func (v Int) Greater(x Val) Val {
	y := Bool(v > x.Int())
	return Val(&y)
}

func (v Int) LessOrEqual(x Val) Val {
	y := Bool(v <= x.Int())
	return Val(&y)
}

func (v Int) GreaterOrEqual(x Val) Val {
	y := Bool(v >= x.Int())
	return Val(&y)
}

//	条件
func (v Int) And(x Val) Val {
	y := Bool(v.Bool() && x.Bool())
	return Val(&y)
}

func (v Int) Or(x Val) Val {
	y := Bool(v.Bool() || x.Bool())
	return Val(&y)
}
func (v Int) Xor(x Val) Val {
	y := v.Int() ^ x.Int()
	return Val(&y)
}

func (v Int) Not() Val {
	if v == 0 {
		v = Int(1)
	} else {
		v = Int(0)
	}
	return &v
}

func (v Int) BitNot() Val {
	v = ^v
	return &v
}

func (v Int) Minus() Val {
	v = -v
	return &v
}

//func (v Int) Inc() {
//	v++
//}

//func (v Int) Dec() {
//	v--
//}

func (v *Int) Inc() {
	*v++
}

func (v *Int) Dec() {
	*v--
}

//	チャネル
func (_ Int) In(_ Val) {
	intPanic.NotSupport("in")
}

func (_ Int) Out() (y Val) {
	intPanic.NotSupport("out")
	return
}

func (_ Int) Close() {
	intPanic.NotSupport("close")
}

func (_ Int) Index(_ Int) (d Val) {
	intPanic.NotSupport("index")
	return
}

func (_ Int) Map(_ Val) (d Val) {
	intPanic.NotSupport("map")
	return
}

//	関数
func (_ Int) Call(_ ...Val) (y []Val) {
	intPanic.NotSupport("call")
	return
}

//func (_ Int) Stmts() (y []ast.Stmt) {
//	panic("int cannot stmts")
//	return
//}

//	インターフェースで定義されたメソッドをポインターメソッドとして定義している場合は、
//	構造体のポインタを interface 化しなければならない
func (v Int) Copy() Val {
	//	x := Int(v)
	//	return Val(&x)
	return &v
}

//func (v Int) Equal(v2 Val) Bool {
//	return Bool(v == v2)
//}
