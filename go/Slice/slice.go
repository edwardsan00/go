package main

import "fmt"

func main() {
	//Slice son partes de un array
	var x []int
	fmt.Println("Esto es un slice ", x)

	j := []int{1, 2, 4}
	fmt.Println(j)

	y := make([]int, 5, 10) //Make define el slice pro tipo longitud y capacidad
	fmt.Println("slice: ", y)
	fmt.Println("Longitud: ", len(y))
	fmt.Println("Capacidad: ", cap(y))

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // Declaramos el array
	fmt.Println("Array", arr)
	a := arr[2:5] //Definimos el slice y su tamaño
	fmt.Println("Slice: ", a)
	b := arr[3:5]
	fmt.Println("Slice: ", b)
	b[0] = 20
	fmt.Println(arr) // Modifica el espacio en memoria desde el slice

}
