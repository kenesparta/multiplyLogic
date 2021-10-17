package multiplyLogic

import "math"

const accuracy = 1000

// Factor Main struct to get the data from the user
type Factor struct {
	First  *float64 `json:"first_factor"`
	Second *float64 `json:"second_factor"`
}

// Product Shows the final result using the accuracy
func (f *Factor) Product() float64 {
	return math.Round((*f.First)*(*f.Second)*accuracy) / accuracy
}

// AreValidNumbers Verifies if the two values are set in the JSON request
// If one of these values is not set, the result is false
func (f *Factor) AreValidNumbers() bool {
	return f.First != nil && f.Second != nil
}

// IsProductInfinite Validate if the product exceeds the bounds of float64
func (f *Factor) IsProductInfinite() bool {
	return f.Product() == math.Inf(0) || f.Product() == math.Inf(-1)
}
