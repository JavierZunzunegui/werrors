package werrors

import (
	"runtime"
	"strconv"
	"strings"
)

type FrameError struct {
	frame [1]uintptr
}

func FWrap(inErr, outErr error) error {
	if inErr == nil {
		return nil
	}
	if outErr == nil {
		return outErr
	}

	pc := [1]uintptr{}

	// Never 0.
	_ = runtime.Callers(2, pc[:])

	err := Wrap(inErr, &FrameError{
		frame: pc,
	})

	return Wrap(err, outErr)
}

func (f FrameError) Error() string {
	frames := runtime.CallersFrames(f.frame[:])
	frame, _ := frames.Next()

	file := frame.File
	if i := strings.LastIndex(frame.File, "/"); i != -1 {
		file = file[i+1:]
	}

	return file + "(" + strconv.Itoa(frame.Line) + ")"
}
