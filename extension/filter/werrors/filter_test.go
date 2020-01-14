package werrors_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/JavierZunzunegui/werrors/extension/filter/werrors"
)

func TestWrapError_Filter(t *testing.T) {
	err := errors.New("base")
	err = werrors.Wrap(err, errors.New("wrap 1"))
	err = werrors.Wrap(err, errors.New("wrap 2"))
	wErr := err.(*werrors.WrapError)

	if got, expected := wErr.Filter(func(e error) bool { return true }), "wrap 2: wrap 1: base"; got.Error() != expected {
		t.Errorf("*WrapError.Filter(true): expected %q, got %q", expected, got)
	}

	if got := wErr.Filter(func(e error) bool { return false }); got != nil {
		t.Errorf("*WrapError.Filter(false): expected nil, got %q", got)
	}

	if got, expected := wErr.Filter(func(e error) bool { return e.Error() == "base" }), "base"; got.Error() != expected {
		t.Errorf("*WrapError.Filter(e.Error()==\"base\"): expected %q, got %q", expected, got)
	}

	if got, expected := wErr.Filter(func(e error) bool { return strings.Contains(e.Error(), "wrap") }), "wrap 2: wrap 1"; got.Error() != expected {
		t.Errorf("*WrapError.Filter(Contains(e.Error(),\"wrap\")): expected %q, got %q", expected, got)
	}
}
