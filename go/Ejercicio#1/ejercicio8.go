package main

import "fmt"

func main() {
	var salto []int
	var salto2 []int
	var base, intervalo, base2, intervalo2, limite int
	limite = 28
	fmt.Scan(&base, &intervalo, &base2, &intervalo2)
	salto = jump(base, intervalo, limite)
	salto2 = jump(base2, intervalo2, limite)
	fmt.Println(salto)
	fmt.Println(salto2)
}
func jump(base, intervalo, limite int) []int {
	var jump []int
	var acum int
	jump = append(jump, base)
	for i := 0; i < (limite / intervalo); i++ {
		if i < 1 {
			acum += base + intervalo
		} else {
			acum += intervalo
		}
		jump = append(jump, acum)
	}
	return jump
}
