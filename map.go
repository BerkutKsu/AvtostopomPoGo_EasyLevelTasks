package main

import (
	"fmt"
)

// 1
// Что происходит?
// Как сделать так, чтобы работало?
func main() {
	//var m map[string]int
	m := make(map[string]int) // не просто объявляем мапу, а инициализируем ее
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
