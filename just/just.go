package just

// Err discards the value and returns the error
//
// Example:
//
//	err := errors.Join(
//		just.Err(io.Copy(w,r)),
//		just.Err(fmt.Fprint(w, "bye"),
//		)
func Err[T any](_ T, e error) error {
	return e
}
