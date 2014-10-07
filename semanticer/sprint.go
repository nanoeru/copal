package semanticer

import (
	"fmt"
	"github.com/nanoeru/copal/val"
)

func Sprint(vs ...val.Val) string {
	str := ""
	for i := 0; i < len(vs); i++ {
		str += fmt.Sprint(vs[i].String())
		if i < len(vs)-1 {
			str += fmt.Sprint(" ")
		}
	}
	return str
}

func Sprintln(vs ...val.Val) string {
	return Sprint(vs...) + fmt.Sprintln()
}
