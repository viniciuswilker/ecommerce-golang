package router

import (
	"ecommerce-golang/internal/router/rotas/rotas"

	"github.com/gorilla/mux"
)

func CarregarRotas() *mux.Router {

	r := mux.NewRouter()

	r = rotas.Configurar(r)

	return r

}
