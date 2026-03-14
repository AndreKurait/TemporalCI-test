package main

func BrokenFunction() int {
	return "this is not an int" // type mismatch - won't compile
}
