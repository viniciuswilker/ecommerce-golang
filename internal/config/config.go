package config

import (
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	Porta = 0
)

func CarregarConfig() {

	var err error

	if err = gotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 8080
	}

}
