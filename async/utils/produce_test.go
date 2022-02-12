package utils_test

import (
	"funk/async/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduce(t *testing.T) {
	output := utils.Produce(func(output chan<- int) {
		output <- 5
		output <- 6
	})

	assert.Equal(t, 5, <-output)
	assert.Equal(t, 6, <-output)

	select {
	case _, ok := <-output:
		assert.False(t, ok)
	}
}
