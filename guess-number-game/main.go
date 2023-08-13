package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	time1 := time.Now()

	hora := time1.Hour()

	//sentencia if
	if hora < 12 {
		fmt.Println("mañana")
	} else if hora < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}

	if time2 := time.Now(); time2.Hour() < 12 {
		fmt.Println("mañana")
	}

	switch time3 := time.Now(); {
	case time3.Hour() < 12:
		fmt.Println("mañana")
	case time3.Hour() < 17:
		fmt.Println("Tarde")
	default:
		fmt.Println("Noche")

	}

	//sentencia switch
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("El sistema operativo es Windows")
	case "linux":
		fmt.Println("El sistema operativo es Linux")
	case "darwin":
		fmt.Println("El sistema operativo es Mac")
	default:
		fmt.Println("otro sistema operativo")

	}

	//bulce for
	//bucle infiniro

	/*for {

	}*/

	//bucle con una condicion
	indice := 0
	for indice < 10 {
		fmt.Println(indice)
		indice++
	}

	//bulce for tipico
	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}

		fmt.Printf("Sofia %d \n", i)
	}

	fmt.Println(suma8B(10, 15))

	saludo := saludar("sofia")
	fmt.Println(saludo)

	sum, mult := suma(10, 3)
	fmt.Println("la suma es:", sum)
	fmt.Println("la multiplicacion es:", mult)

	//proyecto
	fmt.Println(rand.Intn(20))

	jugar()

}

func suma8B(a int8, b int8) int8 {
	return a + b
}

func saludar(name string) string {
	return fmt.Sprintf("hola %s", name)
}

func suma(a, b int) (int, int) {
	sumaNums := a + b
	multNums := a * b
	return sumaNums, multNums

}

func suma2(a, b int) (sumaNums, multNums int) {
	sumaNums = a + b
	multNums = a * b
	return sumaNums, multNums
}

func jugar() {
	numRandom := rand.Intn(100)

	var inputNum int
	var intentos int
	const maxIntentos = 10

	for intentos <= maxIntentos {
		fmt.Printf("Ingresa un numero (Intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&inputNum)

		if inputNum == numRandom {
			fmt.Println("Adivinaste el numero")
			jugarAgain()
			return
		} else if inputNum < numRandom {
			fmt.Println("El numero a adivibar es mayor")
		} else if inputNum > numRandom {
			fmt.Println("El numero a adivibar es menor")
		}

		intentos++
	}

	fmt.Println("SE acabaron los intentos el numero era: ", numRandom)
	jugarAgain()
}

func jugarAgain() {
	var eleccion string
	fmt.Println("Quieres jugar de nuevo s/n")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Fin del juego")
	default:
		fmt.Println("Eleccion invalida")
		jugarAgain()
	}

}
