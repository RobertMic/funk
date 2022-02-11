package slices

func Map[A any, B any](mapper func(A) B) func([]A) []B {
	return func(input []A) []B {
		output := make([]B, len(input))
		for i, v := range input {
			output[i] = mapper(v)
		}
		return output
	}
}

func Reduce[A any, B any](reducer func(A, B) B, start B) func([]A) B {
	return func(input []A) B {
		reduced := start
		for _, v := range input {
			reduced = reducer(v, reduced)
		}
		return reduced
	}
}

func Chain2[A any, B any, C any](fn1 func(A) B, fn2 func(B) C) func(A) C {
	return func(input A) C {
		return fn2(fn1(input))
	}
}
