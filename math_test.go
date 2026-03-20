package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct{ input, want int }{
		{5, 5}, {-5, 5}, {0, 0}, {-1, 1},
	}
	for _, tt := range tests {
		if got := Abs(tt.input); got != tt.want {
			t.Errorf("Abs(%d) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{1, 2, 2}, {5, 3, 5}, {-1, -2, -1}, {0, 0, 0},
	}
	for _, tt := range tests {
		if got := Max(tt.a, tt.b); got != tt.want {
			t.Errorf("Max(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{1, 2, 1}, {5, 3, 3}, {-1, -2, -2}, {0, 0, 0},
	}
	for _, tt := range tests {
		if got := Min(tt.a, tt.b); got != tt.want {
			t.Errorf("Min(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 0}, {1, 1}, {2, 1}, {5, 5}, {10, 55}, {-1, 0},
	}
	for _, tt := range tests {
		if got := Fibonacci(tt.n); got != tt.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{12, 8, 4}, {7, 13, 1}, {0, 5, 5}, {100, 75, 25}, {-12, 8, 4},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 97}
	for _, n := range primes {
		if !IsPrime(n) {
			t.Errorf("IsPrime(%d) = false, want true", n)
		}
	}
	notPrimes := []int{0, 1, 4, 9, 15, 100}
	for _, n := range notPrimes {
		if IsPrime(n) {
			t.Errorf("IsPrime(%d) = true, want false", n)
		}
	}
}
