package semanticer

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/nanoeru/copal/val"
	"io"
	"os"
)

//	ファイルチャネル読み込み
func ChanReadLine(ch *val.Chan, file io.Reader) {
	func() {
		reader := bufio.NewReader(file)
		for {
			bytes, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			ch.In(val.String(string(bytes)))
		}
		ch.Close()
	}()
}

//	ファイルチャネル書き込み
func ChanWriteLine(ch *val.Chan, file io.Writer) {
	func() {
		writer := bufio.NewWriter(file)
		for {
			str := ch.Out()
			//	チャネルが閉じられた
			if str.IsNil() {
				break
			}

			_, err := writer.WriteString(string(str.String()))
			if err != nil {
				ch.Close()
				break
			}
			err = writer.Flush()
			if err != nil {
				ch.Close()
				break
			}
		}
	}()
}

//	CSVファイルチャネル読み込み
func ChanCSVReadLine(ch *val.Chan, file io.Reader) {
	func() {
		reader := csv.NewReader(file)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			ch.In(val.StringsToSlice(record...))
		}
		ch.Close()
	}()
}

//	ファイル読み込み
func FRead(fp string) *val.Chan {
	ch := val.NewChan(0)
	f, err := os.Open(fp)
	if err != nil {
		fmt.Println(err)
		ch.Close()
		return ch
	}
	go ChanReadLine(ch, f)
	return ch
}

//	ファイル書き込み
func FWrite(fp string) *val.Chan {
	ch := val.NewChan(0)
	f, err := os.Create(fp)
	if err != nil {
		fmt.Println(err)
		ch.Close()
		return ch
	}
	//	内容を初期化
	err = f.Truncate(0)
	if err != nil {
		fmt.Println(err)
		ch.Close()
		return ch
	}
	go ChanWriteLine(ch, f)
	return ch
}

//	CSVファイル書き込み
func CSVRead(fp string) *val.Chan {
	ch := val.NewChan(0)
	f, err := os.Open(fp)
	if err != nil {
		fmt.Println(err)
		ch.Close()
		return ch
	}
	go ChanCSVReadLine(ch, f)
	return ch
}
