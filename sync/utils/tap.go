package utils

// Tap is a helper that lets you observe a value in a chain
// without modifying it. The value is returned, unmodified so the
// chain can continue.
func Tap[A any](tapper func(A)) func(A) A {
	return func(a A) A {
		tapper(a)
		return a
	}
}
