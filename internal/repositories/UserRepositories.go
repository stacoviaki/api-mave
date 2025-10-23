package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
)

// UserRepositories é a camada responsável por acessar o banco de dados
// e executar operações na tabela `users` (SELECT, INSERT, UPDATE, DELETE).
type UserRepositories struct {
	connection *sql.DB // conexão ativa com o PostgreSQL
}

// NewUserRepositories cria uma nova instância do repositório de usuários,
// recebendo a conexão com o banco como dependência.
func NewUserRepositories(connection *sql.DB) UserRepositories {
	return UserRepositories{
		connection: connection,
	}
}

// GetUsers busca todos os usuários cadastrados no banco de dados.
func (us *UserRepositories) GetUsers() ([]models.User, error) {
	query := "SELECT * FROM public.users"

	// Executa a query no banco.
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.User{}, err
	}

	var userList []models.User // lista com todos os usuários
	var userObj models.User    // armazena temporariamente cada linha

	// Percorre cada linha retornada e preenche os dados no struct
	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.UserName,
			&userObj.Email,
			&userObj.PasswordHash,
			&userObj.CreatedAt,
			&userObj.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []models.User{}, err
		}

		// Adiciona o usuário lido à lista
		userList = append(userList, userObj)
	}

	rows.Close() // fecha o cursor do banco
	return userList, nil
}

// GetUserById busca um único usuário pelo seu UUID.
func (us *UserRepositories) GetUserById(uuid_user uuid.UUID) (*models.User, error) {
	// Prepara a query SQL com parâmetro (evita SQL Injection)
	query, err := us.connection.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user models.User

	// Executa a query e armazena o resultado no struct
	err = query.QueryRow(uuid_user).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		// Se não encontrar nenhuma linha, retorna nil sem erro
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &user, nil
}

// CreateUser insere um novo usuário na tabela `users` e retorna o UUID gerado.
func (us *UserRepositories) CreateUser(user models.User) (uuid.UUID, error) {
	var id uuid.UUID

	// Cria o comando SQL com RETURNING id para capturar o UUID criado
	query, err := us.connection.Prepare(`
		INSERT INTO users (user_name, email, password_hash)
		VALUES ($1, $2, $3) RETURNING id
	`)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	// Executa o INSERT e pega o ID retornado pelo banco
	err = query.QueryRow(user.UserName, user.Email, user.PasswordHash).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	query.Close()
	return id, nil
}

// UpdateUser atualiza os dados de um usuário existente no banco.
func (repo *UserRepositories) UpdateUser(id_user uuid.UUID, user models.User) (*models.User, error) {
	// Query que atualiza e retorna o registro atualizado
	query := `
		UPDATE users
		SET user_name = $1, email = $2, password_hash = $3
		WHERE id = $4
		RETURNING id, user_name, email, password_hash;
	`

	// Executa a query e armazena o resultado em uma linha
	row := repo.connection.QueryRow(query, user.UserName, user.Email, user.PasswordHash, id_user)

	var updatedUser models.User

	// Lê os valores retornados e preenche o struct
	err := row.Scan(&updatedUser.ID, &updatedUser.UserName, &updatedUser.Email, &updatedUser.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

// DeleteUser remove um usuário do banco de dados pelo seu UUID.
func (us *UserRepositories) DeleteUser(id uuid.UUID) (uuid.UUID, error) {
	var deletedID uuid.UUID

	// Executa o DELETE e retorna o ID apagado
	query := `DELETE FROM users WHERE id = $1 RETURNING id;`
	err := us.connection.QueryRow(query, id).Scan(&deletedID)
	if err != nil {
		fmt.Println("Erro ao deletar usuário:", err)
		return uuid.Nil, err
	}

	return deletedID, nil
}
