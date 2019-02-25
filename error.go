package randx

import (
	"fmt"
	"math/rand"
)

func RandomError(n int) error {
	x := rand.Int()
	if x%n == 0 {
		return fmt.Errorf("%d is divisable by %d", x, n)
	}
	return nil
}

// section: do
