package main

import "fmt"

// funcion variadica
func suma(nums ...int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	return totalSum
}

// recibe n cantidad de datos y de distintos tipos
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v \n", dato, dato)
	}
}

// recursividad
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func saludar(nombre string, callback func(string)) {
	callback(nombre)
}

func multiplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(callback func(int) int, number int) int {
	return callback(number)
}

/*
Funcion de orden superior
Es aquella funcion que recibe como argumento una funcion y
devuelve tambien una funcion como resultado
*/
func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

/*
Closures
un clousure es una funciona anonima que se define dentro de otra funcion
y que puede acceder a los valores de la funcion externa
*/
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	getTotal := suma(12, 15, 16, 12, 16, 10)
	fmt.Println(getTotal)
	imprimirDatos("Zaida", 25, true, 23.6, "jiemna")
	fmt.Println(factorial(5))

	//funciones aninimas
	hola := func(nombre string) {
		fmt.Printf("Hola %s \n", nombre)
	}

	saludar("Zaida", hola)

	resultOne := aplicar(multiplicar, 10)
	resultTwo := aplicar(triplicar, 10)
	fmt.Println(resultOne)
	fmt.Println(resultTwo)

	resultThree := double(addOne, 3)
	fmt.Println(resultThree)

	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
