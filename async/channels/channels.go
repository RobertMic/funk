package channels

import (
	"context"
	"funk/async/utils"
)

// Map consumes from an `input` channel and passes the values to the `mapper`.
// The return of `mapper` then gets produced to an output channel.
//
// Map runs in its own go routine.
func Map[A, B any](mapper func(value A) B) func(ctx context.Context, input <-chan A) <-chan B {
	return func(ctx context.Context, input <-chan A) <-chan B {
		return utils.Produce(func(output chan<- B) {
			utils.Consume(func(v A) {
				output <- mapper(v)
			})(ctx, input)
		})
	}
}

// Filter consumes from an `input` channel and filters out any messages that the
// `filter` returns `false` for. Any messages that `filter` returns `true` are
// forwarded to the output channel
func Filter[A any](filter func(value A) bool) func(ctx context.Context, input <-chan A) <-chan A {
	return func(ctx context.Context, input <-chan A) <-chan A {
		return utils.Produce(func(output chan<- A) {
			utils.Consume(func(v A) {
				if filter(v) {
					output <- v
				}
			})(ctx, input)
		})
	}
}

// Reduce will consume from a channel, calling `reducer` on any values it receives.
// Once the `input` channel is closed, the final call to `reducer` will be sent to the
// output channel.
// The inital call to `reducer` will use `start` as the `aggregate`, subsequent calls
// will use the return of `reducer` as `aggregate`.
//
// Reduce runs in its own go routine.
func Reduce[A, B any](reducer func(value A, aggregate B) B, start B) func(ctx context.Context, input <-chan A) <-chan B {
	return func(ctx context.Context, input <-chan A) <-chan B {
		return utils.Produce(func(output chan<- B) {
			aggregate := start
			utils.Consume(func(a A) {
				aggregate = reducer(a, aggregate)
			})(ctx, input)
			output <- aggregate
		})
	}
}
