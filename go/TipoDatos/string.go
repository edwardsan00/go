package main

import "fmt"

func main() {
	var cadena string
	cadena = "Hola"
	fmt.Println(cadena)
	fmt.Println(len(cadena)) // len Cuenta el tamaño de la cadena
	fmt.Println(cadena[2])   //Se puede acceder por indice
	fmt.Println(cadena[0:2]) //Accediendo a slice de la cadena

}
