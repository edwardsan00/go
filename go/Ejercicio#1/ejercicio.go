package main

import "fmt"

func main() {
	var numero1, numero2, contador int
	fmt.Println("Digite el primer numero: ")
	fmt.Scanf("%d\n", &numero1)
	fmt.Println("Digite el segundo numero: ")
	fmt.Scanf("%d\n", &numero2)
	for i := numero1; i <= numero2; i++ {
		if i%2 != 0 {
			contador++
			fmt.Printf("%d Es un numero inpar\n", i)
		}
	}
	fmt.Printf("Entre el %d y el %d hay %d  numeros impares\n", numero1, numero2, contador)
}
