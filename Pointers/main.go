package main

import "fmt"

func main() {
	a := make(map[int]string)
	a[1] = "Hello"
	a[2] = "World"
	a[3] = "!"
	fmt.Println(a)
	pointerExample1(a)
	fmt.Println(a)
	pointerExample2(&a)
	fmt.Println(a)
}

//map[KeyType]ValueType
func pointerExample1(a map[int]string) {
	a[1] = "Hell"
}

//map[KeyType]ValueType
func pointerExample2(a *map[int]string) {
	(*a)[1] = "Hells"
}
