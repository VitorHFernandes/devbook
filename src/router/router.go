package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// * Generate - retorna um router com as rotas configuradas.
func Generate() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
