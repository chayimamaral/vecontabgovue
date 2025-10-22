package repos

import (
	"database/sql"
	// ASSUMIR o nome do seu módulo/projeto é 'meu-projeto/api'
	"vecontabgovue/backend/internal/models" // <--- IMPORT CORRETO DO PACOTE 'models'

	_ "github.com/lib/pq"
)

type RegistroRepository struct {
	DB *sql.DB
}

// Agora, usamos models.RegistroData
func (r *RegistroRepository) CreateRegistro(data models.RegistroData) error {
	query := `INSERT INTO empresas (nome, email_admin, senha_admin) VALUES ($1, $2, $3)`
	// Note que data.NomeEmpresa funciona porque RegistroData foi exportada (começa com letra maiúscula)
	_, err := r.DB.Exec(query, data.NomeEmpresa, data.EmailAdmin, data.SenhaAdmin)
	return err
}
