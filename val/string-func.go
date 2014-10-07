package val

import (
	"github.com/mitchellh/colorstring"
	"regexp"
	"strings"
)

//	正規表現を色付きに
func StringRegDraw(src string, color string, pattern string) string {
	reg := regexp.MustCompilePOSIX(pattern)
	return colorstring.Color(reg.ReplaceAllStringFunc(src, func(s string) string {
		//	間の色解除
		colorReg := regexp.MustCompilePOSIX("\x1b\\[[0-9]*;?[0-9]*;?[0-9]*;?m")
		s = colorReg.ReplaceAllString(s, "")
		return "[" + color + "]" + s + "[default]"
	}))
}

//	文字列を色付きに
func StringDraw(src string, color string, pattern string) string {
	return colorstring.Color(strings.Replace(src, pattern, "["+color+"]"+pattern+"[default]", -1))
}
