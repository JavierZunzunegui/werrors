package werrors_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/JavierZunzunegui/werrors"
)

func TestWrap(t *testing.T) {
	t.Run("double nil input", func(t *testing.T) {
		if wErr := werrors.Wrap(nil, nil); wErr != nil {
			t.Errorf("Wrap(nil, nil): expecting nil, got: %q", wErr)
		}
	})

	t.Run("nil inError", func(t *testing.T) {
		if wErr := werrors.Wrap(nil, errors.New("some error")); wErr != nil {
			t.Errorf("Wrap(nil, err): expecting nil, got: %q", wErr)
		}
	})

	t.Run("nil outError", func(t *testing.T) {
		wErr := werrors.Wrap(errors.New("some error"), nil)
		if wErr == nil {
			t.Error("Wrap(err, nil): expecting non-nil, got nil")
			return
		}
		if wErr.Error() != "some error" {
			t.Errorf("Wrap(err, nil).Error(): expecting %q, got %q", "some error", wErr)
		}
	})

	t.Run("unwrapped inError", func(t *testing.T) {
		wErr := werrors.Wrap(errors.New("some error"), errors.New("wrapping error"))
		if wErr == nil {
			t.Error("Wrap(err, err): expecting non-nil, got nil")
			return
		}

		const expected = "wrapping error: some error"
		if wErr.Error() != expected {
			t.Errorf("Wrap(err, err).Error(): expecting %q, got %q", expected, wErr)
		}

		if _, ok := wErr.(*werrors.WrapError); !ok {
			t.Errorf("Wrap(err, err): expecting type *WrapError, got %T", wErr)
		}
	})

	t.Run("wrapped inError", func(t *testing.T) {
		inErr := werrors.Wrap(errors.New("some error"), errors.New("wrapping error 1"))
		wErr := werrors.Wrap(inErr, errors.New("wrapping error 2"))
		if wErr == nil {
			t.Error("Wrap(*WrapError, err): expecting non-nil, got nil")
			return
		}

		const expected = "wrapping error 2: wrapping error 1: some error"
		if wErr.Error() != expected {
			t.Errorf("Wrap(*WrapError, err).Error(): expecting %q, got %q", expected, wErr)
		}

		if _, ok := wErr.(*werrors.WrapError); !ok {
			t.Errorf("Wrap(*WrapError, err): expecting type *WrapError, got %T", wErr)
		}
	})
}

type IError struct {
	error
}

func (iErr IError) Is(err error) bool {
	// Is(err) is true if both errors have the same prefix.
	// It is a bit absurd, but good enough for testing.
	iPrefix := iErr.Error()
	if s := strings.Split(iPrefix, " "); len(s) > 0 {
		iPrefix = s[0]
	}

	prefix := err.Error()
	if s := strings.Split(prefix, " "); len(s) > 0 {
		prefix = s[0]
	}

	return iPrefix == prefix
}

func TestWrapError_Is(t *testing.T) {
	t.Run("simple errors", func(t *testing.T) {
		inErr := errors.New("some error")
		outErr := errors.New("wrapping error")
		wErr := werrors.Wrap(inErr, outErr)
		if !errors.Is(wErr, inErr) {
			t.Error("Is(Wrap(err1, err2), err1): should be true")
		}
		if !errors.Is(wErr, outErr) {
			t.Error("Is(Wrap(err1, err2), err2): should be true")
		}
		if errors.Is(wErr, errors.New("unknown error")) {
			t.Error("Is(Wrap(err1, err2), err3): should be false")
		}
	})

	t.Run("Is errors", func(t *testing.T) {
		inErr := IError{errors.New("some error")}
		outErr := IError{errors.New("wrapping error")}
		wErr := werrors.Wrap(inErr, outErr)
		if !errors.Is(wErr, inErr) {
			t.Error("Is(Wrap(IError1, IError2), IError1): should be true")
		}
		if !errors.Is(wErr, outErr) {
			t.Error("Is(Wrap(IError1, IError2), IError2): should be true")
		}
		if errors.Is(wErr, errors.New("unknown error")) {
			t.Error("Is(Wrap(IError1, IError2), err3): should be false")
		}
		if !errors.Is(wErr, errors.New("some other error")) {
			t.Error("Is(Wrap(IError1, IError2), err3~IError1): should be true")
		}
		if !errors.Is(wErr, errors.New("wrapping other error")) {
			t.Error("Is(Wrap(IError1, IError2), err2~IError1): should be true")
		}
	})
}

type UnknownError struct {
	error
}

func TestWrapError_As(t *testing.T) {
	t.Run("simple errors", func(t *testing.T) {
		inErr := IError{errors.New("some error")}
		outErr := errors.New("wrapping error")
		wErr := werrors.Wrap(inErr, outErr)
		if holder := errors.New(""); !errors.As(wErr, &holder) {
			t.Error("Is(Wrap(err1, IError2), *type(err1)): should be true")
		}

		if holder := error(IError{}); !errors.As(wErr, &holder) {
			t.Error("Is(Wrap(err1, IError2), *IError2): should be true")
		}
		if holder := error(UnknownError{}); !errors.As(wErr, &holder) {
			t.Error("Is(Wrap(err1, IError2), *UnknownError): should be false")
		}
	})
}
