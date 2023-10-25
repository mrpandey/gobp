package util

// Returns a pointer that holds the value v. Returned pointer does not point to original value.
// Usage: Ptr(4), Ptr(uint(4)), Ptr(false).
func Ptr[T any](v T) *T {
	return &v
}
