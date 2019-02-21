package randx

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomSleep(i int) {
	// sleep for a random amount of time to add "chaos"
	time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
}
