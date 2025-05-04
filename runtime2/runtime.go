package runtime2

import (
	"fmt"
	"log/slog"
	"runtime"
	"strings"
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

// String returns the string representation of f.
func (f *Frames) String() string {
	if f == nil {
		return "(no stack)"
	}
	var b strings.Builder
	for i, frame := range f.Frames {
		fmt.Fprintf(&b, "called from %s (%s:%d)", frame.Function, frame.File, frame.Line)
		if i != len(f.Frames)-1 {
			b.WriteByte('\n')
		}
	}
	if !f.Complete {
		if b.Len() > 0 {
			b.WriteByte('\n')
		}
		fmt.Fprintf(&b, "(rest of stack elided)")
	}
	return b.String()
}

// Stack returns the call stack of the caller. If no frame is available, Stack returns nil.
// If skip is 0, the returned frames start from the caller of Stack,
// if skip is 1, they start from the caller of the caller of Stack, etc.
// If nFrames is greater than 0, at most nFrames frames will be returned, otherwise all frames will be returned.
// From Go/src/log/slog/value.go.
func Stack(skip, nFrames int) *Frames {
	skip += 2 // skip [Callers, Stack]
	var pcs []uintptr
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
		return nil
	}
	var ret = Frames{Complete: true}
	frames := runtime.CallersFrames(pcs[:n])
	i := 0
	for {
		frame, more := frames.Next()
		ret.Frames = append(ret.Frames, frame)
		if !more {
			break
		}
		i++
		if nFrames > 0 && i >= nFrames {
			ret.Complete = false
			break
		}
	}
	return &ret
}
