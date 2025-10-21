package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
)

type UserRepositories struct {
	connection *sql.DB
}

func NewUserRepositories(connection *sql.DB) UserRepositories {
	return UserRepositories{
		connection: connection,
	}
}

func (us *UserRepositories) GetUsers() ([]models.User, error) {
	query := "SELECT * FROM public.users"
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.User{}, err
	}
	var userList []models.User
	var userObj models.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.UserName,
			&userObj.Email,
			&userObj.PasswordHash,
			&userObj.CreatedAt,
			&userObj.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			return []models.User{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil

}
func (us *UserRepositories) GetUserById(uuid_user uuid.UUID) (*models.User, error) {
	query, err := us.connection.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var user models.User
	err = query.QueryRow(uuid_user).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		//Não encontrou nenhum registro com o id que foi passado
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &user, nil
}

func (us *UserRepositories) CreateUser(user models.User) (uuid.UUID, error) {
	var id uuid.UUID

	query, err := us.connection.Prepare(`
		INSERT INTO users (user_name, email, password_hash)
		VALUES ($1, $2, $3) RETURNING id
	`)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	err = query.QueryRow(user.UserName, user.Email, user.PasswordHash).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	query.Close()
	return id, nil
}
