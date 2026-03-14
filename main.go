package main

import "fmt"

func main() {
	fmt.Println(Greet("World"))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
// Trigger CI test Sat Mar 14 13:46:06 UTC 2026
