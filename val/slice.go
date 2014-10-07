package val

import (
	"reflect"
	"strings"
)

func NewSliceN(n int, vals ...Val) Slice {
	slice := make([]Val, n)
	return append(slice, vals...)
}

func NewSlice(vals ...Val) Slice {
	return vals
}

func (_ Slice) Int() (d Int) {
	slicePanic.NotSupport("int")
	return
}

func (_ Slice) Float() (d Float) {
	slicePanic.NotSupport("float")
	return
}

func (v Slice) String() String {
	str := String("[")
	for i, v := range v {
		if i != 0 {
			str += ", "
		}
		str += v.String()
	}
	str += "]"
	return str
}

func (_ Slice) Type() String {
	return String("slice")
}

func (v Slice) Bool() Bool {
	return v != nil
}

func (_ Slice) Chan() (d Chan) {
	slicePanic.NotSupport("chan")
	return
}

func (_ Slice) Nil() (d Nil) {
	slicePanic.NotSupport("nil")
	return
}

func (_ Slice) Func() (y Func) {
	slicePanic.NotSupport("func")
	return
}

func (v Slice) Slice() Slice {
	return v
}

func (v Slice) IsNil() Bool {
	return false
}

func (v Slice) Len() (x Int) {
	return Int(len(v))
}

func (v Slice) Cap() Int {
	return Int(cap(v))
}

func (_ Slice) Add(_ Val) (d Val) {
	slicePanic.NotSupport("+")
	return
}

func (_ Slice) Sub(_ Val) (d Val) {
	slicePanic.NotSupport("-")
	return
}

func (_ Slice) Mul(_ Val) (d Val) {
	slicePanic.NotSupport("*")
	return
}

func (_ Slice) Div(_ Val) (d Val) {
	slicePanic.NotSupport("/")
	return
}

func (_ Slice) Rem(_ Val) (d Val) {
	slicePanic.NotSupport("%")
	return
}

func (_ Slice) Shl(_ Val) (d Val) {
	slicePanic.NotSupport("<<")
	return
}

func (_ Slice) Shr(_ Val) (d Val) {
	slicePanic.NotSupport(">>")
	return
}

//func (v Slice) AddInt(x Int) (d Int) {
//	slicePanic.NotSupport("addint")
//	return
//}

//func (v Slice) SubInt(x Int) (d Int) {
//	slicePanic.NotSupport("subint")
//	return
//}

func (v Slice) Equal(x Val) Val {
	y := Bool(reflect.DeepEqual(v, x.Slice())) && !x.IsNil()
	return Val(&y)
}

func (v Slice) NotEqual(x Val) Val {
	y := Bool(!reflect.DeepEqual(v, x.Slice())) || x.IsNil()
	return Val(&y)
}

func (_ Slice) Less(_ Val) (d Val) {
	slicePanic.NotSupport("<")
	return
}

func (_ Slice) Greater(_ Val) (d Val) {
	slicePanic.NotSupport(">")
	return
}

func (_ Slice) LessOrEqual(_ Val) (d Val) {
	slicePanic.NotSupport("<=")
	return
}

func (_ Slice) GreaterOrEqual(_ Val) (d Val) {
	slicePanic.NotSupport(">=")
	return
}

func (_ Slice) And(_ Val) (d Val) {
	slicePanic.NotSupport("&&")
	return
}

func (_ Slice) Or(_ Val) (d Val) {
	slicePanic.NotSupport("||")
	return
}

func (_ Slice) Xor(_ Val) (d Val) {
	slicePanic.NotSupport("^^")
	return
}
func (_ Slice) BitAnd(_ Val) (d Val) {
	slicePanic.NotSupport("&")
	return
}

func (_ Slice) BitOr(_ Val) (d Val) {
	slicePanic.NotSupport("|")
	return
}

func (_ Slice) BitXor(_ Val) (d Val) {
	slicePanic.NotSupport("^")
	return
}

func (_ Slice) Not() (d Val) {
	slicePanic.NotSupport("!")
	return
}
func (_ Slice) BitNot() (d Val) {
	slicePanic.NotSupport("~")
	return
}

func (_ Slice) Minus() (d Val) {
	slicePanic.NotSupport("-")
	return
}

func (_ Slice) Inc() {
	slicePanic.NotSupport("++")
	return
}

func (v Slice) Dec() {
	slicePanic.NotSupport("--")
	return
}

func (_ Slice) In(_ Val) {
	slicePanic.NotSupport("in")
}

func (_ Slice) Out() (y Val) {
	slicePanic.NotSupport("out")
	return
}

func (_ Slice) Close() {
	slicePanic.NotSupport("close")
}

func (v Slice) Index(x Int) (y Val) {
	return v[x]
}

func (v Slice) Map(key Val) (ret Val) {
	switch key.String() {
	case "join":
		ret = NewFunc(func(vs ...Val) []Val {
			str := strings.Join(SliceToStrings(v), string(vs[0].String()))
			return []Val{String(str)}
		})
	case "filter":
		ret = NewFunc(func(vs ...Val) []Val {
			slice := Slice(make([]Val, 0))
			for _, e := range v {
				if vs[0].Func().Call(e)[0].Bool() {
					slice = append(slice, e)
				} else {
				}
			}
			return []Val{slice}
		})
	case "unfilter":
		ret = NewFunc(func(vs ...Val) []Val {
			slice := Slice(make([]Val, 0))
			for _, e := range v {
				if vs[0].Func().Call(e)[0].Bool() {
				} else {
					slice = append(slice, e)
				}
			}
			return []Val{slice}
		})
	default:
		return v[key.Int()]
	}
	return
}

func (_ Slice) Call(_ ...Val) (y []Val) {
	slicePanic.NotSupport("call")
	return
}

//	インターフェースで定義されたメソッドをポインターメソッドとして定義している場合は、
//	構造体のポインタを interface 化しなければならない
func (v Slice) Copy() Val {
	x := Slice(v)
	return Val(&x)
}

//	中身を含めてのコピー
//func (v Slice) DeepCopy() Val {
//	x := Slice(v)
//	return Val(&x)
//}

//func (v Slice) Equal(v2 Val) Bool {
//	return Bool(v == v2)
//}
