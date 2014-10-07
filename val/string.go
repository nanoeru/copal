package val

import (
	"fmt"
	"strconv"
	"strings"
)

func NewString(v string) Val {
	tmp := String(v)
	return &tmp
}

func (v String) Int() Int {
	//	base -> 0 -> 基数自動認識
	i, err := strconv.ParseInt(string(v), 0, 64)
	if err != nil {
		stringPanic.Panic(err)
	}
	return Int(i)
}

func (v String) Float() Float {
	f, err := strconv.ParseFloat(string(v), 64)
	if err != nil {
		stringPanic.Panic(err)
	}
	return Float(f)
}

func (v String) String() String {
	return v
}

func (_ String) Type() String {
	return String("string")
}

func (v String) Bool() Bool {
	return v != ""
}

func (v String) Chan() (y Chan) {
	stringPanic.NotSupport("chan")
	return
}

func (v String) Nil() (y Nil) {
	stringPanic.NotSupport("nil")
	return
}

func (v String) Func() (y Func) {
	stringPanic.NotSupport("func")
	return
}

func (v String) Slice() (y Slice) {
	stringPanic.NotSupport("slice")
	return
}

func (v String) IsNil() Bool {
	return false
}

func (v String) Len() Int {
	return Int(len(v))
}

func (_ String) Cap() (dummy Int) {
	stringPanic.NotSupport("cap")
	return
}

//func (v *String) Add(x Val) {
//	*v += x.String()
//}

//func (v *String) SelfAdd(x Val) {
//	*v += x.String()
//}
//func (v *String) SelfSub(_ Val) {
//	stringPanic.NotSupport("subed")
//}
//func (v *String) SelfMul(_ Val) {
//	stringPanic.NotSupport("muled")
//}
//func (v *String) SelfDiv(_ Val) {
//	stringPanic.NotSupport("dived")
//}

func (v String) Add(x Val) Val {
	return v + x.String()
	//v.SelfAdd(x)
	//return &v
}
func (_ String) Sub(_ Val) Val {
	stringPanic.NotSupport("-")
	return nil
}
func (_ String) Mul(_ Val) Val {
	stringPanic.NotSupport("*")
	return nil
}
func (_ String) Div(_ Val) Val {
	stringPanic.NotSupport("/")
	return nil
}

func (_ String) Rem(_ Val) Val {
	stringPanic.NotSupport("%")
	return nil
}

func (_ String) Shl(_ Val) Val {
	stringPanic.NotSupport("<<")
	return nil
}

func (_ String) Shr(_ Val) Val {
	stringPanic.NotSupport(">>")
	return nil
}

//func (_ String) AddInt(_ Int) (y Int) {
//	panic("bool cannot addint")
//	return
//}

//func (_ String) SubInt(_ Int) (y Int) {
//	panic("bool cannot addint")
//	return
//}

//func (v *String) AddInt(x Int) {
//	*v += x.String()
//}

//func (v *String) SubInt(x Int) {
//	stringPanic.NotSupport("subed int")
//}

//func (v *String) MulInt(x Int) {
//	stringPanic.NotSupport("muled int")
//}

//func (v *String) DivInt(x Int) {
//	stringPanic.NotSupport("dived int")
//}

//func (v *String) AddString(x String) {
//	*v += x
//}

//func (v *String) SubString(x String) {
//	stringPanic.NotSupport("subed String")
//}

//func (v *String) MulString(x String) {
//	stringPanic.NotSupport("muled String")
//}

//func (v *String) DivString(x String) {
//	stringPanic.NotSupport("dived String")
//}

func (v String) Equal(x Val) Val {
	result := v == x.String()
	y := Bool(result) && !x.IsNil()
	return Val(&y)
}

func (v String) NotEqual(x Val) Val {
	y := Bool(v != x.String()) || x.IsNil()
	return Val(&y)
}

func (v String) Less(x Val) Val {
	y := Bool(v < x.String())
	return Val(&y)
}

func (v String) Greater(x Val) Val {
	y := Bool(v > x.String())
	return Val(&y)
}

func (v String) LessOrEqual(x Val) Val {
	y := Bool(v <= x.String())
	return Val(&y)
}

func (v String) GreaterOrEqual(x Val) Val {
	y := Bool(v >= x.String())
	return Val(&y)
}

