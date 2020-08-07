package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	s1, s2 := "hello", "world"

	fmt.Println(swap(s2, s1))
}
