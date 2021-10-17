package multiplyLogic

import "math"

const factor = 1000

// Multiply The result is a product of two decimal numbers
func Multiply(p, s float64) float64 {
	return math.Round(p*s*factor) / factor
}
