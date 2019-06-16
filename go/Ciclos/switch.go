package main

import "fmt"

func main() {
	numero := 30
	switch {
	case numero > 23:
		fmt.Println("Es mayor que 23")
		fallthrough // Continua evaluando
	case numero > 25:
		fmt.Println("Es mayor que 25")
	default:
		fmt.Println("Es un numero")
	}
}
