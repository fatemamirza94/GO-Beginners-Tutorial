package main

import "fmt"

// define a type constraint that includes more than one primitive type with the union | operator
type Number interface {
	int64 | float64
}

func main() {
	// create maps of ints and float values
	ints := map[string]int64{
		"first":  132,
		"second": 65,
	}

	floats := map[string]float64{
		"first":  32.3,
		"second": 156.2,
	}

	// print sum of all values in map (type is inferred)
	fmt.Printf("Generic sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}

// function for summing the values (int or float) in a map
// declare type parameters in [] brackets, e.g. K and V. Type constraints (comparable/Number) specify permissible type arguments.
// declare function parameters in () brackets as usual, where it is possible to use the type parameters K and V
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
