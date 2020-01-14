package werrors_test

import (
	"errors"
	"testing"

	"github.com/JavierZunzunegui/werrors/extension/frame/werrors"
)

func TestFWrap(t *testing.T) {
	wErr := werrors.FWrap(errors.New("base"), errors.New("wrapper"))

	// NOTE: This includes a line number and file name, as such it is very sensible to refactoring.
	const expected = "wrapper: frame_test.go(11): base"
	if wErr.Error() != expected {
		t.Errorf("FWrap(err,err): expected %q, got %q", expected, wErr.Error())
	}
}
