package werrors

import (
	"bytes"
)

func (w *WrapError) Error() string {
	var buf bytes.Buffer
	w.errorToBuf(&buf)

	return buf.String()
}

func (w *WrapError) errorToBuf(buf *bytes.Buffer) {
	buf.WriteString(w.payload.Error())

	if w.next == nil {
		return
	}

	buf.WriteString(": ")
	w.next.errorToBuf(buf)
}
