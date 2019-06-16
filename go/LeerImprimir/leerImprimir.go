package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Leer e imprimir datos con fmt
	var edad int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Tu edad es %d\n", edad)

	//Leer e imprimir con bufio y os
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Coloca tu nombre: ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}
}
