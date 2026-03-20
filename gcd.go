package main

// GCD returns the greatest common divisor of a and b using Euclid's algorithm.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}
