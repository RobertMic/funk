package utils_test

import (
	"context"
	"funk/async/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsume(t *testing.T) {
	input := make(chan int, 5)
	for i := 0; i < 5; i++ {
		input <- 5
	}
	close(input)

	assert := assert.New(t)
	got := 0
	utils.Consume(func(a int) {
		assert.Equal(5, a)
		got++
	})(context.Background(), input)

	assert.Equal(5, got)
}
