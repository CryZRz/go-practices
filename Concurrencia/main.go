package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	//declaracion del canal
	//canal := make(chan int)
	//de esta froma se mandan datos al canal
	//canal <- 15
	//de esta manera recbimos datos del canal
	//valor := <-canal

	//fmt.Println(<-canal)

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	chanelApis := make(chan string)

	for _, api := range apis {
		//checkAPI(api) //sin usar concurrencia
		go checkAPI(api, chanelApis) //usando concurrencia
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-chanelApis)
	}

	alapsed := time.Since(start)
	fmt.Printf("¡Lista! !Tomo %v segundos¡\n", alapsed.Seconds())

}

func checkAPI(api string, chanel chan string) {
	_, err := http.Get(api)

	if err != nil {
		chanel <- fmt.Sprintf("Error al consultar la api %s \n", api)
		return
	}

	chanel <- fmt.Sprintf("La api %s esta en funcionamiento\n", api)

}
