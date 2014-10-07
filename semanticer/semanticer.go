package semanticer

import (
	"bufio"
	"github.com/nanoeru/copal/ast"
	"github.com/nanoeru/copal/val"
	"github.com/nanoeru/copal/vm"
	"io"
	"os"
)

var (
	id string = "@"
)

type Semanticer struct {
	//	変数記憶用
	VarMap map[string]*val.Stack
	//	削除変数記憶用
	//	DelVarMap map[string]int
	//	引数・返り値記憶用
	//ValMap  map[int]val.Stack
	*val.Stack
	Namespace string
	//Counter   int
}

//	意味解析
func New() *Semanticer {
	id += "*"
	s := &Semanticer{
		VarMap:    make(map[string]*val.Stack),
		Stack:     val.NewStack(),
		Namespace: id,
	}
	return s
}

func (s Semanticer) MainAnalyze(stmts []ast.Stmt) *vm.FuncBlock {
	return s.parseMainStmts(stmts)
}

func (s Semanticer) Analyze(stmts []ast.Stmt) *vm.FuncBlock {
	return s.parseStmts(stmts)
}

func (s *Semanticer) Init() {
	//	標準入力チャネル
	stdinCh := val.NewChan(0)
	stack := s.GetVarStack("stdin")
	stack.Push(stdinCh)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			bytes, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			stdinCh.In(val.String(string(bytes)))
		}
		stdinCh.Close()
	}()
	//	TODO	初期化時に必要となるPushに対するPopの追加?!
}

//	初期設定変数
func (s *Semanticer) SetStringArgs(key string, args []string) {
	stack := s.GetVarStack(key)
	stack.Push(val.StringsToSlice(args...))
	//	TODO	初期化時に必要となるPushに対するPopの追加?!
}

func (s *Semanticer) GetVarStack(key string) *val.Stack {
	//stack, ok := s.VarMap[s.Namespace+key]
	//if !ok {
	//	stack, ok = s.VarMap[key]
	//	if !ok {
	//		stack = val.NewStack()
	//		s.VarMap[key] = stack
	//	} else {
	//		stack = val.NewStack()
	//		s.VarMap[key] = stack
	//	}
	//}
	stack, ok := s.VarMap[key]
	if !ok {
		stackLog.Info("	", s.Namespace, "create", key, "stack")
		stack = val.NewStack()
		s.VarMap[key] = stack
	}
	return stack
}

func (s *Semanticer) Copy() *Semanticer {
	id += ":"
	sc := Semanticer{
		VarMap:    make(map[string]*val.Stack),
		Stack:     val.NewStack(),
		Namespace: id,
	}
	//fmt.Println("s.VarMap", s.VarMap)
	//fmt.Println("sc.VarMap", sc.VarMap)
	////	実行時に行うべきこと?!
	//for k, v := range s.VarMap {
	//	stack := val.NewStack()
	//	stack.Push(v.Get())
	//	sc.VarMap[k] = stack
	//}
	//fmt.Println("s.VarMap", s.VarMap)
	//fmt.Println("sc.VarMap", sc.VarMap)
	return &sc
}

////	代入のみ
//func (s *Semanticer) SetVar(key string, x val.Val) {
//	stack, ok := s.VarMap[key]
//	if !ok {
//		panic(fmt.Sprintln("var-stack-map set var error not exist [", key, "]"))
//	}
//	//fmt.Println("var-stack-map set", key, x.String())
//	stack.Set(x)
//}

////	変数マップスタックへプッシュ
//func (s *Semanticer) PushVar(key string, x val.Val) {
//	stack, ok := s.VarMap[key]
//	if !ok {
//		//fmt.Println("var-stack-map create new stack", key)
//		stack = val.NewStack()
//		s.VarMap[key] = stack
//	}
//	//fmt.Println("var-stack-map push", key, x.String())
//	stack.Push(x)
//}

////	取得のみ
//func (s *Semanticer) GetVar(key string) val.Val {
//	stack, ok := s.VarMap[key]
//	if !ok {
//		panic(fmt.Sprintln("var-stack-map get var error not exist [", key, "]"))
//	}
//	return stack.Get()
//}

//func (s *Semanticer) PopVar(key string) val.Val {
//	stack, ok := s.VarMap[key]
//	if !ok {
//		panic(fmt.Sprintln("var-stack-map pop var error not exist [", key, "]"))
//	}
//	return stack.Pop()
//}

//func (s *Semanticer) StackVal() {
//	s.PushVal(nil)
//}

//func (s *Semanticer) PushVal(index int, x val.Val) {
//	stack, ok := s.ValMap[index]
//	if !ok {
//		stack = val.NewStack()
//		s.ValMap[index] = stack
//	}
//	stack.Push(x)
//}

//func (s *Semanticer) PopVal(index int) val.Val {
//	stack, ok := s.ValMap[index]
//	if !ok {
//		panic("PopVal Error")
//	}
//	return stack.Pop()
//}
