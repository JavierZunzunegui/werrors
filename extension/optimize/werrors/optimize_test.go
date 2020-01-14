package werrors_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/JavierZunzunegui/werrors/extension/optimize/werrors"
)

func scenarios() []struct {
	name string
	wErr *werrors.WrapError
} {
	wrapNTimes := func(n int) *werrors.WrapError {
		err := errors.New("base")
		for i := 0; i < n; i++ {
			err = werrors.Wrap(err, errors.New("wrapper "+strconv.Itoa(i)))
		}
		return err.(*werrors.WrapError)
	}

	return []struct {
		name string
		wErr *werrors.WrapError
	}{
		{
			"depth-1",
			wrapNTimes(1),
		},
		{
			"depth-2",
			wrapNTimes(2),
		},
		{
			"depth-3",
			wrapNTimes(3),
		},
		{
			"depth-5",
			wrapNTimes(5),
		},
		{
			"depth-10",
			wrapNTimes(10),
		},
		{
			"depth-20",
			wrapNTimes(20),
		},
	}
}

func BenchmarkWrapError_LegacyError(b *testing.B) {
	for _, scenario := range scenarios() {
		scenario := scenario

		b.Run(scenario.name, func(b *testing.B) {
			b.ResetTimer()
			var msg string
			for i := 0; i < b.N; i++ {
				msg = scenario.wErr.LegacyError()
			}
			_ = msg // To avoid it being optimized away.
		})
	}
}

func BenchmarkWrapError_Error(b *testing.B) {
	for _, scenario := range scenarios() {
		scenario := scenario

		b.Run(scenario.name, func(b *testing.B) {
			b.ResetTimer()
			var msg string
			for i := 0; i < b.N; i++ {
				msg = scenario.wErr.Error()
			}
			_ = msg // To avoid it being optimized away.
		})
	}
}
