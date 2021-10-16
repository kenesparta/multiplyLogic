package multiplyLogic

import (
	"fmt"
	"os"
	"testing"
)

const coverageAllowed = 0.5

func TestMain(m *testing.M) {
	rc := m.Run()

	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c <= coverageAllowed {
			fmt.Printf("Allowed: %.2f / Current: %.2f\n", coverageAllowed, c)
			rc = -1
		}
	}
	os.Exit(rc)
}

// TestMultiply calls TestMultiply with a two float numbers.
func TestMultiply(t *testing.T) {
	ans := Multiply(6.0, 7.0)
	if ans != 42.0 {
		t.Errorf("The result should be %f", 42.0)
	}
}
