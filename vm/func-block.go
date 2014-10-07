package vm

import ()

type (
	FuncBlock struct {
		Funcs      *NoNameFuncSlice
		DeferFuncs *NoNameFuncSlice
	}
)

func NewFuncBlock() *FuncBlock {
	return &FuncBlock{
		Funcs:      new(NoNameFuncSlice),
		DeferFuncs: new(NoNameFuncSlice),
	}
}

func (v *FuncBlock) CleanDeferFunc() {
	v.DeferFuncs = new(NoNameFuncSlice)
}

func (v *FuncBlock) IsBlank() bool {
	return v.Funcs.IsBlank() && v.Funcs.IsBlank()
}

func (v *FuncBlock) Do() {
	for _, f := range *v.Funcs {
		f()
	}
	for _, f := range *v.DeferFuncs {
		f()
	}
}

func (v *FuncBlock) Combine(fb *FuncBlock) *FuncBlock {
	v.Funcs.Combine(fb.Funcs)
	v.DeferFuncs.Combine(fb.DeferFuncs)
	return v
}
