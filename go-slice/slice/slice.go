package slice

import (
	"unsafe"
)

type Slice struct {
	cap int
	len int
	array unsafe.Pointer
}

func Make(t string, l int, c ...int) *Slice {
	s := &Slice{
		len: l,
		cap: l,
	}
	if len(c) > 0 {
		s.cap = c[0]
	}
	return s
}

func (s *Slice) Len() {
}

func (s *Slice) Cap() {
}

func (s *Slice) Copy() {
}

func (s *Slice) Append() {
}

func (s *Slice) Slice() {
}

func makeArray(t string, length uint) unsafe.Pointer {
	var ptr unsafe.Pointer
	switch t {
	case "int":
		intArr := [0]int{}
		ptr = unsafe.Pointer(&intArr)
	}
	return ptr
}
