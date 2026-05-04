package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	Porta       = 0
	StringBanco = ""
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

	StringBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

}
