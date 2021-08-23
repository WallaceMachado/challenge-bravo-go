package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
