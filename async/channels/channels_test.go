package channels_test

import (
	"context"
	"funk/async/channels"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	input := make(chan int)
	output := channels.Map(func(value int) int {
		return value + 1
	})(context.Background(), input)

	input <- 5
	assert.Equal(t, 6, <-output)

	input <- 6
	assert.Equal(t, 7, <-output)

	close(input)
	select {
	case _, ok := <-output:
		assert.False(t, ok)
	}
}

func TestFilter(t *testing.T) {
	input := make(chan int)
	output := channels.Filter(func(value int) bool {
		return value > 5
	})(context.Background(), input)

	input <- 5
	input <- 5
	input <- 5
	input <- 6
	assert.Equal(t, 6, <-output)

	close(input)
	select {
	case _, ok := <-output:
		assert.False(t, ok)
	}
}

func TestReduce(t *testing.T) {
	input := make(chan int)
	output := channels.Reduce(func(value int, aggregate int) int {
		return value + aggregate
	}, 0)(context.Background(), input)

	for i := 0; i < 5; i++ {
		input <- 1
	}

	close(input)
	assert.Equal(t, 5, <-output)

	select {
	case _, ok := <-output:
		assert.False(t, ok)
	}
}
