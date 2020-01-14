package werrors

func (w *WrapError) Format(sep string) string {
	if w.next == nil {
		return w.payload.Error()
	}
	return w.payload.Error() + sep + w.next.Format(sep)
}
