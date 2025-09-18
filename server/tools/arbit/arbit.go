package arbit

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func GenerateRandDigits(count int) string {
	sliceRes := make([]string, 0, count)

	for i := 0; i < count; i++ {
		r := rand.IntN(10)
		sliceRes = append(sliceRes, fmt.Sprintf("%d", r))
	}

	return strings.Join(sliceRes, "")
}
