package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		if i != 3 {
			fmt.Println("Hola " + strconv.Itoa(i))
		}
	}
}

//Do while
