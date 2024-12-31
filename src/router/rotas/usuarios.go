package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuarios,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.ExcluirUsuario,
		RequerAuth: false,
	},
}
