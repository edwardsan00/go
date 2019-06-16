package main

import "fmt"

type Persona struct { //Estructura Persona
	Nombre string
	Edad   int
}
type Estudiante struct { //Estructura Estudiante
	Persona //Hereda de persona todos sus datos
	Carrera string
}
type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	jose := Estudiante{
		Persona{ //Asigno valores a persona
			Nombre: "Edward",
			Edad:   24,
		}, //Asigno valores a Estudiante
		"Informatica",
	}
	fmt.Println(jose)
	juan := Profesor{
		Estudiante{
			Persona{
				Nombre: "Juan",
				Edad:   25,
			},
			"Ing Civil",
		},
		"Informatica",
	}
	fmt.Println(juan)
	//Acceder a datos en las Estructuras
	fmt.Printf("Nombre: %s, Edad: %d, Carrera %s, Carrera Prof. %s\n", juan.Nombre, juan.Edad, juan.Estudiante.Carrera, juan.Carrera)

	//Otro metodo de nueva instancia
	var pedro Profesor
	pedro.Nombre = "Pedro"
	pedro.Edad = 30
	pedro.Estudiante.Carrera = "Produccion"
	pedro.Carrera = "Contabilidad"
	fmt.Println(pedro)
}
