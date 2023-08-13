package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		//return 0, errors.New("No se puede divir por 0")
		return 0, fmt.Errorf("No se puede divir por 0")
	}
	return dividendo / divisor, nil
}

func divideTwo(dividendo, divisor int) (int, error) {
	validateZero(divisor)
	return dividendo / divisor, nil
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede divir por 0")
	}
}

func divideThree(dividendo, divisor int) (int, error) {
	defer func() {
		//captura si ocurre un panico recover recoje el panico similar al catch de otros lenguajes
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	return dividendo / divisor, nil
}

func main() {
	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(num)
	}

	result, error := divide(10, 2)

	if error != nil {
		fmt.Println("Error: ", error)
		return
	} else {
		fmt.Println("Resultado: ", result)
	}

	//uso defer
	defer fmt.Println(3) //pospone la ejecucion y se guarda en una pila de ejecucion
	defer fmt.Println(1)
	defer fmt.Println(2)

	file, errorFile := os.Create("hola.txt")

	if errorFile != nil {
		fmt.Println(errorFile)
		return
	}

	defer file.Close()

	_, errorWrite := file.Write([]byte("Christian es muy guapo"))

	if errorWrite != nil {
		fmt.Println(errorFile)
		return
	}

	//panic y cover
	divideTwo(100, 2)
	divideThree(200, 5)
	divideTwo(100, 4)

	log.Println("Este es un mensaje de registro")
	//log.Fatal("Este es un mensaje de error")

	log.SetPrefix("main()")
	fileLog, errorFile2 := os.OpenFile("info.lg", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if errorFile2 != nil {
		log.Fatal(errorFile2)
	}

	defer file.Close()

	log.SetOutput(fileLog)
	log.Print("Oye soy un log")
}
