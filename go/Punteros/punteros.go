package main

import "fmt"

func main() {
	a := 20
	b := &a //& Accede a la direccion en memoria de a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // * accede al valor en memoria de a
	// *int es diferente a int

}
