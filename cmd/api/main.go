package main

import (
	"ecommerce-golang/internal/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.CarregarConfig()

	r := mux.NewRouter()

	fmt.Printf("Rodando na porta %v\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Porta), r))

}
