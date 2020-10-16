package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
)

type testCase struct {
	name string
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	testCases := []testCase{
		{name: "Two positive numbers", a: 2, b: 4, want: 6},
		{name: "Two negative numbers", a: -2, b: -4, want: -6},
		{name: "Positive and negative numbers", a: 2, b: -5, want: -3},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("%s %f and %f:, got %f want %f", tc.name, tc.a, tc.b, got, tc.want)
		}
	}
}

func TestAddRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		a, b := rand.Float64(), rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if got != want {
			t.Errorf("Adding %f and %f: got %f want %f", a, b, got, want)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []testCase{
		{name: "Two positive numbers with positive result", a: 4, b: 2, want: 2},
		{name: "Two positive numbers with negative result", a: 2, b: 4, want: -2},
		{name: "Two negative numbers with positive result", a: -2, b: -5, want: 3},
		{name: "Two negative numbers with negative result", a: -5, b: -3, want: -2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("%s %f and %f:, got %f want %f", tc.name, tc.a, tc.b, got, tc.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []testCase{
		{name: "Two positive numbers", a: 4, b: 2, want: 8},
		{name: "Two negative numbers", a: -2, b: -2, want: 4},
		{name: "Positive and negative number", a: 2, b: -2, want: -4},
		{name: "Zero number", a: 2, b: 0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("%s %f and %f:, got %f want %f", tc.name, tc.a, tc.b, got, tc.want)
		}
	}
}

type testCaseWithError struct {
	name      string
	a, b      float64
	want      float64
	expectErr bool
}

func TestDivide(t *testing.T) {
	testCases := []testCaseWithError{
		{name: "Normal case", a: 12, b: 3, want: 4, expectErr: false},
		{name: "Divide by zero", a: 12, b: 0, want: 0, expectErr: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if tc.expectErr != (err != nil) {
			t.Fatalf("%s %f and %f: unexpected err: %v", tc.name, tc.a, tc.b, err)
		}

		if !tc.expectErr && got != tc.want {
			t.Errorf("%s %f and %f: got %f want %f", tc.name, tc.a, tc.b, got, tc.want)
		}

	}
}

type testCaseSqrt struct {
	name      string
	a         float64
	want      float64
	expectErr bool
}

func TestSqrt(t *testing.T) {
	testCases := []testCaseSqrt{
		{name: "Round sqrt", a: 4, want: 2, expectErr: false},
		{name: "Floating result", a: 5, want: 2.2360679, expectErr: false},
		{name: "Invalid number", a: -4, want: -1, expectErr: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if tc.expectErr != (err != nil) {
			t.Fatalf("%s %f: unexpected error: %v", tc.name, tc.a, err)
		}

		if !tc.expectErr && !closeEnough(got, tc.want) {
			t.Errorf("%s %f: got %f want %f", tc.name, tc.a, got, tc.want)
		}
	}
}

func closeEnough(a, b float64) bool {
	return math.Abs(a-b) <= 0.00001
}
