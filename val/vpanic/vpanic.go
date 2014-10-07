package vpanic

import (
	"fmt"
)

var Name = "Copal: "

type ValPanic struct {
	Name   string
	Format string
}

func NewValPanic(name string) *ValPanic {
	return &ValPanic{
		Name:   name,
		Format: "%s : %s",
	}
}

func (v *ValPanic) SetFormat(format string) *ValPanic {
	v.Format = format
	return v
}

func (v *ValPanic) Panic(a ...interface{}) {
	a = append([]interface{}{v.Name}, a...)
	panic(Name + fmt.Sprintf(v.Format, a...))
}

func (v *ValPanic) NotSupport(name string) {
	v.SetFormat("%s : not support %s").Panic(name)
}

func (v *ValPanic) DevidedByZero() {
	v.SetFormat("%s : devided by 0").Panic()
}

func GetNullStack() {
	NewValPanic("get null stack").Panic("")
}

func SetNullStack() {
	NewValPanic("set null stack").Panic("")
}

func Do(name, str string) {
	NewValPanic(name).Panic(str)
}
