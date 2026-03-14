package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
		{20, 6765},
	}
	for _, tt := range tests {
		if got := Fibonacci(tt.n); got != tt.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
