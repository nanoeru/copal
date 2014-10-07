package val

type Stack struct {
	data     []Val
	index    int
	length   int
	capacity int
}

func NewStack() *Stack {
	capacity := 16
	return &Stack{
		data:     make([]Val, capacity),
		index:    -1,
		length:   capacity,
		capacity: capacity,
	}
}

func (s *Stack) Len() int {
	return s.index + 1
}

func (s *Stack) Push(value Val) {
	s.index++
	if s.index < s.length {
		s.data[s.index] = value
	} else {
		s.data = append(s.data, value)
		s.length++
		//s.length = len(s.data)
		//s.capacity = cap(s.data)
	}
}

func (s *Stack) Pop() (value Val) {
	//	unsafe
	//value = s.data[s.index]
	//s.index--
	//return
	//	safe
	if s.index >= 0 {
		value = s.data[s.index]
		s.index--
		return
	}
	return nil
}

func (s *Stack) Set(value Val) {
	//	unsafe
	//s.data[s.index] = value
	//return
	//	safe
	if s.index >= 0 {
		s.data[s.index] = value
		return
	}
	panic("set: null stack")
}

func (s *Stack) Get() (value Val) {
	//	unsafe
	//value = s.data[s.index]
	//return
	//	safe
	if s.index >= 0 {
		value = s.data[s.index]
		return
	}
	panic("get: null stack")
	return nil
}
