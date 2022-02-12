package utils

import "context"

func Collect[A any](ctx context.Context, input <-chan A) []A {
	var output []A
	Consume(func(value A) {
		output = append(output, value)
	})(ctx, input)
	return output
}

func CollectChain[A, B any](
	chain func(ctx context.Context, input <-chan A) <-chan B,
) func(ctx context.Context, input <-chan A) []B {
	return func(ctx context.Context, input <-chan A) []B {
		return Collect(ctx, chain(ctx, input))
	}
}
