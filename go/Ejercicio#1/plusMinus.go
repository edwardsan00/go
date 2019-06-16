package main

import (
	"fmt"
)

func main() {
	var numArray, valArray int
	var positive, negative, zero int
	var resultPositive, resultNegative, resultZero float32
	var arr []int
	fmt.Scan(&numArray)
	for i := 0; i < numArray; i++ {
		fmt.Scan(&valArray)
		matriz := map[string]int{}
		name := fmt.Sprintf("num%d", i)
		matriz[name] = valArray
		arr = append(arr, matriz[name])
	}
	size := len(arr)
	for j := 0; j < size; j++ {
		if arr[j] > 0 {
			positive++
		} else if arr[j] < 0 {
			negative++
		} else {
			zero++
		}
	}
	resultPositive = (float32(positive) / float32(size))
	resultNegative = (float32(negative) / float32(size))
	resultZero = (float32(zero) / float32(size))
	fmt.Printf("%.5f\n", resultPositive)
	fmt.Printf("%.5f\n", resultNegative)
	fmt.Printf("%.5f\n", resultZero)
}
