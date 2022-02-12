package slices

// Map maps a slice into another slice using the provided
// `mapper`. Every element of the `input` is passed to the
// `mapper`, the returned value is added to the output.
func Map[A, B any](mapper func(A) B) func([]A) []B {
	return func(input []A) []B {
		output := make([]B, len(input))
		for i, v := range input {
			output[i] = mapper(v)
		}
		return output
	}
}

// Reduce reduces a slice down to a single value using the
// supplied `reducer`. `start` is the initial value to use
// when reducing.
// Every element of the `input` is passed to the `reducer`,
// the returned value is used as the aggregate of the next
// `reducer` call. The initial aggregate passed to `reducer`
//  is `start`. The final value returned from `reducer`
//  is the output.
func Reduce[A, B any](reducer func(A, B) B, start B) func([]A) B {
	return func(input []A) B {
		reduced := start
		for _, v := range input {
			reduced = reducer(v, reduced)
		}
		return reduced
	}
}

// Filter removes any element that returns `false` from the
// supplied `filter` method.
// Every element of `input` is passed to `filter`. If `filter`
// returns `true` the element is kept in the output.
func Filter[A any](filter func(A) bool) func([]A) []A {
	return func(input []A) []A {
		var output []A
		for _, v := range input {
			if filter(v) {
				output = append(output, v)
			}
		}
		return output
	}
}
