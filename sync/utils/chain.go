package utils

// Chain2 composes two functions together. It returns a function
// that can be called with input to call `fn1` and `fn2` on.
//
// The goal of Chain2 is to make chaining functions a bit more readable
// and to allow composing functions together into new functions that
// can achieve some sort of goal.
// For example:
// ```
// func SumAndString(input []int) string {
//    return Chain2(slices.Reduce(...), func(i int) string { return fmt.Sprint(i) })
// }
// ```
func Chain2[A, B, C any](fn1 func(A) B, fn2 func(B) C) func(A) C {
	return func(input A) C {
		return fn2(fn1(input))
	}
}

func Chain3[A, B, C, D any](fn1 func(A) B, fn2 func(B) C, fn3 func(C) D) func(A) D {
	return func(input A) D {
		return fn3(fn2(fn1(input)))
	}
}

func Chain4[A, B, C, D, E any](fn1 func(A) B, fn2 func(B) C, fn3 func(C) D, fn4 func(D) E) func(A) E {
	return func(input A) E {
		return fn4(fn3(fn2(fn1(input))))
	}
}
