package main

import "strings"

// Anagram checks if two strings are anagrams of each other.
func Anagram(a, b string) bool {
	a, b = strings.ToLower(a), strings.ToLower(b)
	if len(a) != len(b) {
		return false
	}
	counts := make(map[rune]int)
	for _, c := range a {
		counts[c]++
	}
	for _, c := range b {
		counts[c]--
		if counts[c] < 0 {
			return false
		}
	}
	return true
}
