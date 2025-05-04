package slices2

// Filter creates a new slice which contains values from source
// that pass the keep function.
func Filter[S ~[]E, E any](source S, keep func(E) bool) (result S) {
	for _, v := range source {
		if keep(v) {
			result = append(result, v)
		}
	}
	return
}

// Map creates a new slice populated with the results of calling f
// on every element in the source slice.
func Map[E1 any, E2 any](source []E1, f func(E1) E2) (result []E2) {
	for _, v := range source {
		result = append(result, f(v))
	}
	return
}

// Reduce executes a reducer callback function on each element of the source slice,
// in order, passing in the return value from the calculation on the preceding element.
// The final result is the return value of running the reducer across all elements of
// source slice.
// The first time that the reducer is run, accVal is initValue.
func Reduce[E, R any](source []E, reducer func(accVal R, curVal E, index int) R, initVal R) (result R) {
	result = initVal
	for i, e := range source {
		result = reducer(result, e, i)
	}
	return
}

// Fill fills slice s with value v.
func Fill[E any](s []E, v E) {
	for i := range s {
		s[i] = v
	}
}

// Repeat returns a new slice whose content is v repeated n times.
func Repeat[E any](v E, n int) []E {
	result := make([]E, n)
	Fill(result, v)
	return result
}
