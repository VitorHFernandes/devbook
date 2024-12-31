package seguranca

import "golang.org/x/crypto/bcrypt"

//* Hash - Recebe uma string e coloca um hash nela.
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//* VerificarSenha - Compara uma senha com um hash e retorna caso sejam iguais.
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
