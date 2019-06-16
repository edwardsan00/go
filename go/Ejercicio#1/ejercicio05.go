package main

import "fmt"

func main() {
	size := 3
	var num1, num2 int
	var alice []int
	var bob []int
	var pointAlice, pointBob int
	for i := 0; i < size; i++ {
		fmt.Scan(&num1)
		a := map[string]int{}
		nombre := fmt.Sprintf("num%d", i)
		a[nombre] = num1
		alice = append(alice, a[nombre])
	}
	for j := 0; j < size; j++ {
		fmt.Scan(&num2)
		b := map[string]int{}
		nombre := fmt.Sprintf("num%d", j)
		b[nombre] = num2
		bob = append(bob, b[nombre])
	}
	for i := 0; i < size; i++ {
		switch {
		case alice[i] > bob[i]:
			pointAlice++
		case bob[i] > alice[i]:
			pointBob++
		}
	}
	fmt.Printf("%d %d", pointAlice, pointBob)
}
