package semanticer

import (
	"fmt"
	"github.com/nanoeru/copal/val"
	"github.com/nanoeru/fondot"
	"log"
	"time"
)

//	ArgN 0	引数0個
//	ArgN 1	引数1個
//	ArgN -1	引数1個目から可変長引数
type FuncStruct struct {
	ArgN int
	F    func(...val.Val) []val.Val
}

func NewFuncStrucy(argn int, f func(...val.Val) []val.Val) FuncStruct {
	return FuncStruct{
		ArgN: argn,
		F:    f,
	}
}

var dummy = NewFuncStrucy(0,
	func(vs ...val.Val) []val.Val {
		return nil
	},
)

var builtinFuncMap = map[string]FuncStruct{
	"append": NewFuncStrucy(-2,
		func(vs ...val.Val) []val.Val {
			appendSlices := vs[1:]
			v := append(vs[0].Slice(), appendSlices...)
			return []val.Val{&v}
		},
	),
	"bool": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].Bool()
			return []val.Val{&v}
		},
	),
	"cap": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].Cap()
			return []val.Val{&v}
		},
	),
	//"camera": NewFuncStrucy(0,
	//	func(vs ...val.Val) []val.Val {
	//		//Camera()
	//		return nil
	//	},
	//),
	"chan": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].Chan()
			return []val.Val{&v}
		},
	),
	"new_chan": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			vals := make([]val.Val, len(vs))
			for i := 0; i < len(vs); i++ {
				vals[i] = val.NewChan(int(vs[0].Int()))
			}
			return vals
		},
	),
	"clip_read": NewFuncStrucy(0,
		func(vs ...val.Val) []val.Val {
			v := val.String(ClipBoardRead())
			return []val.Val{&v}
		},
	),
	"clip_write": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			ClipBoardWrite(string(vs[0].String()))
			return nil
		},
	),
	"close": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			for i := 0; i < len(vs); i++ {
				vs[i].Close()
			}
			return nil
		},
	),
	"csv_read": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			return []val.Val{CSVRead(string(vs[0].String()))}
		},
	),
	//"defer": func(vs ...val.Val) []val.Val {
	//	return []val.Val{vs[0].Close()}
	//},
	//"error": func(vs ...val.Val) []val.Val {
	//	return []val.Val{vs[0].Error()}
	//},
	"exec": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].String()
			Exec(string(v))
			return nil
		},
	),
	"exit": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			Exit(int(vs[0].Int()))
			return nil
		},
	),
	//"f": NewFuncStrucy(1,
	//	func(vs ...val.Val) []val.Val {
	//		v := vs[0].Func()
	//		F(v)
	//		//			return []val.Val{&v}
	//		return nil
	//	},
	//),
	"fread": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			return []val.Val{FRead(string(vs[0].String()))}
		},
	),
	"fwrite": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			return []val.Val{FWrite(string(vs[0].String()))}
		},
	),
	"server": NewFuncStrucy(2,
		func(vs ...val.Val) []val.Val {
			v := val.String(Server(string(vs[0].String()), vs[1].Func()).Error())
			return []val.Val{&v}
		},
	),
	"http_server": NewFuncStrucy(2,
		func(vs ...val.Val) []val.Val {
			v := val.String(HttpServer(string(vs[0].String()), vs[1].Func()).Error())
			return []val.Val{&v}
		},
	),
	"file_server": NewFuncStrucy(2,
		func(vs ...val.Val) []val.Val {
			v := val.String(FileServer(string(vs[0].String()), string(vs[1].String())).Error())
			return []val.Val{&v}
		},
	),
	//"fread": func(vs ...val.Val) []val.Val {
	//	return []val.Val{vs[0].Float()}
	//},
	"float": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].Float()
			return []val.Val{&v}
		},
	),
	//"import": func(vs ...val.Val) []val.Val {
	//	return []val.Val{vs[0].Len()}
	//},
	"int": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].Int()
			return []val.Val{&v}
		},
	),
	//"interface": func(vs ...val.Val) []val.Val {
	//	return []val.Val{vs[0].Interface()}
	//},
	"len": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].Len()
			return []val.Val{&v}
		},
	),
	"log": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			log.Print(Sprintln(vs...))
			return nil
		},
	),
	//"new": func(vs ...val.Val) []val.Val {
	//	return []val.Val{vs[0].Len()}
	//},
	"dot_print": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			fondot.Draw(Sprint(vs...), fondot.Grad)
			return nil
		},
	),
	"dot_println": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			fondot.Draw(Sprint(vs...), fondot.GradMono)
			fmt.Println()
			return nil
		},
	),
	"print": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			fmt.Print(Sprint(vs...))
			return nil
		},
	),
	"println": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			fmt.Print(Sprintln(vs...))
			return nil
		},
	),
	"info": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			Info(Sprint(vs...))
			return nil
		},
	),
	"warn": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			Warn(Sprint(vs...))
			return nil
		},
	),
	"error": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			Error(Sprint(vs...))
			return nil
		},
	),
	"fatal": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			Fatal(Sprint(vs...))
			return nil
		},
	),
	"panic": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			panic(vs[0].String())
			return nil
		},
	),
	"string": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].String()
			return []val.Val{&v}
		},
	),
	"slice": NewFuncStrucy(-1,
		func(vs ...val.Val) []val.Val {
			v := val.NewSlice(vs...)
			return []val.Val{&v}
		},
	),
	"sleep": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			v := vs[0].Int()
			time.Sleep(time.Duration(v) * time.Millisecond)
			return []val.Val{&v}
		},
	),
	"led7": NewFuncStrucy(7,
		func(vs ...val.Val) []val.Val {
			NumberLamp(vs)
			return nil
		},
	),
	"cir_nand": NewFuncStrucy(2,
		func(vs ...val.Val) []val.Val {
			chAns := val.NewChan(0)
			//	同一入力
			if vs[0].Equal(vs[1]).Bool() {
				go func() {
					for {
						//	1行記述の場合は1個目の条件を満たすと後の条件を判断しないので注意
						//chAns.In(!(vs[0].Chan().Out().Bool() && vs[1].Chan().Out().Bool()))
						//fmt.Println("chan-wait")
						a := vs[0].Out()
						if a.IsNil() {
							return
						}
						//fmt.Println("chan-wait:a", a.String())
						//fmt.Println("chan-wait:b(a)", a.String())
						c := val.Bool(!(a.Bool() && a.Bool()))
						//fmt.Println("chan-wait:c", c.String())
						chAns.In(c)
						//chAns <- c
						//chAns.In(!a.Bool() && b.Bool())
						//chAns.In(val.Bool(!(vs[0].Chan().Out().Bool() && vs[1].Chan().Out().Bool())))
						//chAns.In(!(vs[0].Chan().Out().Bool() && vs[1].Chan().Out().Bool()))
						//fmt.Println("chan-done")
					}
				}()
			} else {
				//	TODO	無名関数なのでスコープに注意
				go func(chA, chB val.Val) {
					for {
						//	1行記述の場合は1個目の条件を満たすと後の条件を判断しないので注意
						//chAns.In(!(vs[0].Chan().Out().Bool() && vs[1].Chan().Out().Bool()))
						//fmt.Println("chan-wait")
						a := chA.Out()
						//fmt.Println("chan-wait:a", a.String())
						b := chB.Out()
						if a.IsNil() || b.IsNil() {
							chAns.In(val.NewNil())
							return
						}
						//fmt.Println("chan-wait:b", b.String())
						c := val.Bool(!(a.Bool() && b.Bool()))
						//fmt.Println("chan-wait:c", c.String())
						chAns.In(c)
						//chAns <- c
						//chAns.In(!a.Bool() && b.Bool())
						//chAns.In(val.Bool(!(vs[0].Chan().Out().Bool() && vs[1].Chan().Out().Bool())))
						//chAns.In(!(vs[0].Chan().Out().Bool() && vs[1].Chan().Out().Bool()))
						//fmt.Println("chan-done")
					}
				}(vs[0], vs[1])
			}
			return []val.Val{chAns}
		},
	),
	"cir_and": NewFuncStrucy(-2,
		func(vs ...val.Val) []val.Val {
			chAns := val.NewChan(0)
			go func(chs []val.Val) {
				for {
					//fmt.Println("chan-and-wait")
					result := val.Bool(true)
					for i := 0; i < len(chs); i++ {
						out := chs[i].Out()
						if out.IsNil() {
							chAns.In(out)
							return
						}
						//fmt.Println("chan-and-out")
						result = result && out.Bool()
					}
					chAns.In(val.Bool(result))
					//fmt.Println("chan-and-done")
				}
			}(vs)
			return []val.Val{chAns}
		},
	),
	"cir_or": NewFuncStrucy(-2,
		func(vs ...val.Val) []val.Val {
			chAns := val.NewChan(0)
			go func(chs []val.Val) {
				for {
					//fmt.Println("chan-or-wait")
					result := val.Bool(false)
					for i := 0; i < len(chs); i++ {
						out := chs[i].Out()
						if out.IsNil() {
							chAns.In(out)
							return
						}
						//fmt.Println("chan-or-out")
						result = result || out.Bool()
					}
					chAns.In(val.Bool(result))
					//fmt.Println("chan-or-done")
				}
			}(vs)
			return []val.Val{chAns}
		},
	),
	"cir_not": NewFuncStrucy(1,
		func(vs ...val.Val) []val.Val {
			chAns := val.NewChan(0)
			go func(ch val.Val) {
				for {
					//fmt.Println("chan-not-wait")
					out := ch.Out()
					if out.IsNil() {
						chAns.In(out)
						return
					}
					chAns.In(!out.Bool())
					//fmt.Println("chan-not-done")
				}
			}(vs[0])
			return []val.Val{chAns}
		},
	),
	"cir_split": NewFuncStrucy(2,
		func(vs ...val.Val) []val.Val {
			num := int(vs[0].Int())
			ret := make([]val.Val, num)
			for i := 0; i < num; i++ {
				ret[i] = val.NewChan(0)
			}
			go func(ch val.Val) {
				for {
					//fmt.Println("chan-split-wait")
					out := ch.Out()
					if out.IsNil() {
						for i := 0; i < num; i++ {
							ret[i].In(out)
						}
						return
					}
					for i := 0; i < num; i++ {
						ret[i].In(out.Bool())
					}
					//fmt.Println("chan-split-done")
				}
			}(vs[1])
			return ret
		},
	),
}
