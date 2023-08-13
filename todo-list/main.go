package main

import "fmt"

// Estructuras
type Persona struct {
	nombre    string
	apellidos string
	edad      int
}

//metodo de una estructura recibe el reveptor
func (p *Persona) diHola() {
	fmt.Println("Hola mi nombre es: ", p.nombre)
}

func main() {
	var array [5]int
	var array2 = [5]int{1, 2, 3, 4, 5}
	var array3 = [...]int{10, 20, 30, 40, 50}
	array[0] = 1
	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)

	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}

	for _, value := range array3 {
		fmt.Println(value * 2)
	}

	//matriz
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz)

	for i := 0; i < len(matriz); i++ {
		for y := 0; y < len(matriz[i]); y++ {
			fmt.Println(matriz[i][y])
		}
	}

	//rebanadas o listas
	diasSemana := []string{"lunes", "martes", "miercoles", "jueves", "viernes", "sabado", "domingo"}

	diasRebanda := diasSemana[0:3]
	fmt.Println(diasRebanda)

	fmt.Println(len(diasRebanda))
	fmt.Println(cap(diasRebanda))

	var rebanada []int
	rebanada = append(rebanada, 25, 30, 15)
	fmt.Println(rebanada)

	rebanada = append(rebanada[:1], rebanada[3:]...)
	fmt.Println(rebanada)

	//crea una lista de tipo string con 5 espacios inciales y 10 de capacidad inicial
	nombres := make([]string, 3, 5)
	fmt.Println(nombres)
	nombres[0] = "Karla"
	nombres[1] = "Christian"
	nombres[2] = "Aron"

	nombres = append(nombres, "Jiemna", "Zaida", "Jasso")
	fmt.Println(nombres)

	numeros := []int{1, 2, 3, 4, 5}
	numeros2 := make([]int, 5)

	copy(numeros2, numeros)

	fmt.Println(numeros)
	fmt.Println(numeros2)

	var persona Persona
	persona.nombre = "Christian"
	persona.apellidos = "Ramos"
	persona.edad = 19

	fmt.Println(persona)

	persona2 := Persona{"Jiemna", "Navarro", 19}
	fmt.Println(persona2)
	fmt.Println(persona2.nombre)

	var x = 10
	var p *int = &x
	editar(p)
	fmt.Println(*p)

	persona2.diHola()
}

func editar(x *int) {
	*x = 20
}
