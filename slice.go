package main

import "fmt"

// 1
// Какой будет результат выполнения приложения?
//// [a q c]
func main() {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"
	fmt.Println(a)
}

// 2
// Что выведет код и почему?
//// [5 5 5 5 5 5] и [1 2 3 4 5], потому что после того, как мы добавили в срез a еще один элемент, срез a был пересоздан и как следствие а и sl перестали ссылаться на один базовый массив
func mod(a []int) {
	a = append(a, 125)
	for i := range a {
		a[i] = 5
	}
	fmt.Println(a)
}
func main() {
	sl := []int{1, 2, 3, 4, 5}
	mod(sl)
	fmt.Println(sl)
}

// 3
// Что выведет код и почему?
//// [5 5 5 5] и [5 5 5 5], потому что a и sl ссылаются на один базовый массив
func mod(a []int) {
	for i := range a {
		a[i] = 5
	}
	fmt.Println(a)
}
func main() {
	sl := make([]int, 4, 8)
	sl[0] = 1
	sl[1] = 2
	sl[2] = 3
	sl[3] = 5
	mod(sl)
	fmt.Println(sl)
}
