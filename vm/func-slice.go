package vm

import ()

type (
	NoNameFuncSlice []func()
)

func NewNoNameFuncSlice() *NoNameFuncSlice {
	return new(NoNameFuncSlice)
}

func (v NoNameFuncSlice) IsBlank() bool {
	return len(v) != 0
}

func (v NoNameFuncSlice) Do() {
	for _, f := range v {
		f()
	}
}

func (v *NoNameFuncSlice) Append(f ...func()) {
	*v = append(*v, f...)
}

func (v *NoNameFuncSlice) Combine(fs *NoNameFuncSlice) *NoNameFuncSlice {
	for _, f := range *fs {
		*v = append(*v, f)
	}
	return v
}

func (v NoNameFuncSlice) Compact() func() {
	return func() {
		for _, f := range v {
			f()
		}
	}
}
