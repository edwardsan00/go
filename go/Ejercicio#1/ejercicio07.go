package main

import (
	"fmt"
)

func main() {
	var size int
	size = 3
	// fmt.Scan(&size)
	// for i := 0; i < size; i++ {
	// 	arr := map[]int
	// 		contenedor := map[string]int{}
	// 		nombcont := fmt.Sprintf("n%d", i)
	// 		contenedor[nombcont] = i
	// 		arr = append(arr, contenedor[nombcont])
	// }
	// fmt.Println(arr)
	aos := [3][]string{}
	// fmt.Scan(&size)
	for i := 0; i < size; i++ {

		// pos := fmt.Sprintf("%d", i)
		aos[1] = []string{"1", "2", "3"}
	}
	fmt.Println(aos)

}
