package main

import (
	"fmt"
)

func main() {
	var size int
	var num, sum int64
	var arr []int64
	fmt.Scan(&size)
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		a := map[string]int64{}
		nombre := fmt.Sprintf("num%d", i)
		a[nombre] = num
		arr = append(arr, a[nombre])
	}
	for _, valor := range arr {
		sum = sum + valor
	}
	fmt.Println(sum)
}
