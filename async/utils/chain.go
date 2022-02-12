package utils

import "context"

func Chain2[A, B, C any](
	fn1 func(ctx context.Context, input <-chan A) <-chan B,
	fn2 func(ctx context.Context, input <-chan B) <-chan C,
) func(ctx context.Context, input <-chan A) <-chan C {
	return func(ctx context.Context, input <-chan A) <-chan C {
		return fn2(ctx, fn1(ctx, input))
	}
}
