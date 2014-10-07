package val

import (
	"strconv"
)

func NewFloat(v float64) Val {
	tmp := Float(v)
	return &tmp
}

//	型変換
func (v Float) Int() Int {
	return Int(v)
}

func (v Float) Float() Float {
	return v
}

func (v Float) String() String {
	return String(strconv.FormatFloat(float64(v), 'f', -1, 64))
}

func (_ Float) Type() String {
	return String("float")
}

func (v Float) Bool() Bool {
	if v != ZeroFloat {
		return OneBool
	} else {
		return ZeroBool
	}
}

func (_ Float) Chan() (y Chan) {
	floatPanic.NotSupport("chan")
	return
}

func (_ Float) Nil() (y Nil) {
	floatPanic.NotSupport("nil")
	return
}

func (_ Float) Func() (y Func) {
	floatPanic.NotSupport("func")
	return
}

//func (v Float) Slice() (y Slice) {
//	//floatPanic.NotSupport("slice")
//	return NewSlice(&v)
//}

func (_ Float) Slice() (y Slice) {
	floatPanic.NotSupport("slice")
	return
}

func (_ Float) IsNil() Bool {
	return false
}

//	組み込み関数
func (_ Float) Len() (x Int) {
	floatPanic.NotSupport("len")
	return
}

func (_ Float) Cap() (dummy Int) {
	floatPanic.NotSupport("cap")
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

//func (v Float) Add(x Val) Val {
//	v += x.Int()
//	//	v.SelfAdd(x)
//	return &v
//	//return v + x.Int()
//}
//func (v Float) Sub(x Val) Val {
//	v -= x.Int()
//	//v.SelfSub(x)
//	return &v
//	//return v - x.Int()
//}
//func (v Float) Mul(x Val) Val {
//	v *= x.Int()
//	//v.SelfMul(x)
//	return &v
//	//return v * x.Int()
//}
//func (v Float) Div(x Val) Val {
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
func (v Float) Add(x Val) Val {
	v += x.Float()
	//	v.SelfAdd(x)
	return &v
	//return v + x.Int()
}
func (v Float) Sub(x Val) Val {
	v -= x.Float()
	//v.SelfSub(x)
	return &v
	//return v - x.Int()
}
func (v Float) Mul(x Val) Val {
	v *= x.Float()
	//v.SelfMul(x)
	return &v
	//return v * x.Int()
}

func (v Float) Div(x Val) Val {
	y := x.Float()
	if y == ZeroFloat {
		floatPanic.DevidedByZero()
	}
	v /= y
	//v.SelfDiv(x)
	return &v
	//return v / x.Int()
}

func (v Float) Rem(x Val) (dummy Val) {
	floatPanic.NotSupport("%")
	return
}

func (v Float) Shl(x Val) (dummy Val) {
	floatPanic.NotSupport("<<")
	return
}

func (v Float) Shr(x Val) (dummy Val) {
	floatPanic.NotSupport(">>")
	return
}

func (v Float) BitAnd(x Val) (dummy Val) {
	floatPanic.NotSupport("&")
	return
}

func (v Float) BitOr(x Val) (dummy Val) {
	floatPanic.NotSupport("|")
	return
}

func (v Float) BitXor(x Val) (dummy Val) {
	floatPanic.NotSupport("^")
	return
}

//func (v Float) AddInt(x Int) Int {
//	return v + x
//}

//func (v Float) SubInt(x Int) Int {
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
func (v Float) Equal(x Val) Val {
	y := Bool(v == x.Float()) && !x.IsNil()
	return Val(&y)
}

func (v Float) NotEqual(x Val) Val {
	y := Bool(v != x.Float()) || x.IsNil()
	return Val(&y)
}

func (v Float) Less(x Val) Val {
	y := Bool(v < x.Float())
	return Val(&y)
}

func (v Float) Greater(x Val) Val {
	y := Bool(v > x.Float())
	return Val(&y)
}

func (v Float) LessOrEqual(x Val) Val {
	y := Bool(v <= x.Float())
	return Val(&y)
}

func (v Float) GreaterOrEqual(x Val) Val {
	y := Bool(v >= x.Float())
	return Val(&y)
}

//	条件
func (v Float) And(x Val) Val {
	y := Bool(v.Bool() && x.Bool())
	return Val(&y)
}

func (v Float) Or(x Val) Val {
	y := Bool(v.Bool() || x.Bool())
	return Val(&y)
}
func (v Float) Xor(x Val) (dummy Val) {
	floatPanic.NotSupport("^")
	return
}

func (v Float) Not() Val {
	if v == 0 {
		v = Float(1)
	} else {
		v = Float(0)
	}
	return &v
}

func (v Float) BitNot() (dummy Val) {
	floatPanic.NotSupport("^")
	return
}

func (v Float) Minus() Val {
	v = -v
	return &v
}

//func (v Float) Inc() {
//	v++
//}

//func (v Float) Dec() {
//	v--
//}

func (v *Float) Inc() {
	*v++
}

func (v *Float) Dec() {
	*v--
}

//	チャネル
func (_ Float) In(_ Val) {
	floatPanic.NotSupport("in")
}

func (_ Float) Out() (y Val) {
	floatPanic.NotSupport("out")
	return
}

func (_ Float) Close() {
	floatPanic.NotSupport("close")
}

func (_ Float) Index(_ Int) (d Val) {
	floatPanic.NotSupport("index")
	return
}

func (_ Float) Map(_ Val) (d Val) {
	floatPanic.NotSupport("map")
	return
}

//	関数
func (_ Float) Call(_ ...Val) (y []Val) {
	floatPanic.NotSupport("call")
	return
}

//func (_ Float) Stmts() (y []ast.Stmt) {
//	panic("int cannot stmts")
//	return
//}

//	インターフェースで定義されたメソッドをポインターメソッドとして定義している場合は、
//	構造体のポインタを interface 化しなければならない
func (v Float) Copy() Val {
	//	x := Int(v)
	//	return Val(&x)
	return &v
}

//func (v Float) Equal(v2 Val) Bool {
//	return Bool(v == v2)
//}
