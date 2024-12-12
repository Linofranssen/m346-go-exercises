package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

type ShortSides struct {
	a float64
	b float64
}

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println(computeHypotenuse(3, 4))
	s1 := ShortSides{a: 3, b: 4}
	fmt.Print(s1.Hypotenuse())
}
