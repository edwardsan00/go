package main

import "fmt"

func main() {
	origen := []int{1, 2, 3}
	destino := []int{4, 5, 6}
	copy(destino, origen) // Primero el destino luego el origen
	fmt.Println(origen, destino)
	//Origen > destino
	origen2 := []int{1, 2, 3}
	destino2 := make([]int, 2)
	copy(destino2, origen2) // Solo copia los espacios disponibles en destino
	fmt.Println(origen2, destino2)

	//Destino > orgien
	origen3 := []int{1, 2}
	destino3 := []int{3, 4, 5}
	copy(destino3, origen3) //Copia los disponibles
	fmt.Println(origen3, destino3)
}
