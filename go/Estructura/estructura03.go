package main

import (
	"fmt"
	"math"
)

type Rectangulo struct {
	alto, ancho float64
}
type Circulo struct {
	radio float64
}

func (r Rectangulo) area() float64 { //Area pasa a ser metodo de Rectangulo, esta dentro de rectangulo
	return r.alto * r.ancho
}
func (c Circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}
func (r *Rectangulo) duplicar(i float64) { //Puntero al rectangulo y que itere en su mismo rectangulo y no en uno nuevo
	r.alto *= i
	r.ancho *= i
}

func main() {
	r1 := Rectangulo{10, 20}
	fmt.Println(r1.area())

	c1 := Circulo{120}
	fmt.Println(c1.area())

	r2 := Rectangulo{40, 50} //Creo el objeto rectangulo
	r2.duplicar(10)          //Paso el parametro y reescribe el valor
	fmt.Println(r2)
}
