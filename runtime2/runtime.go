package runtime2

import (
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strings"

	"github.com/mkch/gg"
)

// Source returns the SourceLine of the caller.
// If the location is unavailable, it returns a Source with zero fields.
// From Go/src/log/slog/record.go
func Source() slog.Source {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Source]

	fs := runtime.CallersFrames(pcs[:])
	f, _ := fs.Next()
	return slog.Source{
		Function: f.Function,
		File:     f.File,
		Line:     f.Line,
	}
}

// Frames is call stack frames used by [Stack].
type Frames struct {
	Frames   []runtime.Frame
	Complete bool // true if Frames is not truncated due to nFrames of Stack.
}

// FprintIndent formats the stack frames in f to w with indentation.
// Each frame is printed in two lines: the function name with base indentation (indent repeated indentLevel times),
// and the file location with one additional indent level (indent repeated indentLevel+1 times).
func (f *Frames) FprintIndent(w io.Writer, indent string, indentLevel int) (n int, err error) {
	indentStr := strings.Repeat(indent, indentLevel)
	if f == nil {
		return fmt.Fprintf(w, "%s(no stack)\n", indentStr)
	}
	const unknown = "???"
	for _, frame := range f.Frames {
		funcName := gg.If(len(frame.Function) > 0, frame.Function, unknown)
		fileName := gg.If(len(frame.File) > 0, frame.File, unknown)
		var nn int
		// This offset is not the same as the one printed by panic().
		// The reason for this is unclear.
		// Don't use it right now.
		//
		// offset := gg.IfFunc(frame.Func != nil,
		// 	func() uintptr { return frame.PC - frame.Func.Entry() },
		// 	func() uintptr { return 0 })
		// nn, err = fmt.Fprintf(w, "%s%s()\n%s%s:%d +0x%x\n",
		// 	indentStr, funcName,
		// 	indentStr+indent, fileName, frame.Line, offset)
		nn, err = fmt.Fprintf(w, "%s%s()\n%s%s:%d\n",
			indentStr, funcName,
			indentStr+indent, fileName, frame.Line)
		if err != nil {
			return
		}
		n += nn
	}
	if !f.Complete {
		var nn int
		nn, err = fmt.Fprintf(w, "%s(rest of stack elided)\n", indentStr)
		if err != nil {
			return
		}
		n += nn
	}
	return
}

// String returns the string representation of f.
func (f *Frames) String() string {
	var b strings.Builder
	_, _ = f.FprintIndent(&b, "", 0)
	return b.String()
}

// Callers returns the program counters of function invocations on the calling goroutine's stack.
// The argument skip is the number of stack frames to skip before recording in pcs, with 0 identifying the
// frame of the caller of [Callers].
// If nFrames is greater than 0, at most nFrames PCs will be returned, otherwise all PCs will be returned.
// If there are more PCs available than returned, more is true.
// From Go/src/log/slog/value.go.
func Callers(skip, nFrames int) (pcs []uintptr, more bool) {
	skip += 2 // skip [runtime.Callers, Callers]
	var n int
	if nFrames > 0 {
		pcs = make([]uintptr, nFrames+1)
		n = runtime.Callers(skip, pcs)
	} else {
		pcs = make([]uintptr, 5)
		n = runtime.Callers(skip, pcs)
		for n == len(pcs) {
			pcs = append(pcs, make([]uintptr, len(pcs))...) // double
			n = runtime.Callers(skip, pcs)
		}
	}
	if n == 0 {
		return nil, false
	}
	if nFrames > 0 && n > nFrames {
		return pcs[:nFrames], true
	}
	return pcs[:n], false
}

// StackFromPC returns the [Frames] corresponding to the given pcs.
func StackFromPC(pcs []uintptr, more bool) *Frames {
	if len(pcs) == 0 {
		return nil
	}
	var ret = Frames{Complete: !more}
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		ret.Frames = append(ret.Frames, frame)
		if !more {
			break
		}
	}
	return &ret
}

// Stack returns the call stack of the caller. If no frame is available, Stack returns nil.
// If skip is 0, the returned frames start from the caller of Stack,
// if skip is 1, they start from the caller of the caller of Stack, etc.
// If nFrames is greater than 0, at most nFrames frames will be returned, otherwise all frames will be returned.
// From Go/src/log/slog/value.go.
func Stack(skip, nFrames int) *Frames {
	skip += 1 // skip [Stack]
	return StackFromPC(Callers(skip, nFrames))
}
