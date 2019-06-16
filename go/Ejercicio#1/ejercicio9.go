package main

import "fmt"

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	var d, h int
	fmt.Scan(&d, &h)
	// e = 1
	sum := make([]int, 0)
	rev := make([]int, 0)
	for i := d; i <= h; i++ {
		sum = append(sum, i)
	}
	for j := 0; j < len(sum); j++ {
		r := reverse(sum[j])
		rev = append(rev, r)
	}
	fmt.Println(sum)
	fmt.Println(rev)
}
