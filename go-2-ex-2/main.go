package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	for i := 2; i < 4; i++ {
		fibs[i+1] = fibs[i-1] + fibs[i]
	}
	// TODO: correct up to index 4 using direct element access

	for j := 4; j <= 7; j++ {
		fibs = append(fibs, fibs[j-1]+fibs[j])
	}
	// TODO: replace 0 with the next Fibonacci number
	// TODO: compute three more Fibonacci numbers and append them

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
