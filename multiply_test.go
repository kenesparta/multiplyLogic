package multiplyLogic

import (
	"testing"
)

// TestMultiply calls TestMultiply with a two float numbers.
func TestMultiply(t *testing.T) {
	ans := Multiply(6.0, 7.0)
	if ans != 42.0 {
		t.Errorf("The result should be %f", 42.0)
	}
}
