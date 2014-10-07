package semanticer

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/nanoeru/copal/val"
)

/*
	 -
	| |
	 -
	| |
	 -

	 0
	5 1
	 6
	4 2
	 3
*/
func NumberLamp(vs []val.Val) {
	func(vs []val.Val) {
		draw := func(color, str string) {
			fmt.Print(colorstring.Color("[_" + color + "_]" + str + "[default]"))
		}
		for {
			off := "black"
			on := "blue"
			ons := make([]string, 7)
			for i := 0; i < 7; i++ {
				out := vs[i].Out()
				if out.IsNil() {
					return
				}
				if out.Bool() {
					ons[i] = on
				} else {
					ons[i] = off
				}
			}
			//	0
			draw(off, "  ")
			draw(ons[0], "    ")
			draw(off, "  ")
			draw(off, "\n")
			//	5 1
			draw(ons[5], "  ")
			draw(off, "    ")
			draw(ons[1], "  ")
			draw(off, "\n")
			draw(ons[5], "  ")
			draw(off, "    ")
			draw(ons[1], "  ")
			draw(off, "\n")
			//	6
			draw(off, "  ")
			draw(ons[6], "    ")
			draw(off, "  ")
			draw(off, "\n")
			//	4 2
			draw(ons[4], "  ")
			draw(off, "    ")
			draw(ons[2], "  ")
			draw(off, "\n")
			draw(ons[4], "  ")
			draw(off, "    ")
			draw(ons[2], "  ")
			draw(off, "\n")
			//	3
			draw(off, "  ")
			draw(ons[3], "    ")
			draw(off, "  ")
			draw(off, "\n")
			//	new line
			draw(off, "\n")
		}
	}(vs)
}
