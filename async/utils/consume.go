package utils

import "context"

// Consume is a helper for consuming from an input channel while respecting a context.
// This method will consume from `input` and pass its contents to `consumer` to do some
// sort of processing.
// Consume will stop passing data to `consumer` when `ctx.Done` has an event or the
// `input` channel gets closed
func Consume[A any](consumer func(value A)) func(ctx context.Context, input <-chan A) {
	return func(ctx context.Context, input <-chan A) {
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-input:
				if !ok {
					return
				}

				consumer(v)
			}
		}
	}
}
