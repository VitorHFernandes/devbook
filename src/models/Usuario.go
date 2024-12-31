package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// * Usuarios - representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// * Preparar - chama os métodos para validar e formatar o usuário recebido.
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

// * validar - Valida se as informações estão vazias ou se são válidas.
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o campo 'nome' é obrigatório e não pode estar vazio")
	}

	if usuario.Nick == "" {
		return errors.New("o campo 'nick' é obrigatório e não pode estar vazio")
	}

	if usuario.Email == "" {
		return errors.New("o campo 'email' é obrigatório e não pode estar vazio")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("o campo 'senha' é obrigatório e não pode estar vazio")
	}

	return nil
}

// * formatar - Remove os espaçamentos de nome, nick e e-mail.
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
