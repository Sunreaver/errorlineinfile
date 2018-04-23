package errorlineinfile

import (
	"fmt"
	"runtime"
)

// Error define Error
type Error struct {
	ok   bool
	line int
	file string
	tips string
}

// New will 新建Error
func New(tips string) (e *Error) {
	defer func() {
		_, e.file, e.line, e.ok = runtime.Caller(2)
	}()
	return &Error{
		tips: tips,
	}
}

func (e *Error) Error() (out string) {
	if e.ok {
		out = fmt.Sprintf("%s:%d ==> %v", e.file, e.line, e.tips)
	} else {
		out = fmt.Sprintf("%s", e.tips)
	}
	return
}
