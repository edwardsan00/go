package main

import "fmt"

//Retornar mas de un valor declarando el retorno arriba
func multiplicar(n1, n2, n3 int) (r1 int, r2 int, r3 int) {
	r1 = n1 * 1
	r2 = n2 * 2
	r3 = n3 * 3
	return
}

//Retorno declarando abajo
func sumar(numero int) (int, int) {
	r1 := numero + 10
	r2 := numero + 20
	return r1, r2
}

//Retornar para almacenar
func restar(numero int) (int, int) {
	r1 := numero - 10
	r2 := numero - 20
	return r1, r2
}
func main() {
	fmt.Println(multiplicar(2, 3, 4))
	fmt.Println(sumar(5))
	_, num2 := restar(100) //Recibo las variables y las almaceno, la primera la descarto
	fmt.Println(num2)
}
