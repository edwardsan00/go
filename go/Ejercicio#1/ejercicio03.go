package main

import (
	"fmt"
	"reflect"
)

var a string

func main() {
	a, b := "Hola", 3
	_ = a
	fmt.Println(b)
	mostrar()
	fmt.Println(reflect.TypeOf(a))
}
func mostrar() {
	fmt.Println(a)
}
