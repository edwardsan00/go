package main

import "fmt"

type Persona struct {
	nombre string
	email  string
	edad   int
}
type Moderador struct {
	Persona
	Foro string
}
type Administrador struct {
	Persona
	Seccion string
}
type Usuario interface { //En interface se colocan los metodos que va a ejecutar
	Nombre() string // Todas las estructuras que tengan metodo nombre y devuelva un string pasa a ser Usuario
	Email() string  // Todas las estructuras que tengan metodo email y devuelva un string pasa a ser Usuario
}

func (m Moderador) AbrirForo() {
	fmt.Println("Abrir foro")
}
func (a Administrador) AbrirGrupo() {
	fmt.Println("Abrir grupo")
}
func (p Persona) Nombre() string { //Metodo que me retorna el nombre
	return p.nombre
}
func (p Persona) Email() string { //Metodo que me retorna el email
	return p.email
}
func Presentarse(p Usuario) { //Funcion que recibe como parametro usuario (que tiene nombre & email) y accede a los metodos d persona
	fmt.Println("\n", p.Nombre())
	fmt.Println("\n", p.Email())
}
func main() {
	alejandro := Persona{"Alejandro", "email", 20} //Nueva persona
	pedro := Moderador{Persona{"Pedro", "email2", 26}, "Juego"}
	juan := Administrador{Persona{"Juan", "email3", 30}, "Pc"}
	Presentarse(alejandro) //Paso nueva persona como parametro
	Presentarse(pedro)
	Presentarse(juan)

	var i Usuario //interfaces son tipos asi que podemos definir variables de ese tipo
	i = alejandro //Alejandro de tipo persona aplica a interface usuario, serian del mismo tipo
	fmt.Println("I alejandro", i)
	fmt.Println("Nombre de i", i.Nombre())
	fmt.Println("Email de i", i.Email())
}
