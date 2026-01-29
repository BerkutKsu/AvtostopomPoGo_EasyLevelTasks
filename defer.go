package main

import "fmt"

// Что выведет код и почему?
//// start    end    3    2    1, потому что для вывода чисел в цикле мы использовали ключевое слово defer и все 3 вызова функции Println сложились в специальный стэк, который
//// выполнился в конце функции main в обратном порядке, что характерно для deferred функций
func main() {
	fmt.Println("start")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}
