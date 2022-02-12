package main

import (
	"context"
	"fmt"
	"funk/async/channels"
	"funk/async/utils"
	autils "funk/async/utils"
)

func main() {
	input := make(chan int)

	chain := autils.Chain2(
		channels.Filter(func(value int) bool {
			return value%2 == 0
		}),
		channels.Map(func(value int) string {
			return fmt.Sprint(value)
		}),
	)

	utils.CollectChain(chain)(context.Background(), input)
}
