package channels

import "context"

const CHANNEL_BUFFER_SIZE = 10

// Map consumes from an `input` channel and passes the values to the `mapper`.
// The return of `mapper` then gets produced to an output channel.
//
// Map runs in its own go routine.
func Map[A, B any](mapper func(value A) B) func(ctx context.Context, input <-chan A) <-chan B {
	return func(ctx context.Context, input <-chan A) <-chan B {
		return Produce(func(output chan<- B) {
			Consume(func(v A) {
				output <- mapper(v)
			})
		})
	}
}

// Filter consumes from an `input` channel and filters out any messages that the
// `filter` returns `true` for. Any messages that `filter` returns `false` for are
// forwarded to the output channel
func Filter[A any](filter func(value A) bool) func(ctx context.Context, input <-chan A) <-chan A {
	return func(ctx context.Context, input <-chan A) <-chan A {
		return Produce(func(output chan<- A) {
			Consume(func(v A) {
				if !filter(v) {
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
		return Produce(func(output chan<- B) {
			aggregate := start
			Consume(func(a A) {
				aggregate = reducer(a, aggregate)
			})
			output <- aggregate
		})
	}
}

// Consume is a helper for consuming from an input channel while respecting a context.
// This method will consume from `input` and pass its contents to `consumer` to do some
// sort of processing.
// Consume will stop passing data to `consumer` when `ctx.Done` has an event or the
// `input` channel gets closed
func Consume[A any](consumer func(A)) func(ctx context.Context, input <-chan A) {
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

// Produce is a helper method that creates a channel with size `CHANNEL_BUFFER_SIZE`
// and spawns a go routine that cleans up the channel when it exits.
// The spawned go routine passes the created channel to `producer`.
// `producer` is expected to only exit when it wants to stop working,
// either because it has no more work or because it needs to stop.
func Produce[A any](producer func(output chan<- A)) <-chan A {
	output := make(chan A, CHANNEL_BUFFER_SIZE)

	go func() {
		defer func() {
			close(output)
		}()
		producer(output)
	}()

	return output
}
