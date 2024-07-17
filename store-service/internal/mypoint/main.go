package mypoint

import (
	"fmt"
	"math"
)

func Calculate_num(a, b int) (int, error) {
	if a < 0 {
		return math.MinInt, fmt.Errorf("positive int only, now a =  %d", a)

	} else if b < 0 {
		return math.MinInt, fmt.Errorf("positive int only, now a =  %d", b)
	}
	return a + b, nil
}
