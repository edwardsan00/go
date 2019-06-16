package main

import "fmt"

type Persona struct { //Declaramos el struct que seria como una clase
	Nombre string
	Edad   int
}

func orden(p1, p2 Persona) (Persona, int) {
	if p1.Edad > p2.Edad {
		return p1, p1.Edad - p2.Edad
	}
	return p2, p2.Edad - p1.Edad
}

func main() {
	//Metodos para acceder a la Estructura Persona
	var p Persona       // P sera instancia de persona
	p.Nombre = "Edward" //accedemos a la variable Nombre de la clase Persona
	p.Edad = 24
	fmt.Println("Persona", p)
	fmt.Println("Nombre:", p.Nombre)

	//Declarando otra Persona con nombre de propiedades
	p2 := Persona{Nombre: "Edward", Edad: 24}
	fmt.Println("Edad", p2.Edad)

	//Declaranto nueva persona sin nombre de propiedades
	p3 := Persona{"Edward", 24}
	fmt.Println("Persona 3", p3)

	//Ejercicio
	tom := Persona{"Tom", 60} //Declaro 3 nuevas personas
	juan := Persona{"Juan", 20}
	pedro := Persona{"Pedro", 36}
	//Ejecuto la funcion Orden
	pMayor1, dif1 := orden(tom, juan) //Almaceno los retornos en las variables declaradas
	pMayor2, dif2 := orden(juan, pedro)
	pMayor3, dif3 := orden(tom, pedro)

	//Imprimo resultado
	fmt.Printf("Entre %s y %s el mayor es %s por %d años\n", tom.Nombre, juan.Nombre, pMayor1.Nombre, dif1)
	fmt.Printf("Entre %s y %s el mayor es %s por %d años\n", juan.Nombre, pedro.Nombre, pMayor2.Nombre, dif2)
	fmt.Printf("Entre %s y %s el mayor es %s por %d años\n", tom.Nombre, pedro.Nombre, pMayor3.Nombre, dif3)
}
