package main

import "fmt"

func imprimir() {
	fmt.Println("Hola")

	defer func() {
		cadena := recover() // Recupera el error que manda el panic()
		fmt.Println(cadena)
	}()
	panic("Error") // panic dispara un error en tiempo de ejecucion
}
func main() {
	imprimir()
	fmt.Print("Hola 2")
}
