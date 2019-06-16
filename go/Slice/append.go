package main

import "fmt"

func main() {
	//Funcion append
	p := make([]byte, 4, 10)       //Creo el slice
	p = []byte{'H', 'O', 'L', 'A'} // Asigno valores al slice
	fmt.Println("Slice P: ", p)
	fmt.Printf("Slice p: %q \n", p) //%q lee en utf8
	for i := 0; i < len(p); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, p[i])
	}
	p = append(p, ' ')  //Agrego otro valor al slice con append(slice, Contenido a agregar)
	fmt.Println(cap(p)) // El slice analiza el tamaño y si no es suficiente crea un array subyacente y crea otro slice
	p = append(p, 'M', 'U', 'N', 'D', 'O')
	fmt.Printf("Slice: %q \n", p)
	fmt.Println(cap(p))

	header := `********************
    Programa Append
    ********************`
	var y []int
	fmt.Println(header)
	for i := 1; i < 15; i++ {
		y = append(y, i)
		fmt.Println("Slice Y: ", y) //El slice duplica su tamaño cuando necesita mas espacio
		fmt.Printf("Longitud: %d Capadicad %d Elementos %d \n", len(y), cap(y), i)
	}
}
