package multiplyLogic

import (
	"fmt"
	"math"
	"os"
	"testing"
)

const coverageAllowed = 0.5

// TestMain Verifies if the coverage percentage is equals to 'coverageAllowed'
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
	tests := []struct {
		input struct {
			first  float64
			second float64
		}
		output float64
	}{
		{
			input: struct {
				first  float64
				second float64
			}{
				first:  0.0,
				second: 0.0,
			},
			output: 0.0,
		},
		{
			input: struct {
				first  float64
				second float64
			}{
				first:  -1.123,
				second: -125.879,
			},
			output: 141.362,
		},
		{
			input: struct {
				first  float64
				second float64
			}{
				first:  102.003,
				second: -12.358,
			},
			output: -1260.553,
		},
		{
			input: struct {
				first  float64
				second float64
			}{
				first:  math.MaxFloat64,
				second: 0,
			},
			output: 0,
		},
		{
			input: struct {
				first  float64
				second float64
			}{
				first:  math.MaxFloat64,
				second: -1257.127,
			},
			output: math.Inf(-1),
		},
		{
			input: struct {
				first  float64
				second float64
			}{
				first:  93651428.108,
				second: 12577425.127,
			},
			output: 1177893825064993.250,
		},
	}

	for _, tt := range tests {
		m := Multiply(tt.input.first, tt.input.second)
		if Multiply(tt.input.first, tt.input.second) != tt.output {
			t.Errorf("%.3f is not equal to %.3f", tt.output, m)
		}
	}
}
