package set_rand

import (
	"math/rand"
	"time"
)

func Rand_bot_id() string{
	reasons := make([]string, 0)
	reasons = append(reasons, "bot03","bot02","bot01")
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return reasons[rand.Intn(len(reasons))]
}
