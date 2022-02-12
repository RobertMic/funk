package utils_test

import (
	"fmt"
	"testing"

	"funk/sync/slices"
	"funk/sync/utils"

	"github.com/stretchr/testify/assert"
)

func TestChain2(t *testing.T) {
	input := []int{0, 1, 2, 3}
	output := utils.Chain2(
		slices.Map(func(a int) int { return a + 1 }),
		slices.Reduce(func(a int, b int) int { return a + b }, 0),
	)(input)

	assert.Equal(t, 10, output)
}

func TestChain2_2(t *testing.T) {
	input := []int{0, 1, 2, 3}
	chain := utils.Chain2(
		slices.Reduce(func(a int, b int) int { return a + b }, 0),
		func(b int) string { return fmt.Sprint(b) },
	)

	assert.Equal(t, "6", chain(input))
}

func TestChain2NestedChain(t *testing.T) {
	input := []int{0, 1, 2, 3}
	firstChain := utils.Chain2(
		slices.Map(func(a int) int { return a + 1 }),
		slices.Map(func(a int) int { return a * 0 }),
	)
	secondChain := utils.Chain2(
		firstChain,
		slices.Map(func(a int) string { return fmt.Sprint(a) }),
	)
	output := utils.Chain2(
		secondChain,
		slices.Reduce(func(a string, b map[string]struct{}) map[string]struct{} { b[a] = struct{}{}; return b }, map[string]struct{}{}),
	)(input)

	assert.Equal(t, map[string]struct{}{"0": {}}, output)
}
