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

	r := router.Generate()

	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