func (_ String) And(_ Val) (y Val) {
	stringPanic.NotSupport("&&")
	return
}

func (_ String) Or(_ Val) (y Val) {
	stringPanic.NotSupport("||")
	return
}
func (_ String) Xor(_ Val) (y Val) {
	stringPanic.NotSupport("^^")
	return
}
func (_ String) BitAnd(_ Val) (y Val) {
	stringPanic.NotSupport("&")
	return
}

func (_ String) BitOr(_ Val) (y Val) {
	stringPanic.NotSupport("|")
	return
}
func (_ String) BitXor(_ Val) (y Val) {
	stringPanic.NotSupport("^")
	return
}
func (_ String) Not() (y Val) {
	stringPanic.NotSupport("!")
	return
}
func (_ String) BitNot() (y Val) {
	stringPanic.NotSupport("~")
	return
}

func (_ String) Minus() (y Val) {
	stringPanic.NotSupport("-")
	return
}

func (v String) Copy() Val {
	x := String(v)
	return Val(&x)
}

func (_ String) Inc() {
	stringPanic.NotSupport("++")
	return
}

func (_ String) Dec() {
	stringPanic.NotSupport("--")
	return
}

func (_ String) In(_ Val) {
	stringPanic.NotSupport("in")
}

func (_ String) Out() (y Val) {
	stringPanic.NotSupport("out")
	return
}

func (_ String) Close() {
	stringPanic.NotSupport("close")
}

func (_ String) Index(x Int) (y Val) {
	stringPanic.NotSupport("index")
	return
}

func StringsToSlice(a ...string) Slice {
	tmp := make(Slice, len(a))
	for i, v := range a {
		tmp[i] = String(v)
	}
	return tmp
}

func SliceToStrings(a Slice) []string {
	tmp := make([]string, len(a))
	for i, v := range a {
		tmp[i] = string(v.String())
	}
	return tmp
}

func (v String) Map(key Val) (ret Val) {
	switch key.String() {
	case "has_prefix":
		ret = NewFunc(func(vs ...Val) []Val {
			flag := strings.HasPrefix(string(v.String()), string(vs[0].String()))
			return []Val{Bool(flag)}
		})
	case "has_suffix":
		ret = NewFunc(func(vs ...Val) []Val {
			flag := strings.HasSuffix(string(v.String()), string(vs[0].String()))
			return []Val{Bool(flag)}
		})
	case "split":
		ret = NewFunc(func(vs ...Val) []Val {
			strs := strings.Split(string(v.String()), string(vs[0].String()))
			return []Val{StringsToSlice(strs...)}
		})
	case "trim_left":
		ret = NewFunc(func(vs ...Val) []Val {
			//	重要：変数に格納した後にアドレス…
			str := String(strings.TrimLeft(string(v.String()), string(vs[0].String())))
			return []Val{&str}
		})
	case "draw":
		ret = NewFunc(func(vs ...Val) []Val {
			str := StringDraw(string(v), string(vs[0].String()), string(vs[1].String()))
			return []Val{String(str)}
		})
	case "reg_draw":
		ret = NewFunc(func(vs ...Val) []Val {
			str := StringRegDraw(string(v), string(vs[0].String()), string(vs[1].String()))
			return []Val{String(str)}
		})
	default:
		panic(fmt.Sprintln("String cannnot have map", "[", key.String(), "]"))
	}
	return
}

func (_ String) Call(_ ...Val) (y []Val) {
	stringPanic.NotSupport("call")
	return
}

//func (_ String) Stmts() (y []ast.Stmt) {
//	stringPanic.NotSupport("stmts")
//	return
//}

//func (v String) Val() Val {
//	if v != ZeroString {
//		return OneBool
//	} else {
//		return ZeroBool
//	}
//}

//func (v *String) Set(v2 CanString) {
//	*v = v2.String()
//}

//func (v *String) Add(v2 CanString) {
//	*v += v2.String()
//}

//func (v *String) Sub(v2 CanString) {
//	stringPanic.NotSupport("Sub")
//}

//func (v *String) Mul(v2 CanString) {
//	stringPanic.NotSupport("Mul")
//}

//func (v *String) Div(v2 CanString) {
//	stringPanic.NotSupport("Div")
//}

//func (v String) Equal(v2 Val) Bool {
//	return Bool(v == v2)
//}
