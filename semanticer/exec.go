package semanticer

import (
	//"bufio"
	//"flag"
	"fmt"
	//"io"
	"os"
	"os/exec"
	"strings"
)

//	パイプによりコマンド実行
//	[&&]分割
//	TODO	1行ごとに情報を初期化するか否か(cdなどについて)
func Exec(cmdStr string) {
	//debugFlag := func() bool {
	//	debugFlag := flag.Bool("d", false, "debug flag")
	//	flag.Parse()
	//	return *debugFlag
	//}()

	//reader := bufio.NewReader(os.Stdin)
	//for {
	//bytes, _, err := reader.ReadLine()
	//if err == io.EOF {
	//break
	//}
	for _, str := range strings.Split(cmdStr, "&&") {
		str = strings.TrimSpace(str)
		strs := strings.Split(str, " ")
		cmd := strs[0]
		var args []string
		if len(cmd) > 1 {
			args = strs[1:]
		}
		//if debugFlag {
		//	fmt.Println(str)
		//}
		execCmd := exec.Command(cmd, args...)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		if err := execCmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
	//}
}
