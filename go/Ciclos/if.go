package main

import "fmt"

func main() {
	edad := 18
	condMaxima := 30
	condMinima := 20
	if edad+2 >= condMinima && edad <= condMaxima {
		fmt.Println("Puede Pasar")
	} else if edad != condMinima || edad != condMaxima {
		fmt.Println("Es Distinto")
	} else {
		fmt.Println("No puede pasar")
	}
}
