package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors["azul"])
	colors["negro"] = "#000000"

	fmt.Println(colors)
	valor, ok := colors["blanco"]

	if ok {
		fmt.Println("si existe la clave")
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave")
		fmt.Println(ok)
	}

	delete(colors, "rojo")
	fmt.Println(colors)

	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s \n", clave, valor)
	}

}
