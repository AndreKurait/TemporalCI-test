package main

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns the larger of a or b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of a or b.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Fibonacci returns the nth Fibonacci number.
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// GCD returns the greatest common divisor of a and b.
func GCD(a, b int) int {
	a, b = Abs(a), Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// IsPrime returns true if n is a prime number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
// triggered 2026-03-20T04:37:03Z
// verified
