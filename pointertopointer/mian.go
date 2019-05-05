package main

import "fmt"

func main() {
	i := 10
	s := &i //指针
	fmt.Println(s)
	ss := &s //指针的指针
	fmt.Println(**ss)
	l := 15
	s = &l
	fmt.Println(**ss)
	//=================================
	s5 := new(rune)
	fmt.Println(&s5)
	change(s5)
	fmt.Println(string(*s5))
}

func change(a *rune) {
	fmt.Println(&a)
	*a = 65
}
