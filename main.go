package main

import (
	"fmt"
)

// CONSTANTS
var n int = 0

func f(x float64) float64 {
	return (x - (((x * x) - (3 * x) - 4) / ((2 * x) - 3)))
}

func main() {
	newton_rafson(1.75, f)
}

// func tan_f(x float64) float64 {
// 	x = x - ((math.Tan(x) - x - 1) / (math.Tan(x) * math.Tan(x)))
// 	return x
// }

func newton_rafson(x float64, f func(float64) float64) {
	n++
	x = f(x)
	fmt.Printf("x_%d = %f\n", n, x)
	if n < 300 {
		newton_rafson(x, f)
	} else {
		return
	}
}
