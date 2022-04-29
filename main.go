package main

import (
	"api-TalkCond/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando a API")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
