package val

import (
	"github.com/nanoeru/copal/val/vpanic"
)

//	利点：初期化せずに使える
//		データの上書きによる損失が存在しない
//	欠点：スライス型より遅い
type Stack struct {
	top  *Element
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

type Element struct {
	value Val
	next  *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s Stack) Copy() *Stack {
	return &s
}

func (s *Stack) Push(value Val) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value Val) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Set(value Val) {
	if s.size > 0 {
		s.top.value = value
		return
	}
	vpanic.SetNullStack()
}

func (s *Stack) Get() (value Val) {
	if s.size > 0 {
		value = s.top.value
		return
	}
	vpanic.GetNullStack()
	return nil
}
