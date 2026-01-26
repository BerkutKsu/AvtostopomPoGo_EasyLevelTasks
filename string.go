package main

import "fmt"

// Расскажи подробно что происходит
// Как сделать так, чтобы работало?

func main() {
	str := "Привет"
	//str[2] = 'e' - неправильно, теперь как правильно:
	word := []rune(str)
	word[2] = 'е'
	str = string(word)
	fmt.Println(str)
}
