package semanticer

import (
	"fmt"
	"github.com/nanoeru/copal/val"
	"net"
	"net/http"
	"os"
	"strings"
)

//	文字列(src)の先頭に文字列(prefix)付加
func SetPrefix(src, prefix string) string {
	if strings.HasPrefix(src, prefix) {
		return src
	}
	return prefix + src
}

//	静的ファイルサーバー
func FileServer(hp, root string) error {
	//	ディレクトリ確認
	info, err := os.Stat(root)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return err
	}
	fmt.Println("Http Static File Server")
	fmt.Println("Please access", hp)

	return http.ListenAndServe(SetPrefix(hp, ":"), http.FileServer(http.Dir(root)))
	//; err != nil {
	//	log.Printf("Server Error: %v", err)
	//}
}

func F(f val.Func) {
	v := val.String("cat")
	f.Call(v)
}

//func (h Hello) ServeHTTP(
//	w http.ResponseWriter,
//	r *http.Request) {
//	fmt.Fprint(w, "Hello!")
//}

type HttpServerStruct struct {
	val.Func
}

func (v *HttpServerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(v.Func.Call(val.String(r.URL.Path))[0].String()))
}

//	httpサーバー
func HttpServer(port string, f val.Func) error {
	fmt.Println("Server is running!\nAccess to [http://localhost:" + port + "]")
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, string(f.Call(val.String(r.URL.Path))[0].String()))
	//})
	return http.ListenAndServe(SetPrefix(port, ":"), &HttpServerStruct{f})
}

//	そのままのサーバー
func Server(port string, f val.Func) error {
	listener, err := net.Listen("tcp", SetPrefix(port, ":"))
	if err != nil {
		return err
	}
	fmt.Println("[Listening]", port)

	//var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept:", err)
			return err
		}
		fmt.Println("[Accepted]")
		//wg.Add(1)
		go func(conn net.Conn) {
			defer func() {
				conn.Close()
				//wg.Done()
			}()
			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("Error reading:", err)
					return
				}
				fmt.Println("received ", n, " bytes of data")
				fmt.Println(string(buf))

				//f.Call()
				//f.RetCall()
				//ret = NewFunc(func(vs ...Val) []Val {
				//	slice := Slice(make([]Val, 0))
				//	for _, e := range v {

				//		if vs[0].Func().Call(e)[0].Bool() {
				//			slice = append(slice, e)
				//		} else {
				//		}
				//	}
				//	return []Val{slice}
				//})
				res := []byte(string(f.Call(val.String(string(buf)))[0].String()))

				//	echo
				_, err = conn.Write(res)
				if err != nil {
					fmt.Println("Error send echo:", err)
					return
				} else {
					fmt.Println("[Sended echo]")
				}
			}
		}(conn)
	}
}
