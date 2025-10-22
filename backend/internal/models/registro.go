package models

// RegistroData define a estrutura de dados (campos de input ou de DB) para o Registro/Empresa
// As tags `json` e `db` são comuns em Go.
type RegistroData struct {
	ID          string `json:"id" db:"id"`
	NomeEmpresa string `json:"nome_empresa" db:"nome_empresa"`
	EmailAdmin  string `json:"email_admin" db:"email_admin"`
	SenhaAdmin  string `json:"senha_admin" db:"senha_admin"`
	// Adicione aqui todos os campos que representam a empresa e o primeiro usuário.
}
