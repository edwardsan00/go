package main

import "fmt"

type Dinero int //Declaramos un nuevo tipo con type seguido del nombre y el tipo de origan

func (d Dinero) String() string { //Metodo.. declaramos antes del nombre la variable(d) y a la cual hace referencia(Dinero)
	return fmt.Sprintf("$%d", d)
}

func main() {
	var sueldo Dinero
	sueldo += 25000
	fmt.Println("Sueldo", sueldo)

	aumento := 10000
	sueldo += Dinero(aumento)     //Convirtiendo aumento de tipo int a tipo Dinero(dinero int)
	fmt.Println("Sueldo", sueldo) //Imrpime sueldo + $
}
