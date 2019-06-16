package main

import "fmt"

func main() {
	cant := 6
	shar := "#"
	// space := "!"
	for i := 0; i < cant; i++ {
		for k := cant; k > i; k-- {
			fmt.Println("Formato ", k)
			fmt.Println(shar)
			shar += "#"
		}

	}

}
