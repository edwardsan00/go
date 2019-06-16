package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Convertir a String valor Entero
	numero1 := 23
	//strconv - Libreria
	//strconv.Itoa(parametro) - Funcion Enteros a String
	edad_str := strconv.Itoa(numero1)
	fmt.Println("Mi edad es " + edad_str)

	//Convertir a Enteros valor String
	numero2 := "23"
	//strconv.Atoi(parametro) - Funcion String a Enteros
	// _ Recibe parametro pero la desecha(Hace referencia al valor Err devuelto por go)
	edad_int, _ := strconv.Atoi(numero2)
	fmt.Println(edad_int + 10)
}



package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Meal Cost: ")
    mealCost,_ := reader.ReadString('\n')
    fmt.Println("Tip Percent: ")
    tipPercent,_ := reader.ReadString('\n')
    fmt.Println("Tax Percent: ")
    taxPercent,_ := reader.ReadString('\n')
    //Calculations
    mealCost_int,_ := strconv.Atoi(mealCost)
    tipPercent_int,_ := strconv.Atoi(tipPercent)
    taxPercent_int,_ := strconv.Atoi(taxPercent)
    tip := ((mealCost_int * tipPercent_int)/100)
    tax := ((mealCost_int * taxPercent_int)/100)
    totalCost := mealCost_int + tip + tax
    fmt.Println(totalCost)
}