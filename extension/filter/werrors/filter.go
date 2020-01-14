package werrors

func (w *WrapError) Filter(f func(error) bool) *WrapError {
	out := &WrapError{}
	last := out

	for ; w != nil; w = w.next {
		if f(w.payload) {
			last.next = &WrapError{payload: w.payload}
			last = last.next
		}
	}

	return out.next
}
