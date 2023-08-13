package main

import (
	"fmt"
	"poo/animal"
	"poo/book"
)

func main() {
	myFirsBook := book.NewBook("C++ programing languaje", "Bjaren Sturten", 800)
	mySecondBook := book.NewBook("C programing languaje", "Dennis Ritche", 600)

	myFirsBook.PrintInfo()
	mySecondBook.PrintInfo()
	myFirsBook.SetPages(750)
	myFirsBook.PrintInfo()

	fmt.Println(mySecondBook.GetPages())

	myTextBook := book.NewTextBook("Java programing languaje", "James Gosling", 550, "Suns Microsistem", "university")

	myTextBook.PrintInfo()

	book.Print(myTextBook)

	miDog := animal.Perro{Nombre: "Nacho"}
	miCat := animal.Gato{Nombre: "Chispas"}
	//miDogTwo := animal.Perro{Nombre: "Firulas"}
	//miCatTwo := animal.Gato{Nombre: "Peluche"}

	animal.HacerSonido(&miDog)
	animal.HacerSonido(&miCat)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Nacho"},
		&animal.Gato{Nombre: "Chispas"},
		&animal.Perro{Nombre: "Firulas"},
		&animal.Gato{Nombre: "Peluche"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
