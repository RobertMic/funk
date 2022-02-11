package slices_test

import (
	"fmt"
	"funk/slices"
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

func TestChain2(t *testing.T) {
	input := []int{0, 1, 2, 3}
	output := slices.Chain2(
		slices.Map(func(a int) int { return a + 1 }),
		slices.Reduce(func(a int, b int) int { return a + b }, 0),
	)(input)

	assert.Equal(t, 10, output)
}

func TestChain2NestedChain(t *testing.T) {
	input := []int{0, 1, 2, 3}
	firstChain := slices.Chain2(
		slices.Map(func(a int) int { return a + 1 }),
		slices.Map(func(a int) int { return a * 0 }),
	)
	secondChain := slices.Chain2(
		firstChain,
		slices.Map(func(a int) string { return fmt.Sprint(a) }),
	)
	output := slices.Chain2(
		secondChain,
		slices.Reduce(func(a string, b map[string]struct{}) map[string]struct{} { b[a] = struct{}{}; return b }, map[string]struct{}{}),
	)(input)

	assert.Equal(t, map[string]struct{}{"0": {}}, output)
}
