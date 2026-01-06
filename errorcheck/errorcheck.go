package errorcheck

// Handler is a function type for handling errors.
type Handler func(error)

// Must checks the error and returns the value if err is nil.
// If err is not nil, it calls the provided handler and panics.
// The intended usage is to wrap this function with a custom handler as a
// function that only takes (T, error) and returns T.
func Must[T any](handler Handler, v T, err error) T {
	MustOK(handler, err)
	return v
}

// Must2 is like [Must] but accepts and returns two values.
func Must2[T1 any, T2 any](handler Handler, v1 T1, v2 T2, err error) (T1, T2) {
	MustOK(handler, err)
	return v1, v2
}

// Must3 is like [Must] but accepts and returns three values.
func Must3[T1 any, T2 any, T3 any](handler Handler, v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	MustOK(handler, err)
	return v1, v2, v3
}

// Must4 is like [Must] but accepts and returns four values.
func Must4[T1 any, T2 any, T3 any, T4 any](handler Handler, v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	MustOK(handler, err)
	return v1, v2, v3, v4
}

// MustOK checks the error and calls the provided handler if err is not nil.
func MustOK(handler Handler, err error) {
	if err != nil {
		handler(err)
		panic(err)
	}
}
