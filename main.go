package main

import (
	"challeng-bravo/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Escutando na porta 3000")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":3000", nil), r)
}
