package utils

const CHANNEL_BUFFER_SIZE = 10

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
