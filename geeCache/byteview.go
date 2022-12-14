package geeCache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneByte(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneByte(b []byte) (cb []byte) {
	cb = make([]byte, len(b))
	copy(cb, b)
	return
}
