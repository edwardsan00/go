package main

import "fmt"

func imprimirNombre(nombre string) { //Sin retorno
	fmt.Println("El nombre es:", nombre)
}

func suma(n1 int, n2 int) int { //Con retorno
	return n1 + n2
}
func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

func main() {
	imprimirNombre("Edward")
	fmt.Println(suma(10, 10))
	fmt.Println(resta(20, 10))
}
