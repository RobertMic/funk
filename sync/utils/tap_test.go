package utils_test

import (
	"funk/sync/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTap(t *testing.T) {
	assert := assert.New((t))

	assert.Equal(0, utils.Tap(func(a int) {})(0))
}
