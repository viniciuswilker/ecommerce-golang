package main

import (
	"ecommerce-golang/internal/config"
	"ecommerce-golang/internal/database"
	"ecommerce-golang/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.CarregarConfig()
	database.ConectaBanco()

	r := router.CarregarRotas()

	fmt.Printf("Rodando na porta %v\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Porta), r))

}
