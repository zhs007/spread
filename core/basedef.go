package spreadcore

import (
	"math/rand"
	"time"
)

const RATE_VALUE = 10000

func init() {
	rand.Seed(time.Now().Unix())
}
