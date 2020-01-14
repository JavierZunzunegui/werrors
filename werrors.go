// Package werrors is a proposed extension to package errors.
package werrors

import "errors"

// WrapError is the main element of this proposal.
// It represents a linked list of errors, and provides the wrapping functionality.
type WrapError struct {
	payload error
	next    *WrapError
}

// Error provides a default format to how wrapped errors will be serialised:
// "{err1}: {err2}: ... : {errN}"
func (w *WrapError) Error() string {
	if w.next == nil {
		return w.payload.Error()
	}
	return w.payload.Error() + ": " + w.next.Error()
}

// Unwrap is required by the current error wrapping functionality in pkg/errors.
// It's existence is unfortunate, this proposal describes why we should move away from it.
// It returns either a non-nil *WrapError or nil.
func (w *WrapError) Unwrap() error {
	if w.next == nil {
		// Avoiding nil-typed error.
		return nil
	}
	return w.next
}

// Is is the custom method required by pkg/errors.Is.
func (w *WrapError) Is(err error) bool {
	return errors.Is(w.payload, err)
}

// As is the custom method required by pkg/errors.As.
func (w *WrapError) As(i interface{}) bool {
	return errors.As(w.payload, i)
}

// Wrap is the only way wrapped errors are to be produced.
// The first argument is the error being wrapped, the second is the error being added to it.
// It returns either a non-nil *WrapError or nil.
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
