package main

import (
	"fmt"
	"github.com/nanoeru/copal/semanticer"
	"github.com/nanoeru/copal/shell/parser"
	"github.com/nanoeru/copal/shell/scanner"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

import (
	"github.com/nanoeru/nlogger"
)

var mainLog = nlogger.NewLogger()

func init() {
	//	並列化設定
	runtime.GOMAXPROCS(runtime.NumCPU())
	mainLog.Info.Idle()
}

func main() {
	scanner := scanner.New()
	if len(os.Args) >= 2 {
		bytes, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		scanner.Init(string(bytes))
	} else {
		scanner.Init(`println("Usage: 1st arg is file-path")`)
	}

	mainLog.Info("[Lexer-start]")
	stmts, err := parser.ParseScanner(scanner)
	mainLog.Info("[Lexer-end]")
	if err != nil {
		mainLog.Error("[Parse-err]")
		mainLog.Error(err)
	}

	t1 := time.Now().UnixNano()
	s := semanticer.New()
	s.Init()
	args := os.Args
	if len(os.Args) >= 2 {
		args = args[2:]
	} else {
		args = args[1:]
	}
	s.SetStringArgs("$", args)
	{
		defer func() {
			if e := recover(); e != nil {
				if strings.HasPrefix(fmt.Sprint(e), "Copal:") {
					//	インタプリタ上のpanic
					mainLog.Error(e)
				} else {
					//	予期せぬpanic
					//	適当数容量確保
					buf := make([]byte, 10000)
					runtime.Stack(buf, false)
					mainLog.Fatal("Stack Trace :", e, string(buf))
				}
			}
		}()
		s.MainAnalyze(stmts).Do()
	}
	t2 := time.Now().UnixNano()
	sub := time.Duration(t2 - t1)
	mainLog.Info("time:", float64(sub)/float64(time.Second))
}
