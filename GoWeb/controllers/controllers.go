package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	templateLaod, error := template.ParseFiles("templates/base.html", "templates/index.html")

	if error != nil {
		http.Error(w, "error al cargar el html", http.StatusInternalServerError)
	}

	/*
		data := struct {
			Title   string
			Message string
		}{
			Title:   "Pagina inicio",
			Message: "Bienvenido a piedra papel y tijera",
		}
	*/

	//error = templateLaod.Execute(w, nil)
	error = templateLaod.ExecuteTemplate(w, "base", nil)

	if error != nil {
		http.Error(w, "error al renderizar el html", http.StatusInternalServerError)
	}

}

func NewGameController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintln(w, "Crear nuevo juego")
}

func GameController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintln(w, "Jugo")
}

func PlayController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintln(w, "JUgar")
}

func AboutController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintln(w, "Acerca")
}
