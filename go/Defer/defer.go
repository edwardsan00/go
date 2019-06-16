package main

import (
	"fmt"
	"os" //Permite ejecutar operaciones del sistema operativo
)

func primera() {
	fmt.Println("Primera")
}
func segunda() {
	fmt.Println("Segunda")
}
func main() {
	//defer primera() // Se ejecuta al terminar todo el ambito de la funcion aqui se ejecutar la funcion primera() luego de segunda()
	primera()
	segunda()
	//Abrimos el archivo
	f, err := os.Open("archivo.txt")
	//Verificamos errores con la funcion open
	if err != nil {
		panic(err)
	}
	defer f.Close() // Se cierra el archivo al final de leer todo el codigo
	//Alamacenamos el texto en un slice
	data := make([]byte, 50)
	//Leemos el slice(que contiene el texto) con la funcion read()
	c, err := f.Read(data)
	//Verificamos errores con la funcion read
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cantidad de bytes leidos: %d\nTexto leido: %s\nError: %v", c, data, err)
	//f.Close()
}
