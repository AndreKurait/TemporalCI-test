package main

// Clamp restricts a value to a range [min, max].
func Clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
// dedup verification Mon Mar 16 10:38:48 UTC 2026
