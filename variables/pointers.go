package variables

// ToPtr converts a value to a pointer.
//
// Arguments:
//   - v: the value to convert to a pointer
//
// Returns:
//   - the pointer to the value
func ToPtr[T any](v T) *T {
	return &v
}
