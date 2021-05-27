package main

import (
	"fmt"

	"github.com/fakovacic/test-actions/internal/tested"
)

func main() {
	fmt.Println(
		tested.Sum(2, 10),
	)
}
