package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (perro *Perro) Sonido() {
	fmt.Println(perro.Nombre + "Hace gua gua")
}

type Gato struct {
	Nombre string
}

func (gato *Gato) Sonido() {
	fmt.Println(gato.Nombre + "Hace miau miau")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
