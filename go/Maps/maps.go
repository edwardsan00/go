package main

import "fmt"

func main() {
	x := make(map[string]string, 2)
	x["nombre"] = "Alejandro"
	x["Edad"] = "20"
	fmt.Println(x["nombre"])
	fmt.Println(x["Edad"])
	dias := map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miercoles",
	}
	fmt.Println(dias)
	mes := map[string]int{
		"Enero":   01,
		"Febrero": 02,
		"Marzo":   03,
	}
	fmt.Println(mes["Enero"])

	//Borrar Maps
	delete(dias, 1)
	fmt.Println(dias)

	if mes, ok := mes["Jueves"]; ok { //ok es una variable que estoy reciviendo con un parametro boleano y lo evaluo
		if mes == 02 {
			fmt.Println("No es enero")
		} else {
			fmt.Println("Es Enero")
		}
	} else {
		fmt.Println("No es un mes")
	}
	fmt.Println(len(mes))

	//Recorriendo el map con range
	for mes, fecha := range mes {
		fmt.Printf("El mes es %s y numero %d \n", mes, fecha)
	}
}
