package main

import "fmt"

func main() {
	var x [3]int //Array de 3 valores
	fmt.Println(x)

	var k [3][2]float64 // Array de 3 valores con 2 dentro
	fmt.Println(k)

	x[1] = 25 //Accedo y asigno al indice 1 de la variable x
	fmt.Println(x)
	fmt.Println(x[1]) //Accedo al indice 1

	y := [5]int{1, 2, 3, 4, 5} //Creo y asigno valores a los indices
	fmt.Println(y)

	j := [...]int{1, 2, 3, 4} //Omito la cantidad pero go la completa
	fmt.Println(j)

	i := [...]int{
		1,
		2,
		3,
		4,
	} //Asigando valores a los indices
	fmt.Println(i)

	f := [...]float64{1.2, 10.5, 8.3}
	fmt.Println(f)

	fmt.Println(len(f))     // Ver tamaño del array con len()
	fmt.Println(len(f) - 1) // Accediendo al ultimo elemento

	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	fmt.Println(a == b)
	fmt.Println(a == c)
}
