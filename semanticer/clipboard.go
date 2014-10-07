package semanticer

import (
	"github.com/atotto/clipboard"
)

func ClipBoardWrite(text string) {
	_ = clipboard.WriteAll(text)
}

func ClipBoardRead() (text string) {
	text, _ = clipboard.ReadAll()
	return
}
