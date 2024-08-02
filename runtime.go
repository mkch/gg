package gg

import (
	"log/slog"
	"runtime"
)

// Source returns the SourceLine of the caller.
// If the location is unavailable, it returns a Source with zero fields.
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
