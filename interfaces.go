package main

import "fmt"

// Что выведет код и почему?
//// Выведет false, потому что A() вернет интерфейсное значение nil с типом nil, а B() вернет nil с типом *main.impl
type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}
