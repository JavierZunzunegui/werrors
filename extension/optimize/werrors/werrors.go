package werrors

import "errors"

type WrapError struct {
	payload error
	next    *WrapError
}

func (w *WrapError) LegacyError() string {
	if w.next == nil {
		return w.payload.Error()
	}
	return w.payload.Error() + ": " + w.next.LegacyError()
}

func (w *WrapError) Unwrap() error {
	if w.next == nil {
		// Avoiding nil-typed error.
		return nil
	}
	return w.next
}

func (w *WrapError) Is(err error) bool {
	return errors.Is(w.payload, err)
}

func (w *WrapError) As(i interface{}) bool {
	return errors.As(w.payload, i)
}

func Wrap(inErr, outErr error) error {
	if inErr == nil {
		return nil
	}
	if outErr == nil {
		return inErr
	}

	if _, ok := outErr.(*WrapError); ok {
		// TODO: not expanding it yet, requires a full copy of the linked error list.
		// Using it is an anti-pattern.
		panic("TODO - outErr as *WrapError")
	}

	if inWErr, ok := inErr.(*WrapError); ok {
		return &WrapError{payload: outErr, next: inWErr}
	}

	return &WrapError{payload: outErr, next: &WrapError{payload: inErr}}
}
