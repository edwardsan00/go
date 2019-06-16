package main

import "fmt"

func f1() {
	fmt.Println("Entrando a  f1")
	f2()
	fmt.Println("Saliendo de f1")
}
func f2() {
	fmt.Println("Entrando a  f2")
	fmt.Println("Saliendo de f2")
}
func main() {
	fmt.Println("Entrando a Main")
	f1()
	fmt.Println("Saliendo de Main")
}
