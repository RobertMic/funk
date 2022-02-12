package slices_test

import (
	"fmt"
	"funk/sync/slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	input := []int{0, 1, 2, 3}
	output := slices.Map(func(a int) string { return fmt.Sprint(a) })(input)

	assert := assert.New(t)

	assert.Equal([]string{"0", "1", "2", "3"}, output)
}

func TestMapNested(t *testing.T) {
	input := []int{0, 1, 2, 3}
	output := slices.Map(func(a int) []int { return []int{a} })(input)

	assert := assert.New(t)

	assert.Equal([][]int{{0}, {1}, {2}, {3}}, output)
}

func TestReduce(t *testing.T) {
	input := []int{0, 1, 2, 3}
	output := slices.Reduce(func(a int, b int) int { return a + b }, 0)(input)

	assert := assert.New(t)

	assert.Equal(6, output)
}
