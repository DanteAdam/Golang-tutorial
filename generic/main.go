package main

import "fmt"

func main() {
	result := add(2.25, 2.5)
	fmt.Println(result)
}

func add[T int | string | float64](a, b T) T {
	return a + b

}
