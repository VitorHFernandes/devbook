package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
)

// * Logger - Escreve informações da requisição no terminal.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("| Request received |",
			"Method", r.Method,
			"URI", r.RequestURI,
			"HOST", r.Host,
		)
	}
}

// * Autenticar - Verifica se o usuário realizando a requisição está autenticado.
func Autenticar(next http.HandlerFunc) http.HandlerFunc { //? func (w http.ResponseWriter, r *http.Request)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}
}
