package util_test

import (
	"funk/sync/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTap(t *testing.T) {
	assert := assert.New((t))

	assert.Equal(0, util.Tap(func(a int) {})(0))
}
