package randx

import (
	"fmt"
	"math/rand"
)

func RandomError() error {
	x := rand.Int()
	if x%3 == 0 {
		return fmt.Errorf("%d is divisable by 3", x)
	}
	return nil
}

// section: do
