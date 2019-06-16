package main

import "fmt"

func main() {
	nombres := []string{
		"Edward",
		"Jesus",
		"Sanchez",
	}
	for index, nombre := range nombres {
		fmt.Printf("El nombre es %q esta en el index %d \n", nombre, index)
	}
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
	for index := range nombres {
		fmt.Println(index)
	}
	cadena := "Hola"
	for index, letra := range cadena {
		fmt.Printf("Letra %q y index %d \n", letra, index)
	}


 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var total int
    arr := [6]int{1,2,3,4,10,11}
   for _, valor := range arr{
            total = total + valor
        }
    fmt.Println(total)

}
