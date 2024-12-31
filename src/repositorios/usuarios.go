package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// * Usuarios - representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// * NovoRepositorioDeUsuarios() - Cria um repositório de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// * Criar() - Insere um usuário no banco de dados.
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios(nome, nick, email, senha) values(?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// * Buscar - Retorna todos os usuários filtrando por nome ou nick.
func (repositorios Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorios.db.Query("SELECT id, nome, nick, email, criadoEm from usuarios WHERE nome LIKE ? OR nick LIKE ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// * BuscarPorID - Retorna um usuário do banco de dados.
func (repositorios Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := repositorios.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		return models.Usuario{}, erro //TODO - Não permite nil, usar models.Usuario{} para retorno vazio (ausência de usuário).
	}
	defer linhas.Close()

	var usuario models.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}

// * Atualizar - Realiza alterações no banco em determinado usuário filtrado pelo ID.
func (repositorios Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, erro := repositorios.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// * Excluir - Realiza a exclusão de um usuário no banco de dados filtrado pelo ID.
func (repositorios Usuarios) Excluir(ID uint64) error {
	statement, erro := repositorios.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil

}

// * BuscarPorEmail - Busca um usuário por e-mail e retorna seu ID e senha com hash.
func (repositorios Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorios.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}
