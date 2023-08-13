package main

import (
	"fmt"
	"goweb/controllers"
	"log"
	"net/http"
)

func main() {
	//crear enrutador
	router := http.NewServeMux()

	router.HandleFunc("/", controllers.HomeController)
	router.HandleFunc("/new", controllers.NewGameController)
	router.HandleFunc("/game", controllers.GameController)
	router.HandleFunc("/play", controllers.PlayController)
	router.HandleFunc("/about", controllers.AboutController)

	port := ":3000"
	fmt.Println("Server on port: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
