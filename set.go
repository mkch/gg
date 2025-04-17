package gg

// Set is a generic set data structure where each element of type T is unique.
type Set[T comparable] map[T]struct{}

// Add adds the given value to the set.
// If the value is not already present in the set, it will be added.
// If the value already exists, this method does nothing.
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// Contains checks if the given value is present in the set.
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}
