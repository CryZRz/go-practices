package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func printList(list ...any) {
	for _, li := range list {
		fmt.Println(li)
	}
}

type integer int

// restricciones
type Numbers interface {
	~int | ~float64 | ~float32 | uint
}

func sumTwo[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

/*
~se usa para aproximar tipos de datos
*/
func sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

type Product[T uint | string] struct {
	Id          T
	Description string
	Price       float32
}

func main() {
	var num1 integer = 255
	var num2 integer = 225

	printList(sum[float64](1.32, 25.32, 25.12, 30.25, 15.5))
	printList(sum[integer](num1, num2))
	printList(sum[int](1, 25, 25, 30, 15))
	printList(sumTwo[uint](1, 25, 25, 30, 15))

	producto := Product[uint]{1, "Zapato", 50}
	producto2 := Product[string]{"12", "Zapato", 50}
	fmt.Println(producto)
	fmt.Println(producto2)

}
