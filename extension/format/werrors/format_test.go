package werrors_test

import (
	"errors"
	"testing"

	"github.com/JavierZunzunegui/werrors/extension/format/werrors"
)

func TestWrapError_Format(t *testing.T) {
	err := errors.New("base")
	err = werrors.Wrap(err, errors.New("wrap 1"))
	err = werrors.Wrap(err, errors.New("wrap 2"))
	wErr := err.(*werrors.WrapError)

	if got, expected := wErr.Format(": "), "wrap 2: wrap 1: base"; got != expected {
		t.Errorf("*WrapError.Format(\": \"): expected %q, got %q", expected, got)
	}

	if got, expected := wErr.Format(" - "), "wrap 2 - wrap 1 - base"; got != expected {
		t.Errorf("*WrapError.Format(\" - \"): expected %q, got %q", expected, got)
	}
}
