package main

import "fmt"

func main() {
	var arr []int
	var num, size, suma int
	fmt.Scan(&size)
	for i := 1; i <= size; i++ {
		fmt.Scan(&num)
		a := map[string]int{}
		name := fmt.Sprintf("num%d", i)
		a[name] = num
		arr = append(arr, a[name])
	}
	for _, num := range arr {
		suma = suma + num
	}
	fmt.Println(suma)
}
