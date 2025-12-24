// package errortrace provides utilities to work with errors containing stack traces.
package errortrace

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mkch/gg"
	"github.com/mkch/gg/runtime2"
)

// Error is an error containing stack frames.
type Error struct {
	error
	frames     []uintptr
	moreFrames bool
}

// Error implements error interface.
func (e *Error) Error() string {
	return e.error.Error()
}

// Unwrap returns the underlying error.
func (e *Error) Unwrap() error {
	return e.error
}

// StackFrames returns the stack frames captured when the error was created.
func (e *Error) StackFrames() *runtime2.Frames {
	return runtime2.StackFromPC(e.frames, e.moreFrames)
}

// fprintErrorIndentImpl prints all Errors in the entire error chain to w.
// It is a helper function for fprintErrorIndent.
// Each level of Error is indented by an additional indent level.
// If lookForCause is true, it prints "Caused by:" before printing the error message.
// It returns whether any Error is printed, the number of bytes written, and any error encountered.
func fprintErrorIndentImpl(w io.Writer, indent string, indentLevel int, e error, lookForCause bool) (printed bool, n int, err error) {
	var nn int
	var pp bool
	// Print Error if available
	if errStack, ok := e.(*Error); ok {
		indentStr := strings.Repeat(indent, indentLevel)
		// Print "Caused by:" if needed.
		if lookForCause {
			nn, err = fmt.Fprintln(w, "\n"+indentStr+"Caused by:")
			n += nn
			if err != nil {
				return
			}
		}
		// Print error message.
		nn, err = fmt.Fprintln(w, indentStr+e.Error())
		printed = true
		n += nn
		if err != nil {
			return
		}
		// Print stack from frames.
		if frames := errStack.StackFrames(); frames != nil && len(frames.Frames) > 0 {
			// Do not print marker if only one frame and marked complete.
			var needPrintMarker = len(frames.Frames) > 1 || !frames.Complete
			nn, err = io.WriteString(w, "\n"+indentStr+gg.If(needPrintMarker, "===== STACK TRACE =====\n", ""))
			n += nn
			if err != nil {
				return
			}
			nn, err = frames.FprintIndent(w, indent, indentLevel)
			n += nn
			if err != nil {
				return
			}
			nn, err = io.WriteString(w, indentStr+gg.If(needPrintMarker, "=======================\n", "\n"))
			n += nn
			if err != nil {
				return
			}
		}
		// Look for causes via Unwrap.
		pp, nn, err = fprintErrorIndentImpl(w, indent, indentLevel+1, errStack.Unwrap(), true)
		printed = printed || pp
		n += nn
		return
	}

	// Unwrap cause.
	if uw, ok := e.(interface{ Unwrap() error }); ok {
		if cause := uw.Unwrap(); cause != nil {
			pp, nn, err = fprintErrorIndentImpl(w, indent, indentLevel, cause, lookForCause)
			printed = printed || pp
			n += nn
		}
		return
	}
	// Unwrap causes.
	if uw, ok := e.(interface{ Unwrap() []error }); ok {
		for _, cause := range uw.Unwrap() {
			pp, nn, err = fprintErrorIndentImpl(w, indent, indentLevel, cause, lookForCause)
			printed = printed || pp
			n += nn
			if err != nil {
				return
			}
		}
	}
	return
}

// fprintErrorIndent prints all [Error] instances in the entire
// error chain to w with tab("\t") indentation.
// If no Error is found in the chain, it prints the error message only.
func fprintErrorIndent(w io.Writer, e error) (n int, err error) {
	const indent = "\t"
	printed, nn, err := fprintErrorIndentImpl(w, indent, 0, e, false)
	n += nn
	if err != nil {
		return
	}
	if !printed {
		// No Error found in the chain, print the error message only.
		nn, err = fmt.Fprintln(w, e)
		n += nn
		if err != nil {
			return
		}
	}
	return
}

// Fprint prints all [Error] instances in the entire error chain rooted
// at e to w.
// If no [Error] is found in the chain, it prints the error message only.
// This function is intended to be used for printing [Error] instances in
// a wrapper.
// See example of [WithStack].
func Fprint(w io.Writer, e error) (n int, err error) {
	return fprintErrorIndent(w, e)
}

// Print calls [Fprint]([os.Stderr], e).
func Print(e error) (n int, err error) {
	return Fprint(os.Stderr, e)
}

// Sprint returns the string representation of the output of [Fprint].
func Sprint(e error) string {
	var sb strings.Builder
	Fprint(&sb, e) // string.Builder.Write never returns error.
	return sb.String()
}

// Format implements [fmt.Formatter].
// With verb 'v' and the '+' flag, it prints all Error
// instances in the error chain, each with its message and stack trace.
// Otherwise, it falls back to the default formatting.
// See the example of [WithStack].
func (e *Error) Format(f fmt.State, verb rune) {
	if verb == 'v' && f.Flag('+') {
		fprintErrorIndent(f, e)
		return
	}
	// Fallback to default formatting.
	fmt.Fprintf(f, fmt.FormatString(f, verb), e.error)
}

// WithStackFrames returns an Error that wraps err and contains stack frames
// start from the caller of WithStackFrames.
// The argument skip is the number of stack frames to skip before recording, with 0 identifying
// starting from the caller of WithStackFrames.
// The number of stack frames captured is limited to nFrames.
// If nFrames is less than or equal to zero, a reasonable default value is used.
// If forceComplete is true the stack frames capture is considered complete even
// if more than nFrames stack frames exist.
// If err is nil, nil is returned.
func WithStackFrames(err error, skip, nFrames int, forceComplete bool) *Error {
	if err == nil {
		return nil
	}
	if nFrames <= 0 {
		nFrames = maxStackDepth
	}
	pcs, more := runtime2.Callers(skip+1, nFrames) // skip [withStackN].
	return &Error{
		error:      err,
		frames:     pcs,
		moreFrames: more && !forceComplete,
	}
}

const maxStackDepth = 32

// WithStack returns an Error that wraps err and contains stack frames start
// from the caller of WithStack.
// The number of stack frames captured has a reasonable limit. If more stack frames exist,
// a marker message is included in the stack trace.
// If the maximum number of stack frames is in consideration, use [WithStackFrames] instead.
// If err is nil, nil is returned.
func WithStack(err error) *Error {
	return WithStackFrames(err, 1, maxStackDepth, false) // Skip [WithStack].
}

// WithFileLine returns an Error that wraps err and contains
// the stack frame of the caller of WithFileLine.
// If err is nil, nil is returned. See example of [WithStack].
func WithFileLine(err error) *Error {
	return WithStackFrames(err, 1, 1, true) // Skip [WithFileLine], capture only one frame, complete.
}

// ErrorfStack is like [fmt.Errorf] but returns an Error that contains stack frames start
// from the caller of ErrorfStack.
// ErrorfStack acts as wrapping the return value of fmt.ErrorfStack(format, args...) with [WithStack].
func ErrorfStack(format string, args ...any) *Error {
	err := fmt.Errorf(format, args...)
	return WithStackFrames(err, 1, maxStackDepth, false) // Skip [ErrorfStack].
}

// ErrorfFileLine is like [fmt.Errorf] but returns an Error that contains
// the stack frame of the caller of ErrorfFileLine.
// ErrorfFileLine acts as wrapping the return value of fmt.Errorf(format, args...) with [WithFileLine].
func ErrorfFileLine(format string, args ...any) *Error {
	err := fmt.Errorf(format, args...)
	return WithStackFrames(err, 1, 1, true) // Skip [ErrorfFileLine], capture only one frame, complete.
}
