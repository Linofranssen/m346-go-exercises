package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float64, maxPoints float64) float64 {
	return gotPoints/maxPoints*5 + 1
}

func main() {
	// TODO: call the function computeGrade
	fmt.Print(computeGrade(23.5, 30), " & ", computeGrade(28, 28))
}
