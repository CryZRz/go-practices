package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Constantes usar a nivel de paquete
const pi = 3.1416

const (
	PI    = math.Pi
	EULER = math.E
	x     = 100
	y     = 0b1010 //Binario
	z     = 0o12   //Octal
	w     = 0xFF   //Hexadecimal
)

const (
	Domingo = iota + 1 //Incrementa el valor secuencialmete empezando desde 0
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	//Declaracion de variables
	var firstName, lastName string
	var age int

	var (
		name = "Jhon" //Se infiere el tipo
		work = "Arquitec"
		time int
	)

	firstName = "Christian"
	lastName = "Ramos"
	age = 19

	name = "Jhon"
	time = 123

	//Infiere el tipo sin usar "var", solo se usa para declarar
	country := "Mexico"
	country = "Greace"

	fmt.Println(firstName, lastName, age)
	fmt.Println(name, work, time, country)

	fmt.Println(x, y, z, w, EULER)

	fmt.Println(Viernes)

	//Tipos de datos
	//var integer uint = 25
	//var float float32 = math.E
	fmt.Println(math.MinInt64, math.MaxInt64)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxUint8)

	//valueBool := true
	//valueBool = false

	fullName := "Christian Ramos \t(alias \"CryzDev\")"
	fmt.Println(fullName)

	//byte alamacena datos no negativos hasta 255 similar a uint8
	var byteAlc byte = 'a'
	palabra := "christian"

	fmt.Println(byteAlc)
	fmt.Println(palabra[0])

	var emoji rune = '🌵'
	fmt.Println(emoji)

	//conversion de tipo
	var numeroP int16 = 50
	var numeroP2 int32 = 100

	stringP := "100"
	i, _ := strconv.Atoi(stringP)

	fmt.Println(i + 20)
	fmt.Println(numeroP + int16(numeroP2))

	numeroP3 := 42
	stringP = strconv.Itoa(numeroP3)
	fmt.Println(stringP)

	//funciones libreria fmt

	//printf formatear informacion
	nameUriel := "uriel"
	ageUriel := 19
	fmt.Printf("Hola me llamo: %s y tengo %d años. \n", nameUriel, ageUriel)

	greeting := fmt.Sprintf("Hola me llamo: %s y tengo %d años. \n", "Jieman", 19)
	fmt.Print(greeting)

	//Entrada de datos
	var nameKarla string
	var ageKarla int

	fmt.Print("Ingresa tu nombre")
	fmt.Scanln(&nameKarla)

	fmt.Print("Ingresa tu edad")
	fmt.Scanln(&ageKarla)

	fmt.Printf("Hola me llamo: %v y tengo %d años. \n", nameKarla, ageKarla)
}
