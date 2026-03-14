package main

import "fmt"

func main() {
	fmt.Println(Greet("World"))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
// CI integration test Sat Mar 14 13:47:22 UTC 2026
// PR CI test Sat Mar 14 13:47:58 UTC 2026
