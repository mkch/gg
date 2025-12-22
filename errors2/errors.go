// package errors2 provides error types and functions to work with errors containing stack traces.
package errors2

import (
	"fmt"
	"io"
	"strings"

	"github.com/mkch/gg/runtime2"
)

// ErrorWithStack is an error that contains stack frames.
type ErrorWithStack struct {
	error
	frames     []uintptr
	moreFrames bool
}

// Unwrap returns the underlying error.
func (e *ErrorWithStack) Unwrap() error {
	return e.error
}

// StackFrames returns the stack frames captured when the error was created.
func (e *ErrorWithStack) StackFrames() *runtime2.Frames {
	return runtime2.StackFromPC(e.frames, e.moreFrames)
}

// fprintErrorChainIndent prints the entire error chain to w with indentation.
// If indent is empty, a tab character("\t") is used as the indent string.
// Each level of the error chain is indented by an additional indent level.
// *ErrorWithStack in the chain will be printed with [ErrorWithStack.FprintIndent].
func fprintErrorChainIndent(w io.Writer, indent string, indentLevel int, e error) (n int, err error) {
	if e == nil {
		return io.WriteString(w, "<nil>\n")
	}
	if indent == "" {
		indent = "\t"
	}
	indentStr := strings.Repeat(indent, indentLevel)
	var nn int
	// Print the error message.
	if nn, err = fmt.Fprintf(w, "%s%T: %s\n", indentStr, e, e.Error()); err != nil {
		return nn, err
	}
	// Print stack if available.
	if errStack, ok := e.(*ErrorWithStack); ok {
		nn, err = errStack.fprintStack(w, indent, indentLevel)
		if err != nil {
			return
		}
	}
	causedBy := "\n" + indentStr + "Caused by:\n"
	// Unwrap cause.
	if uw, ok := e.(interface{ Unwrap() error }); ok {
		if cause := uw.Unwrap(); cause != nil {
			nn, err = io.WriteString(w, causedBy)
			n += nn
			if err != nil {
				return
			}
			nn, err = fprintErrorChainIndent(w, indent, indentLevel+1, cause)
			n += nn
			if err != nil {
				return
			}
		}
		return
	}
	// Unwrap causes.
	if uw, ok := e.(interface{ Unwrap() []error }); ok {
		nn, err = io.WriteString(w, causedBy)
		n += nn
		if err != nil {
			return
		}
		for _, cause := range uw.Unwrap() {
			nn, err = fprintErrorChainIndent(w, indent, indentLevel+1, cause)
			n += nn
			if err != nil {
				return
			}
		}
	}
	return
}

// fprintStack prints the stack frames of e to w with indentation.
func (e *ErrorWithStack) fprintStack(w io.Writer, indent string, indentLevel int) (n int, err error) {
	indentStr := strings.Repeat(indent, indentLevel)
	if len(e.frames) == 0 {
		return
	}
	var nn int
	frames := e.StackFrames()
	nn, err = io.WriteString(w, "\n"+indentStr+"===== STACK TRACE =====\n")
	n += nn
	if err != nil {
		return
	}
	nn, err = frames.FprintIndent(w, indent, indentLevel)
	n += nn
	if err != nil {
		return
	}
	nn, err = io.WriteString(w, indentStr+"=======================\n")
	n += nn
	return
}

// Format implements [fmt.Formatter] interface.
// When the verb is 'v' and the '+' flag is specified, it prints the error message
// and the string representation of stack frames to f.
// Otherwise, it falls back to the default formatting.
// See the example of [WithStack].
func (e *ErrorWithStack) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('+') {
		fprintErrorChainIndent(f, "\t", 0, e)
		return
	}
	// Fallback to default formatting.
	fmt.Fprintf(f, fmt.FormatString(f, verb), e)
}

// WithStackFrames returns an ErrorWithStack that wraps err and contains stack frames
// start from the caller of WithStackFrames.
// The argument skip is the number of stack frames to skip before recording, with 0 identifying the
// frame of the caller of WithStackFrames.
// The number of stack frames captured is limited to nFrames.
// If nFrames is less than or equal to zero, a reasonable default value is used.
// If err is nil, nil is returned.
func WithStackFrames(err error, skip, nFrames int) *ErrorWithStack {
	if err == nil {
		return nil
	}
	if nFrames <= 0 {
		nFrames = maxStackDepth
	}
	pcs, more := runtime2.Callers(skip+1, nFrames) // skip [withStackN].
	return &ErrorWithStack{
		error:      err,
		frames:     pcs,
		moreFrames: more,
	}
}

const maxStackDepth = 32

// WithStack returns an ErrorWithStack that wraps err and contains stack frames start
// from the caller of WithStack.
// The number of stack frames captured has a reasonable limit.
// If the maximum number of stack frames is in consideration, use [WithStackFrames] instead.
// If err is nil, nil is returned.
func WithStack(err error) *ErrorWithStack {
	return WithStackFrames(err, 1, maxStackDepth) // Skip [WithStack].
}

// Errorf is like [fmt.Errorf] but returns an ErrorWithStack that contains stack frames start
// from the caller of [Errorf].
// Errorf acts as wrapping the return value of fmt.Errorf(format, args...) with [WithStack].
func Errorf(format string, args ...any) *ErrorWithStack {
	err := fmt.Errorf(format, args...)
	return WithStackFrames(err, 1, maxStackDepth) // Skip [Errorf].
}
