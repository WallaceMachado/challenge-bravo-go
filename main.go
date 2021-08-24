package main

import (
	"challeng-bravo/src/config"
	"challeng-bravo/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Loader()

	fmt.Println("Server is running!")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
